package user

type PostUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `password:"password" validate:"required"`
}

type PostUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
