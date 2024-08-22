package usecase

import (
	"api/model"
	"api/repository"
)

type ScheduleUsecase interface {
	CreateSchedule(schedule model.Schedule) (model.Schedule, error)
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
