package models

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"math"
	"time"
)

type User struct {
	gorm.Model
	Passport        string  `json:"passportNumber" gorm:"size:50;unique;not null"`
	FirstName       string  `json:"FirstName" gorm:"size:100;not null;default:'Иван'"`
	LastName        string  `json:"LastName" gorm:"size:100;not null;default:'Иванов'"`
	Patronymic      string  `json:"Patronymic" gorm:"size:100;not null;default:'Иванович'"`
	Address         string  `json:"Address" gorm:"size:255;not null;default:'г. Москва, ул. Пушкина, д. 1'"`
	CountTask       int     `json:"countTask" gorm:"not null;default:0"`
	OverallDuration float64 `json:"overallDuration" gorm:"not null;default:0"`

	TimeEntries []TimeEntry `gorm:"foreignKey:UserPassport;references:Passport"`
}

type TimeEntry struct {
	gorm.Model
	UserPassport string `gorm:"not null"`
	Name         string `gorm:"default:'Template'"`
	StartTime    time.Time
	EndTime      time.Time
	Duration     float64 `gorm:"default:0"`
	User         User    `gorm:"foreignKey:UserPassport;references:Passport"`
}

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func (user User) Create() error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	log.Println("Creating user:", user)
	result := db.Create(&user)
	return result.Error
}

func (user User) Update() error {
	var oldUser User
	if err := db.Where("passport = ?", user.Passport).First(&oldUser).Error; err != nil {
		return err
	}
	if err := db.Model(&oldUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateTask(passport string, taskName string) error {
	var user User
	if err := db.Where("passport = ?", passport).First(&user).Error; err != nil {
		return err
	}
	task := TimeEntry{
		UserPassport: passport,
		StartTime:    time.Now(),
		User:         user,
		Name:         taskName,
	}
	result := db.Create(&task)
	user.CountTask += 1
	db.Save(&user)
	return result.Error
}

func EndTask(passport string, taskID uint64) (float64, error) {
	var task TimeEntry
	var user User
	if err := db.Where("user_passport = ? AND id = ?", passport, taskID).Last(&task).Error; err != nil {
		return 0, err
	}
	if err := db.Where("passport = ?", passport).First(&user).Error; err != nil {
		return 0, err
	}
	duration := time.Now().Sub(task.StartTime)
	user.OverallDuration += math.Round(duration.Seconds())
	task.Duration = math.Round(duration.Seconds())
	task.EndTime = time.Now()
	db.Save(&task)
	db.Save(&user)
	return math.Round(duration.Seconds()), nil
}

func (user User) Delete() error {
	result := db.Delete(&user, "passport = ?", user.Passport)
	return result.Error
}

func (user User) FormattedName() string {
	return user.LastName + " " + user.FirstName
}

func (user User) FormattedCreatedAt() string {
	return user.CreatedAt.Format("15-01-2006")
}

func (user User) FormattedUpdatedAt() string {
	return user.UpdatedAt.Format("15-01-2006")
}

func (entry TimeEntry) FormattedDuration() string {
	sec := int(math.Round(entry.Duration))
	hours := sec / 3600
	minutes := (sec % 3600) / 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, sec%60)
}

func (entry TimeEntry) FormattedStartTime() string {
	return entry.StartTime.Format("02-01-2006 15:04:05")
}

func (entry TimeEntry) FormattedEndTime() string {
	return entry.EndTime.Format("15-01-2006 15:04:05")
}

func (entry TimeEntry) FormattedName() string {
	return entry.User.FormattedName()
}
