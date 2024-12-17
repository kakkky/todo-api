package user

type RegisterUsecaseInputDTO struct {
	Name     string
	Email    string
	Password string
}

type RegisterUsecaseOutputDTO struct {
	ID    string
	Name  string
	Email string
}
