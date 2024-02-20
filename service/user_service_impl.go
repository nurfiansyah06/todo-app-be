package service

import (
	"context"
	"database/sql"
	"todo-app-be/helper"
	"todo-app-be/model/domain"
	"todo-app-be/model/web"
	"todo-app-be/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hash, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.User{
		Email:    request.Email,
		FullName: request.FullName,
		Password: hash,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) GetById(ctx context.Context, userId int) web.UserResponse {
	// tx, err := service.DB.Begin()
	// helper.PanicIfError(err)

	// defer helper.CommitOrRollback(tx)

	// story, err := service.StoryRepository.FindById(ctx, tx, id)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }

	// return helper.ToStoryResponse(story)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetById(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Email = request.Email
	user.FullName = request.FullName
	user.Password, _ = helper.HashPassword(request.Password)

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}
