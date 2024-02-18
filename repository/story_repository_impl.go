package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-app-be/helper"
	"todo-app-be/model/domain"
)

type StoryRepositoryImpl struct {
}

func NewStoryRepository() StoryRepository {
	return &StoryRepositoryImpl{}
}

func (repository *StoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, story domain.Story) domain.Story {
	SQL := "insert into story(text) values (?)"
	result, err := tx.ExecContext(ctx, SQL, story.Text)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	story.Id = int(id)

	return story
}

func (repository *StoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, story domain.Story) domain.Story {
	SQL := "update story set text = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, story.Text, story.Id)
	helper.PanicIfError(err)

	return story
}

func (repository *StoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, story domain.Story) {
	SQL := "delete from story where id = ?"
	_, err := tx.ExecContext(ctx, SQL, story.Id)
	helper.PanicIfError(err)
}

func (repository *StoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, storyId int) (domain.Story, error) {
	SQL := "select id, text from story where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, storyId)
	helper.PanicIfError(err)
	defer rows.Close()

	story := domain.Story{}
	if rows.Next() {
		err := rows.Scan(&story.Id, &story.Text)
		helper.PanicIfError(err)
		return story, nil
	}

	return story, errors.New("story is not found")
}

func (repository *StoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Story {
	SQL := "select id, text from story"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var stories []domain.Story
	for rows.Next() {
		story := domain.Story{}
		err := rows.Scan(&story.Id, &story.Text)
		helper.PanicIfError(err)
		stories = append(stories, story)
	}
	return stories
}
