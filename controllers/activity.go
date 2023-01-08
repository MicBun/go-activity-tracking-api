package controllers

import (
	"github.com/MicBun/go-activity-tracking-api/models"
	"github.com/MicBun/go-activity-tracking-api/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateActivityInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateActivity(c *gin.Context) {
	var input CreateActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	activity := models.Activity{Name: input.Name, Description: input.Description, UserID: userID}
	err = activity.CreateActivity(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"data": activity})
}

func GetActivities(c *gin.Context) {
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	activities := models.Activity{UserID: userID}
	activitiesList, err := activities.GetActivities(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": activitiesList})
}

func UpdateActivity(c *gin.Context) {
	var input CreateActivityInput
	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	activity := models.Activity{ID: uint(activityID), Name: input.Name, Description: input.Description, UserID: userID}
	err = activity.UpdateActivity(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": activity})

}

func DeleteActivity(c *gin.Context) {
	activityID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	activity := models.Activity{ID: uint(activityID), UserID: userID}
	err = activity.DeleteActivity(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Activity deleted successfully."})
}
