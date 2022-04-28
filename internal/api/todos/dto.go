package todos

type Dto struct {
	Id          string      `json:"id"`
	Attributes  *Attributes `json:"attributes"`
	DateCreated string      `json:"date_created"`
	Version     int64       `json:"version"`
}

type Attributes struct {
	Title   string   `json:"title"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Status  string   `json:"status"`
}

type CreateTodoRequest struct {
	Data *Dto `json:"data"`
}

type UpdateTodo struct {
	Data *Dto `json:"data"`
}

type ApiResponse struct {
	Data *Dto `json:"data"`
}
