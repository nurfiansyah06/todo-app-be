package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-app-be/helper"
	"todo-app-be/model/web"
	"todo-app-be/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}

	helper.ReadFromRequestBody(r, &userCreateRequest)

	validate := validator.New()
	err := validate.Struct(userCreateRequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			fieldName := strings.ToLower(e.Field())

			// Check specifically for the "text" field
			if fieldName == "email" {
				errorList[fieldName] = fmt.Sprintf("%s cannot be empty", fieldName)
			}
		}

		// If errorList contains "text" error, return 400 Bad Request
		if _, ok := errorList["email"]; ok {
			logrus.WithFields(
				logrus.Fields{"email": userCreateRequest.Email, "full_name": userCreateRequest.FullName, "password": userCreateRequest.Password}).Warn("Unauthorized access")
			webResponse := web.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   errorList,
			}
			helper.WriteToResponseBodyUnauthorized(w, &webResponse)
			return
		}
	}

	userResponse := controller.UserService.Create(r.Context(), userCreateRequest)
	logrus.WithFields(logrus.Fields{"email": userCreateRequest.Email, "full_name": userCreateRequest.FullName, "password": userCreateRequest.Password}).Info("Register successful")
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *UserControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.GetById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)
	userUpdateRequest.Id = id

	storyResponse := controller.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   storyResponse,
	}

	helper.WriteToResponseBody(w, &webResponse)
}
