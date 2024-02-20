package service

import (
	"context"
	"todo-app-be/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	GetById(ctx context.Context, userId int) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
}
