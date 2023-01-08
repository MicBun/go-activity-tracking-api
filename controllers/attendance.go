package controllers

import (
	"github.com/MicBun/go-activity-tracking-api/models"
	"github.com/MicBun/go-activity-tracking-api/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"time"
)

func ClockIn(c *gin.Context) {
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	attendance := models.Attendance{UserID: userID}
	err = attendance.ClockIn(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "you have already clocked in"})
		return
	}
	c.JSON(200, gin.H{"message": "clock in success"})
}

func ClockOut(c *gin.Context) {
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	attendance := models.Attendance{UserID: userID}
	err = attendance.ClockOut(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "record not found or already clocked out"})
		return
	}

	c.JSON(200, gin.H{"message": "clock out success"})
}

type GetAttendanceOutput struct {
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
}

func GetAttendances(c *gin.Context) {
	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	attendance := models.Attendance{UserID: userID}
	attendanceList, err := attendance.GetAttendances(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var output []GetAttendanceOutput
	for _, attendance := range attendanceList {
		output = append(output, GetAttendanceOutput{
			CheckIn:  attendance.CreatedAt,
			CheckOut: attendance.UpdatedAt,
		})
	}
	
	c.JSON(200, gin.H{"data": output})
}

type GetAttendanceByDateInput struct {
	Date string `json:"date" binding:"required"`
}

type GetAttendanceByDateOutput struct {
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
}

func GetAttendanceByDate(c *gin.Context) {
	var input GetAttendanceByDateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	attendance := models.Attendance{UserID: userID}
	attendanceResult, err := attendance.GetAttendanceByDate(c, input.Date)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": GetAttendanceByDateOutput{
		CheckIn:  attendanceResult.CreatedAt,
		CheckOut: attendanceResult.UpdatedAt,
	}})
}
