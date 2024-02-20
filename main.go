package main

import (
	"log"
	"net/http"
	"todo-app-be/app"
	"todo-app-be/controller"
	"todo-app-be/helper"
	"todo-app-be/middleware"
	"todo-app-be/repository"
	"todo-app-be/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.DBConnect()
	validate := validator.New()

	storyRepository := repository.NewStoryRepository()
	storyService := service.NewStoryService(storyRepository, db, validate)
	storyController := controller.NewStoryController(storyService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := app.NewRouter(storyController, userController)
	log.Println("Application is running at :3000")
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
