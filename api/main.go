package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/router"
	"api/usecase"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	router := router.NewRouter(userController)

	router.Run(":8080")
}
