package web

type StoryCreateRequest struct {
	Text string `json:"text" validate:"required"`
}
