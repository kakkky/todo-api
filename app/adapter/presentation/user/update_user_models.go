package user

type UpdateUserRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
