package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "users-task-project/docs"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	if err := router.SetTrustedProxies([]string{"192.168.1.2", "10.0.0.0/8"}); err != nil {
		return nil
	}

	router.Static("/assets/", "frontend/")
	router.LoadHTMLGlob("frontend/templates/*.html")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", indexHandler)
	router.GET("/reg", registerHTMLHandler)
	router.GET("/upd", updateHTMLHandler)
	router.GET("/del", deleteHTMLHandler)
	router.GET("/man/:Passport", manageHTMLHandler)
	router.GET("/info/:PassportSerie/:PassportNumber", infoUserHandler)
	router.POST("/create/:Passport", createUserHandler)
	router.POST("/man/begin/:Passport/:TaskName", beginTaskHandler)
	router.POST("/man/end/:Passport/:TaskID", endTaskHandler)
	router.GET("/filter", filterHandler)
	router.PUT("/update/:Passport", updateUserHandler)
	router.DELETE("/delete/:Passport", deleteUserHandler)

	return router
}

func Run(router *gin.Engine) {
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error - starting HTTP server: %v", err)
		return
	}
}
