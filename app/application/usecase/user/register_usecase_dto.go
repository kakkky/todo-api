package user

type RegisterUsecaseInputDTO struct {
	Name     string
	Email    string
	Password string
}

type RegisterUsecaseOutputDTO struct {
	Name  string
	Email string
}
