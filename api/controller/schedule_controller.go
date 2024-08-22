package controller

import (
	"api/model"
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScheduleController interface {
	CreateSchedule(c *gin.Context)
	DeleteScheduleByID(c *gin.Context)
	GetShedulesByUserID(c *gin.Context)
	JudgeScheduleByID(c *gin.Context)
}

type scheduleController struct {
	su usecase.ScheduleUsecase
}

func NewScheduleController(su usecase.ScheduleUsecase) ScheduleController {
	return &scheduleController{su}
}

func (sc *scheduleController) CreateSchedule(c *gin.Context) {
	var schedule model.Schedule
	if err := c.BindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSchedule, err := sc.su.CreateSchedule(schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSchedule)
}

func (sc *scheduleController) DeleteScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.su.DeleteScheduleByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}

func (sc *scheduleController) GetShedulesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedules, err := sc.su.GetShedulesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (sc *scheduleController) JudgeScheduleByID(c *gin.Context) {
	scheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err := sc.su.GetScheduleByID(scheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var currentLocation model.CurrentLocation
	if err := c.BindJSON(&currentLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.su.JudgeSchedule(schedule, currentLocation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule judged successfully"})
}
