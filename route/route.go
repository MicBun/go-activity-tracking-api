package route

import (
	"github.com/MicBun/go-activity-tracking-api/controllers"
	"github.com/MicBun/go-activity-tracking-api/middleware"
	"github.com/MicBun/go-activity-tracking-api/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/login", controllers.Login)
	r.POST("/reset", func(context *gin.Context) {
		if err := utils.ResetTable(db); err != nil {
			utils.HandleError(context, err)
		}
		context.JSON(200, gin.H{"message": "reset success"})
	})

	clockIn := r.Group("/clock-in")
	clockIn.Use(middleware.JwtAuthMiddleware())
	clockIn.POST("", controllers.ClockIn)

	clockOut := r.Group("/clock-out")
	clockOut.Use(middleware.JwtAndClockInMiddleware())
	clockOut.POST("", controllers.ClockOut)

	attendance := r.Group("/attendance")
	attendance.Use(middleware.JwtAndClockInMiddleware())
	attendance.GET("", controllers.GetAttendances)
	attendance.GET("/by-date", controllers.GetAttendanceByDate)

	activity := r.Group("/activity")
	activity.Use(middleware.JwtAndClockInMiddleware())
	activity.POST("", controllers.CreateActivity)
	activity.GET("", controllers.GetActivities)
	activity.PUT("/:id", controllers.UpdateActivity)
	activity.DELETE("/:id", controllers.DeleteActivity)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
