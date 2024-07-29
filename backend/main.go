package main

import (
	"github.com/joho/godotenv"
	"log"
	"users-task-project/config"
	"users-task-project/internal/database"
	"users-task-project/internal/handlers"
	"users-task-project/internal/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// @title           Account Task Manager
// @version         1.0
// @description     This is an account task manager app
// @host      localhost:8080
func main() {
	cfg := config.NewConfig()

	db, err := database.InitDB(cfg)
	models.InitDB(db)

	if err != nil {
		log.Fatalf("Error - initializing database: %v", err)
		return
	}

	router := handlers.CreateRouter()
	if router == nil {
		log.Fatalf("Error - cannot create router")
	}

	handlers.Run(router)

}
