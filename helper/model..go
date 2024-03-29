package helper

import (
	"todo-app-be/model/domain"
	"todo-app-be/model/web"
)

func ToStoryResponse(story domain.Story) web.StoryResponse {
	return web.StoryResponse{
		Id:   story.Id,
		Text: story.Text,
	}
}

func ToStoryResponses(stories []domain.Story) []web.StoryResponse {
	var storyResponses []web.StoryResponse
	for _, story := range stories {
		storyResponses = append(storyResponses, ToStoryResponse(story))
	}

	return storyResponses
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
		Password: user.Password,
	}
}
