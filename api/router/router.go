package router

import (
	"api/controller"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.UserController, sc controller.ScheduleController, am middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.Use(am.CORS())

	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)
	r.GET("/users/:id", uc.GetUserByID)

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.POST("", sc.CreateSchedule)
		scheduleGroup.DELETE("/:id", sc.DeleteScheduleByID)
		scheduleGroup.GET("/:id", sc.GetSheduleByID)
		scheduleGroup.GET("users/:user_id", sc.GetShedulesByUserID)
		scheduleGroup.POST("/:id", sc.JudgeScheduleByID)
	}

	return r
}
