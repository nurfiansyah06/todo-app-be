package web

type UserUpdateRequest struct {
	Id       int    `validate:"required"`
	Email    string `json:"email" validate:"email,omitempty" structs:"email,omitempty"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
