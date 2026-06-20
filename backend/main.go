package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tuhu02/backendservices/config"
	"github.com/tuhu02/backendservices/handlers"
	"github.com/tuhu02/backendservices/models"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.User{})
	
	r := gin.Default()

	r.GET("/users" , handlers.GetUsers)

	r.Run(":8080")
}
