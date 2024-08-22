package usecase

import (
	"api/model"
	"api/repository"
	"errors"
	"math"
)

type ScheduleUsecase interface {
	CreateSchedule(schedule model.Schedule) (model.Schedule, error)
	DeleteScheduleByID(scheduleID int) error
	GetScheduleByID(scheduleID int) (model.Schedule, error)
	GetShedulesByUserID(userID int) ([]model.Schedule, error)
	JudgeSchedule(schedule model.Schedule, currentLocation model.CurrentLocation) error
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

func (su *scheduleUsecase) GetScheduleByID(scheduleID int) (model.Schedule, error) {
	schedule, err := su.sr.GetScheduleByID(scheduleID)
	if err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (su *scheduleUsecase) GetShedulesByUserID(userID int) ([]model.Schedule, error) {
	schedules, err := su.sr.GetShedulesByUserID(userID)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (su *scheduleUsecase) JudgeSchedule(schedule model.Schedule, currentLocation model.CurrentLocation) error {
	latitude := currentLocation.Latitude
	longitude := currentLocation.Longitude
	currentTime := currentLocation.CurrentTime

	// 期限切れ判定
	if currentTime.After(schedule.Deadline) {
		return errors.New("the deadline has passed")
	}

	// 位置情報判定
	if !IsWithinRadius(schedule.Latitude, schedule.Longitude, latitude, longitude) {
		return errors.New("you are not within the radius")
	}

	if err := su.sr.DeleteScheduleByID(schedule.ID); err != nil {
		return err
	}

	return nil
}

const radius = 100.0

// 半径100m以内にいるかどうかを判定
func IsWithinRadius(destLat float64, destLon float64, curLat float64, curLon float64) bool {
	// 二点間の距離を計算
	distance := haversine(destLat, destLon, curLat, curLon)

	// 半径100m以内ならtrueを返す
	return distance <= radius
}

// Haversine公式を使って、緯度・経度の2点間の距離を計算
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // 地球の半径（メートル）

	// 度をラジアンに変換
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Haversineの公式
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// 距離を計算
	return R * c
}
