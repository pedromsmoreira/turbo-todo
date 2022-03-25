package todo

type TodoDto struct {
	Id          string          `json:"id"`
	Attributes  *TodoAttributes `json:"attributes"`
	DateCreated string          `json:"date_created"`
	Version     int64           `json:"version"`
}

type TodoAttributes struct {
	Title   string   `json:"title"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Status  string   `json:"status"`
}

type CreateTodoRequest struct {
	Data *TodoDto `json:"data"`
}

type UpdateTodo struct {
	Data *TodoDto `json:"data"`
}

type ApiResponse struct {
	Todo *TodoDto `json:"data"`
}

type ApiError struct {
	Errors []string `json:"errors,omitempty"`
}
