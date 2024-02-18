package repository

import (
	"context"
	"database/sql"
	"todo-app-be/model/domain"
)

type StoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, story domain.Story) domain.Story
	Update(ctx context.Context, tx *sql.Tx, story domain.Story) domain.Story
	Delete(ctx context.Context, tx *sql.Tx, story domain.Story)
	FindById(ctx context.Context, tx *sql.Tx, storyId int) (domain.Story, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Story
}
