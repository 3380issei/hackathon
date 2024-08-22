package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.UserController, sc controller.ScheduleController) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.POST("", sc.CreateSchedule)
		scheduleGroup.DELETE("/:id", sc.DeleteScheduleByID)
		scheduleGroup.GET("/:user_id", sc.GetShedulesByUserID)
		scheduleGroup.POST("/:id", sc.JudgeScheduleByID)
	}

	return r
}
