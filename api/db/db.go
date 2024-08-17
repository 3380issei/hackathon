package db

import (
	"fmt"

	"api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Schedule{})
	fmt.Println("DB connection successfully opened")
	return db, nil
}
