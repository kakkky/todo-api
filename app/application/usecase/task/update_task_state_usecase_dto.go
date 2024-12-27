package task

type UpdateTaskStateUsecaseInputDTO struct {
	ID             string
	LoggedInUserID string
	State          string
}

type UpdateTaskStateUsecaseOutputDTO struct {
	ID      string
	UserId  string
	Content string
	State   string
}
