package task

type PostTaskRequest struct {
	Content string `json:"content" validate:"required"`
	State   string `json:"state" validate:"required"`
}

type PostTaskResponse struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id" `
	Content string `json:"content"`
	State   string `json:"state"`
}
