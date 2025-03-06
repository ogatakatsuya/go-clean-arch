package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	conn := db.NewDB()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(conn)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	userContoller := controller.NewUserController(userUseCase)

	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(conn)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.NewRouter(userContoller, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
