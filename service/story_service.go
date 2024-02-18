package service

import (
	"context"
	"todo-app-be/model/web"
)

type StoryService interface {
	Create(ctx context.Context, request web.StoryCreateRequest) web.StoryResponse
	Update(ctx context.Context, request web.StoryUpdateRequest) web.StoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.StoryResponse
	FindAll(ctx context.Context) []web.StoryResponse
}
