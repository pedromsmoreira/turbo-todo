package todos

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
)

func Routes(r *gin.Engine, cfg *configs.Config) {
	todoRepo, _ := NewCockroachDbTodoRepository(cfg)
	todoSvc := NewTodoService(todoRepo)
	tc := NewTodoController(todoSvc)

	r.GET("/todos", tc.list)
	r.GET("/todos/:id", tc.get)
	r.POST("/todos", tc.create)
	r.PUT("/todos/:id", tc.update)
	r.DELETE("/todos/:id", tc.delete)
}
