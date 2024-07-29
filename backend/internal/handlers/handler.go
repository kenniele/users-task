package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"users-task-project/internal/database"
	"users-task-project/internal/models"
)

func indexHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limitStr := c.DefaultQuery("limit", "6")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 6
	}

	offset := (page - 1) * limit
	users, totalPages, err := database.GetUsers(limit, offset)
	if err != nil {
		log.Printf("Error - getting users: %v", err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
	}

	var pageNumbers []int
	for i := 1; i <= totalPages; i++ {
		pageNumbers = append(pageNumbers, i)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users":       users,
		"totalPages":  totalPages,
		"page":        page,
		"prev":        page - 1,
		"next":        page + 1,
		"limit":       limit,
		"pageNumbers": pageNumbers,
	})
}

func filterHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limitStr := c.DefaultQuery("limit", "6")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 6
	}

	offset := (page - 1) * limit

	filter := models.User{
		Passport:   c.DefaultQuery("Passport", ""),
		FirstName:  c.DefaultQuery("FirstName", ""),
		LastName:   c.DefaultQuery("LastName", ""),
		Patronymic: c.DefaultQuery("Patronymic", ""),
		Address:    c.DefaultQuery("Address", ""),
	}

	users, totalPages, err := database.FilterUsers(filter, limit, offset)
	if err != nil {
		log.Printf("Error - filtering users: %v", err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	var pageNumbers []int
	for i := 1; i <= totalPages; i++ {
		pageNumbers = append(pageNumbers, i)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"users":       users,
		"totalPages":  totalPages,
		"page":        page,
		"limit":       limit,
		"pageNumbers": pageNumbers,
	})
}

func registerHTMLHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"Error": nil,
	})
}

func updateHTMLHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", gin.H{
		"Error": nil,
	})
}

func deleteHTMLHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "delete.html", gin.H{
		"Error": nil,
	})
}

func manageHTMLHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	completedTasks, currentTasks, err := database.GetTasksByPassport(Passport)
	if err != nil {
		log.Printf("Error - getting tasks: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"Error":          nil,
		"completedTasks": completedTasks,
		"currentTasks":   currentTasks,
		"UserPassport":   Passport,
	})
}

// @Summary Get user information by passport
// @Description Get detailed user information using passport data
// @Tags users
// @Param PassportSeries path string true "Passport Series"
// @Param PassportNumber path string true "Passport Number"
// @Success 200 {object} map[string]string "User data"
// @Failure 400 {object} map[string]string "Error"
// @Router /info/{PassportSeries}/{PassportNumber} [get]
func infoUserHandler(c *gin.Context) {
	passportNumber := c.Param("PassportNumber")
	passportSeries := c.Param("PassportSeries")
	passport := passportSeries + " " + passportNumber
	user, err := database.GetUserByPassport(passport)
	if err != nil {
		log.Printf("Error - receiving user from database: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
			"data":  nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
		"data": map[string]string{
			"surname":    user.LastName,
			"name":       user.FirstName,
			"patronymic": user.Patronymic,
			"address":    user.Address,
		},
	})
}

// @Summary Create a new user
// @Description Create a user with the passport and other details
// @Tags users
// @Param Passport path string true "Passport"
// @Param User body models.User true "User data"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /create/{Passport} [post]
func createUserHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	if err := models.PassportValidation(Passport); err != nil {
		log.Printf("Error - binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	log.Printf("Passport: %s", Passport)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Error - binding json - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	user.Passport = Passport

	err := user.Create()
	if err != nil {
		log.Printf("Error - creating user - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
	})
}

// @Summary Begin a task for a user
// @Description Start a new task for a user with a specific passport
// @Tags tasks
// @Param Passport path string true "Passport"
// @Param TaskName path string true "Task name"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /man/begin/{Passport}/{TaskName} [post]
func beginTaskHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	TaskName := c.Param("TaskName")

	if err := models.CreateTask(Passport, TaskName); err != nil {
		log.Printf("Error - creating task - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
	})
}

// @Summary End a task for a user
// @Description End a task for a user with a specific passport and task ID
// @Tags tasks
// @Param Passport path string true "Passport"
// @Param TaskID path int true "Task ID"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /man/end/{Passport}/{TaskID} [post]
func endTaskHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	TaskID, err := strconv.ParseUint(c.Param("TaskID"), 10, 64)
	if err != nil {
		log.Printf("Error - parsing task ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}

	duration, err := models.EndTask(Passport, TaskID)
	if err != nil {
		log.Printf("Error - ending task - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Error":    nil,
		"Duration": duration,
	})
}

// @Summary Update user details
// @Description Update user information using passport
// @Tags users
// @Param Passport path string true "Passport"
// @Param User body models.User true "Updated user data"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /update/{Passport} [put]
func updateUserHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	if err := models.PassportValidation(Passport); err != nil {
		log.Printf("Error - validating passport: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Error - binding json - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	log.Println("New user:", user)

	err := user.Update()
	if err != nil {
		log.Printf("Error - updating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
	})
}

// @Summary Delete a user
// @Description Delete a user with provided passport
// @Tags users
// @Param Passport path string true "Passport"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} map[string]interface{} "Error"
// @Router /delete/{Passport} [delete]
func deleteUserHandler(c *gin.Context) {
	Passport := c.Param("Passport")
	if err := models.PassportValidation(Passport); err != nil {
		log.Printf("Error - validating json - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	var user models.User
	user = models.User{
		Passport: Passport,
	}

	err := user.Delete()
	if err != nil {
		log.Printf("Error - deleting user - %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
	})
}
