package auth

type LoginUsecaseInputDTO struct {
	Email    string
	Password string
}

type LoginUsecaseOutputDTO struct {
	SignedToken string
}
