package usecase

import (
	"api/model"
	"api/repository"
)

type ScheduleUsecase interface {
	CreateSchedule(schedule model.Schedule) (model.Schedule, error)
	DeleteScheduleByID(scheduleID int) error
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
