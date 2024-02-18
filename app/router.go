package app

import (
	"todo-app-be/controller"
	"todo-app-be/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(storyController controller.StoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/stories", storyController.FindAll)
	router.GET("/api/stories/:storyId", storyController.FindById)
	router.POST("/api/stories", storyController.Create)
	router.PUT("/api/stories/:storyId", storyController.Update)
	router.DELETE("/api/stories/:storyId", storyController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
