package task

type GetTaskResponse struct {
	ID       string `json:"id"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Content  string `json:"content"`
	State    string `json:"state"`
}
