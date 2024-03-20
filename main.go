package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	funcdb "github.com/chirzul/recruitment-cudo/src/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	port := os.Getenv("GIN_PORT")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()
	r.POST("/GenerateJSONStructure", func(c *gin.Context) {
		var request struct {
			OrgID string `json:"org_id"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		org, err := funcdb.GenerateJSONStructure(request.OrgID, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get organization: %v", err)})
			return
		}

		c.JSON(http.StatusOK, org)
	})

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
