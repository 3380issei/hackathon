package model

import "time"

type Schedule struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Destination string    `json:"destination"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Deadline    time.Time `json:"deadline"`
	Expired     bool      `json:"expired" default:"false"`
	User        User      `json:"user" gorm:"foreignkey:UserID"`
	UserID      int       `json:"user_id" gorm:"not null"`
}

type CurrentLocation struct {
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CurrentTime time.Time `json:"current_time"`
}
