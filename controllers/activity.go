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

// CreateActivity godoc
// @Summary Create New Activity.
// @Description Creating a new Activity.
// @Tags Activity
// @Param Body body CreateActivityInput true "the body to create a new Activity"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Activity
// @Router /activity [post]
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

// GetActivities godoc
// @Summary Get all Activities from user.
// @Description Get a list of Activities.
// @Tags Activity
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Activity
// @Router /activity [get]
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

// UpdateActivity godoc
// @Summary Update Activity.
// @Description Update Activity By ID.
// @Tags Activity
// @Produce json
// @Param id path string true "Activity id"
// @Param Body body CreateActivityInput true "the body to update Activity"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200
// @Router /activity/{id} [put]
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

// DeleteActivity godoc
// @Summary Delete one Activity.
// @Description Delete a Activity by id.
// @Tags Activity
// @Produce json
// @Param id path string true "Activity id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /activity/{id} [delete]
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
