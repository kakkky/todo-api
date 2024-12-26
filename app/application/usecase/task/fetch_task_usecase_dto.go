package task

type FetchTaskUsecaseInputDTO struct {
	ID string
}

type FetchTaskUsecaseOutputDTO struct {
	ID       string
	UserName string
	UserId   string
	Content  string
	State    string
}
