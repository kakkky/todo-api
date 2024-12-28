package task

type GetUserTaskResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	State   string `json:"state"`
}
