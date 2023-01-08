package config

import (
	"github.com/MicBun/go-activity-tracking-api/models"
	"github.com/MicBun/go-activity-tracking-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.GetEnv("DB_USERNAME", "myuser")
	password := utils.GetEnv("DB_PASSWORD", "mypassword")
	host := utils.GetEnv("DB_HOST", "mysql:3306")
	database := utils.GetEnv("DB_DATABASE", "activity_tracking")
	dsn := username + ":" + password + "@" + host + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Attendance{}, &models.Activity{})
	return db
}
