package db

import (
	"fmt"
	"os"

	"api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	var dsn string

	env := os.Getenv("DB_ENV")
	switch env {
	case "test":
		dsn = "test.db"
	case "prod":
		dsn = "prod.db"
	default:
		dsn = "dev.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Schedule{})
	fmt.Println("DB connection successfully opened for", env, "environment")

	return db, nil
}
