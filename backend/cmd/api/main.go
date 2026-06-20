package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tuhu02/backend/internal/config"
	"github.com/tuhu02/backend/internal/database"
	"github.com/tuhu02/backend/internal/models"
)

func main() {
	cfg := config.Load()
	database.Connect(cfg)

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":" + cfg.AppPort)
}
