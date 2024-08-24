package service

import (
	"api/model"
	"log"
)

type XService struct {
}

func NewXService() XService {
	return XService{}
}

type XServiceInterface interface {
	Post(schedule model.Schedule) error
}

func (xs *XService) Post(schedule model.Schedule) error {
	log.Println("Post to XService")
	return nil
}

func generateContent(user *User) string {
	log.Println("Huh?")
	return "Huh?"
}
