package model

import "time"

type Schedule struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Deadline  time.Time `json:"deadline"`
	User      User      `json:"user" gorm:"foreignkey:UserID"`
	UserID    int       `json:"user_id" gorm:"not null"`
}
