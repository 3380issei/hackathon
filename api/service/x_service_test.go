package service

import (
	"api/model"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	schedule := model.Schedule{
		ID:          1,
		Destination: "Tokyo",
		Latitude:    35.6895,
		Longitude:   139.6917,
		Deadline:    time.Now(),
		Expired:     false,
		UserID:      1,
	}

	xService := NewXService()
	err := xService.Post(schedule)
	if err != nil {
		t.Errorf("Post() error = %v, want nil", err)
	}
}
