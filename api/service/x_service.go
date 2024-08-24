package service

import "api/model"

type XService struct {
}

func NewXService() XService {
	return XService{}
}

type XServiceInterface interface {
	Post(schedule model.Schedule) error
}

func (xs *XService) Post(schedule model.Schedule) error {
	// ここに処理を書く
	return nil
}
