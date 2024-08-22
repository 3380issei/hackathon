package repository

import (
	"api/model"
	"errors"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *model.Schedule) error
	DeleteScheduleByID(scheduleID int) error
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

func (sr *scheduleRepository) DeleteScheduleByID(scheduleID int) error {
	result := sr.db.Where("id = ?", scheduleID).Delete(&model.Schedule{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("schedule not found")
	}
	return nil
}
