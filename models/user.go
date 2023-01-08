package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User
type User struct {
	ID         uint         `json:"id" gorm:"primary_key"`
	Username   string       `json:"username" gorm:"type:varchar(255)"`
	Password   string       `json:"password" gorm:"type:varchar(255)"`
	Activity   []Activity   `json:"activity" gorm:"foreignKey:UserID"`
	Attendance []Attendance `json:"attendance" gorm:"foreignKey:UserID"`
}

func (u *User) LoginUser(ctx *gin.Context) (User, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
