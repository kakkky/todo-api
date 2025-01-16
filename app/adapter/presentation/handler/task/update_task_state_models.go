package task

type UpdateTaskStateRequest struct {
	State string `json:"state" validate:"required"`
}
type UpdateTaskStateResponse struct {
	ID      string `json:"id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
	State   string `json:"state"`
}
