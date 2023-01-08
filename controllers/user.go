package controllers

import (
	"github.com/MicBun/go-activity-tracking-api/models"
	"github.com/MicBun/go-activity-tracking-api/utils/jwtAuth"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userPassword := models.User{Username: input.Username, Password: input.Password}
	user, err := userPassword.LoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "username or password is incorrect."})
		return
	}
	jwtToken, err := jwtAuth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": jwtToken})
}
