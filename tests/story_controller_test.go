package tests

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"todo-app-be/app"
	"todo-app-be/controller"
	"todo-app-be/helper"
	"todo-app-be/middleware"
	"todo-app-be/repository"
	"todo-app-be/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo_app_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	storyRepository := repository.NewStoryRepository()
	storyService := service.NewStoryService(storyRepository, db, validate)
	storyController := controller.NewStoryController(storyService)
	router := app.NewRouter(storyController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateStorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"text": "sehat bang?"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/stories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCreateStoryFailed(t *testing.T) {

}

func TestUpdateStorySuccess(t *testing.T) {

}

func TestUpdateStoryFailed(t *testing.T) {

}

func TestGetStorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/stories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
func TestGetStoryFailed(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/stories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}
func TestDeleteStorySuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/stories/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestDeleteStoryFailed(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/stories/100", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
func TestListStoriesSuccess(t *testing.T) {

}
