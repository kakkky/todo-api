package task

type CreateTaskUsecaseInputDTO struct {
	UserId  string
	Content string
	State   string
}

type CreateTaskUsecaseOutputDTO struct {
	ID      string
	UserId  string
	Content string
	State   string
}
