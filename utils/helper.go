package utils

import (
	"fmt"
	"github.com/MicBun/go-activity-tracking-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func ResetTable(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&models.User{}, &models.Activity{}, &models.Attendance{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Activity{}, &models.Attendance{}); err != nil {
		return err
	}

	for i := 1; i <= 10; i++ {
		username := fmt.Sprintf("user%d", i)
		password := fmt.Sprintf("password%d", i)
		if err := db.Create(&models.User{Username: username, Password: password}).Error; err != nil {
			return err
		}
	}
	return nil
}
