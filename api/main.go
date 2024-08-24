package main

import (
	"api/controller"
	"api/db"
	"api/middleware"
	"api/repository"
	"api/router"
	"api/service"
	"api/usecase"
	"log"
	"os"
	"time"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	xService := service.NewXService()
	scheduleRepository := repository.NewScheduleRepository(db)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepository, xService)
	scheduleController := controller.NewScheduleController(scheduleUsecase)

	authMiddleware := middleware.NewAuthMiddleware()

	router := router.NewRouter(userController, scheduleController, authMiddleware)
	go func() {
		if err := router.Run(os.Getenv("API_ADDRESS")); err != nil {
			log.Fatalf("APIサーバーの起動に失敗しました: %v", err)
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Minute) // 1分ごとに実行
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := scheduleUsecase.ExecuteExpiredSchedules()
				if err != nil {
					log.Printf("定期処理でエラーが発生しました: %v", err)
				}
			}
		}
	}()

	// メインスレッドをブロック
	select {}
}
