package task

type FetchLoggedInUserTasksUsecaseInputDTO struct {
	UserId string
}
type FetchLoggedInUserTasksUsecaseOutputDTO struct {
	ID      string
	Content string
	State   string
}
