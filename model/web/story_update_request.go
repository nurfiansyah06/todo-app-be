package web

type StoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Text string `validate:"required",max=200,min=1`
}
