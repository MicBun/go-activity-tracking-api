package config

import (
	"github.com/MicBun/go-activity-tracking-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	//username := utils.GetEnv("DB_USERNAME", "myuser")
	//password := utils.GetEnv("DB_PASSWORD", "mypassword")
	//host := utils.GetEnv("DB_HOST", "tcp(mysql:3306)")
	//database := utils.GetEnv("DB_DATABASE", "activity_tracking")
	//dsn := username + ":" + password + "@" + host + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//dsn := "host=dpg-cetjjlpgp3jmgl18q2rg-a user=myuser password=yN5Do2dZnQivCkSqulOK9mwWa6VPB5Dx dbname=activity_tracking port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbUrl := "postgres://myuser:yN5Do2dZnQivCkSqulOK9mwWa6VPB5Dx@dpg-cetjjlpgp3jmgl18q2rg-a/activity_tracking"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Attendance{}, &models.Activity{})
	return db
}
