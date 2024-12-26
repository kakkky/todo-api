package task

type FetchUserTasksUsecaseInputDTO struct {
	UserId string
}
type FetchUserTasksUsecaseOutputDTO struct {
	ID      string
	Content string
	State   string
}
