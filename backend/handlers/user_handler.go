package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhu02/backendservices/config"
	"github.com/tuhu02/backendservices/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}
