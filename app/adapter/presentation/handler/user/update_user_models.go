package user

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
