package repository

import (
	"api/model"
	"errors"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(schedule *model.Schedule) error
	DeleteScheduleByID(scheduleID int) error
	GetShedulesByUserID(userID int) ([]model.Schedule, error)
	GetScheduleByID(scheduleID int) (model.Schedule, error)
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

func (sr *scheduleRepository) GetShedulesByUserID(userID int) ([]model.Schedule, error) {
	var schedules []model.Schedule
	if err := sr.db.Where("user_id = ?", userID).Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (sr *scheduleRepository) GetScheduleByID(scheduleID int) (model.Schedule, error) {
	var schedule model.Schedule
	if err := sr.db.Where("id = ?", scheduleID).First(&schedule).Error; err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}
