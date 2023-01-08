package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// Attendance
type Attendance struct {
	ID        uint      `json:"id" gorm:"primary_key,auto_increment"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *Attendance) ClockIn(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)

	var attendance Attendance
	err := db.Where("user_id = ? AND DATE(created_at) = ?", a.UserID, time.Now().Format("2006-01-02")).First(&attendance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = db.Create(&a).Error
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("you have clocked in")
}

func (a *Attendance) ClockOut(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)

	var attendance Attendance
	err := db.Where("user_id = ? AND DATE(created_at) = ?", a.UserID, time.Now().Format("2006-01-02")).First(&attendance).Error
	if err != nil {
		return err
	}

	if attendance.CreatedAt == attendance.UpdatedAt {
		err = db.Model(&attendance).Update("updated_at", time.Now()).Error
		if err != nil {
			return err
		}
		return nil
	}

	return err
}

func (a *Attendance) GetAttendances(ctx *gin.Context) ([]Attendance, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var attendance []Attendance
	err := db.Where("user_id = ?", a.UserID).Find(&attendance).Error
	if err != nil {
		return []Attendance{}, err
	}
	return attendance, nil
}

func (a *Attendance) GetAttendanceByDate(ctx *gin.Context, date string) (Attendance, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var attendance Attendance
	err := db.Where("user_id = ? AND DATE(created_at) = ?", a.UserID, date).First(&attendance).Error
	if err != nil {
		return Attendance{}, err
	}
	return attendance, nil
}
