package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"reflect"
	"users-task-project/config"
	"users-task-project/internal/models"
)

var db *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error - failed to connect to database: %v", err)
		return nil, err
	}
	if err = db.AutoMigrate(&models.User{}, &models.TimeEntry{}); err != nil {
		log.Printf("Error - failed to migrate table: %v", err)
		return nil, err
	}
	log.Printf("Database migrated")
	return db, nil
}

func GetUsers(limit int, offset int) ([]models.User, int, error) {
	var users []models.User
	var totalCount int64

	db.Model(&models.User{}).Count(&totalCount)
	result := db.Offset(offset).Limit(limit).Find(&users)

	if result.Error != nil {
		return []models.User{}, 1, result.Error
	}

	totalPages := int(totalCount) / limit
	if int(totalCount)%limit > 0 {
		totalPages++
	}
	return users, totalPages, nil
}

func GetUserByPassport(passport string) (*models.User, error) {
	var user models.User

	result := db.Where("passport = ?", passport).First(&user)
	return &user, result.Error
}

func GetTasksByPassport(passport string) ([]models.TimeEntry, []models.TimeEntry, error) {
	var completedTasks []models.TimeEntry
	var currentTasks []models.TimeEntry
	resultCompleted := db.Where("user_passport = ? AND duration > ?", passport, 0).Order("duration DESC").Find(&completedTasks)
	resultCurrent := db.Where("user_passport = ? AND duration = ?", passport, 0).Find(&currentTasks)
	if resultCompleted.Error != nil {
		return nil, nil, resultCompleted.Error
	}
	if resultCurrent.Error != nil {
		return nil, nil, resultCurrent.Error
	}
	return completedTasks, currentTasks, nil
}

func FilterUsers(filter models.User, limit int, offset int) ([]models.User, int, error) {
	var users []models.User
	var totalCount int64
	conditions := map[string]interface{}{}

	val := reflect.ValueOf(filter)
	typ := reflect.TypeOf(filter)
	reformat := map[string]string{
		"Passport":   "passport",
		"FirstName":  "first_name",
		"LastName":   "last_name",
		"Patronymic": "patronymic",
		"Address":    "address",
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if field.Kind() == reflect.String && field.String() != "" {
			conditions[reformat[fieldType.Name]] = field.String()
		}
	}

	dbQuery := db.Where(conditions).Find(&users)

	if err := dbQuery.Model(&models.User{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	result := dbQuery.Offset(offset).Limit(limit).Find(&users)

	if result.Error != nil {
		return nil, 1, result.Error
	}

	totalPages := int(totalCount) / limit
	if int(totalCount)%limit > 0 {
		totalPages++
	}
	return users, totalPages, nil
}
