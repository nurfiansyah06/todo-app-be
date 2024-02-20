package web

type UserCreateRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"email,omitempty" structs:"email,omitempty"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
