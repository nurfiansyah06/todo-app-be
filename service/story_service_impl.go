package service

import (
	"context"
	"database/sql"
	"todo-app-be/exception"
	"todo-app-be/helper"
	"todo-app-be/model/domain"
	"todo-app-be/model/web"
	"todo-app-be/repository"

	"github.com/go-playground/validator/v10"
)

type StoryServiceImpl struct {
	StoryRepository repository.StoryRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewStoryService(storyRepository repository.StoryRepository, DB *sql.DB, validate *validator.Validate) StoryService {
	return &StoryServiceImpl{
		StoryRepository: storyRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *StoryServiceImpl) Create(ctx context.Context, request web.StoryCreateRequest) web.StoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	story := domain.Story{
		Text: request.Text,
	}

	story = service.StoryRepository.Save(ctx, tx, story)

	return helper.ToStoryResponse(story)
}

func (service *StoryServiceImpl) Update(ctx context.Context, request web.StoryUpdateRequest) web.StoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	story, err := service.StoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	story.Text = request.Text

	story = service.StoryRepository.Update(ctx, tx, story)

	return helper.ToStoryResponse(story)
}

func (service *StoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	story, err := service.StoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	helper.PanicIfError(err)

	service.StoryRepository.Delete(ctx, tx, story)
}

func (service *StoryServiceImpl) FindById(ctx context.Context, id int) web.StoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	story, err := service.StoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStoryResponse(story)
}

func (service *StoryServiceImpl) FindAll(ctx context.Context) []web.StoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	stories := service.StoryRepository.FindAll(ctx, tx)

	return helper.ToStoryResponses(stories)
}
