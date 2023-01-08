package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	ID          uint      `json:"id" gorm:"primary_key,auto_increment"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *Activity) CreateActivity(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Activity) UpdateActivity(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var activity Activity
	err := db.Where("id = ? AND user_id = ?", a.ID, a.UserID).First(&activity).Error
	if err != nil {
		return err
	}
	err = db.Model(&activity).Updates(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Activity) DeleteActivity(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var activity Activity
	err := db.Where("id = ? AND user_id = ?", a.ID, a.UserID).First(&activity).Error
	if err != nil {
		return err
	}
	err = db.Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Activity) GetActivities(ctx *gin.Context) ([]Activity, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var activities []Activity
	err := db.Where("user_id = ?", a.UserID).Find(&activities).Error
	if err != nil {
		return []Activity{}, err
	}
	return activities, nil
}
