package todo

import "github.com/gin-gonic/gin"

type TodoHandler struct {
	Method   string
	Endpoint string
	handler  gin.HandlerFunc
}

type Handlers struct {
	Handlers []TodoHandler
}

func GetHandlers() *Handlers {
	h := make([]TodoHandler, 4)

	h = append(h, TodoHandler{
		Method:   "GET",
		Endpoint: "/v1/todos/id",
		handler:  Get,
	})

	return &Handlers{
		Handlers: h,
	}
}
