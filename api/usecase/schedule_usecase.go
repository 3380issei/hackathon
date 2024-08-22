package usecase

import (
	"api/model"
	"api/repository"
)

type ScheduleUsecase interface {
	CreateSchedule(schedule model.Schedule) (model.Schedule, error)
	DeleteScheduleByID(scheduleID int) error
	GetShedulesByUserID(userID int) ([]model.Schedule, error)
}

type scheduleUsecase struct {
	sr repository.ScheduleRepository
}

func NewScheduleUsecase(sr repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{sr}
}

func (su *scheduleUsecase) CreateSchedule(schedule model.Schedule) (model.Schedule, error) {
	if err := su.sr.CreateSchedule(&schedule); err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (su *scheduleUsecase) DeleteScheduleByID(scheduleID int) error {
	if err := su.sr.DeleteScheduleByID(scheduleID); err != nil {
		return err
	}
	return nil
}

func (su *scheduleUsecase) GetShedulesByUserID(userID int) ([]model.Schedule, error) {
	schedules, err := su.sr.GetShedulesByUserID(userID)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}
