package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-app-be/helper"
	"todo-app-be/model/web"
	"todo-app-be/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

type StoryControllerImpl struct {
	StoryService service.StoryService
}

func NewStoryController(storyService service.StoryService) StoryController {
	return &StoryControllerImpl{
		StoryService: storyService,
	}
}

func (controller *StoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storyCreateRequest := web.StoryCreateRequest{}

	// Read from request body
	helper.ReadFromRequestBody(r, &storyCreateRequest)

	// Validate the populated request
	validate := validator.New()
	err := validate.Struct(storyCreateRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			fieldName := strings.ToLower(e.Field())

			// Check specifically for the "text" field
			if fieldName == "text" {
				errorList[fieldName] = fmt.Sprintf("%s cannot be empty", fieldName)
			}
		}

		// If errorList contains "text" error, return 400 Bad Request
		if _, ok := errorList["text"]; ok {
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   errorList,
			}
			helper.WriteToResponseBodyUnauthorized(w, &webResponse)
			return
		}
	}

	// If validation passes or no "text" field error, proceed to create the story
	storyResponse := controller.StoryService.Create(r.Context(), storyCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   storyResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *StoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storyUpdateRequest := web.StoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &storyUpdateRequest)

	storyId := params.ByName("storyId")
	id, err := strconv.Atoi(storyId)
	helper.PanicIfError(err)
	storyUpdateRequest.Id = id

	storyResponse := controller.StoryService.Update(r.Context(), storyUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   storyResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *StoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storyId := params.ByName("storyId")
	id, err := strconv.Atoi(storyId)
	helper.PanicIfError(err)

	controller.StoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *StoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storyId := params.ByName("storyId")
	id, err := strconv.Atoi(storyId)
	helper.PanicIfError(err)

	storyResponse := controller.StoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   storyResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *StoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	storyReponses := controller.StoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   storyReponses,
	}

	helper.WriteToResponseBody(w, &webResponse)
}
