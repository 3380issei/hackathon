package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *model.Schedule) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (sr *scheduleRepository) CreateSchedule(schedule *model.Schedule) error {
	if err := sr.db.Create(schedule).Error; err != nil {
		return err
	}
	return nil
}
