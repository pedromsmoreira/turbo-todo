package todos

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	todoRepo := NewInMemoryTodoRepository()
	todoSvc := NewTodoService(todoRepo)
	tc := NewTodoController(todoSvc)

	r.GET("/todos", tc.list)
	r.GET("/todos/:id", tc.get)
	r.POST("/todos", tc.create)
	r.PUT("/todos/:id", tc.update)
	r.DELETE("/todos/:id", tc.delete)
}
