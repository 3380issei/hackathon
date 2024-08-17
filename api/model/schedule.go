package model

import "time"

type Schedule struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserID    int       `json:"user_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Deadline  time.Time `json:"deadline"`
}
