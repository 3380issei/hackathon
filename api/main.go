package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/router"
	"api/usecase"
	"os"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	scheduleRepository := repository.NewScheduleRepository(db)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepository)
	scheduleController := controller.NewScheduleController(scheduleUsecase)

	router := router.NewRouter(userController, scheduleController)
	router.Run(os.Getenv("API_ADDRESS"))
}
