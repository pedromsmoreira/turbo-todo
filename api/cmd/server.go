package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/api/configs"
	"github.com/pedromsmoreira/turbo-todo/api/internal/healthcheck"
	"github.com/pedromsmoreira/turbo-todo/api/internal/todo"
	"github.com/pedromsmoreira/turbo-todo/api/schema"
)

func main() {
	cfg := configs.NewConfig()
	err := schema.CreateSchema(cfg)
	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/v1/ping", healthcheck.Ping)

	todorepo := todo.NewInMemoryTodoRepository()
	todosvc := todo.NewTodoService(todorepo)
	tc := todo.NewTodoController(todosvc)

	v1 := r.Group("/v1")
	{
		v1.GET("/todos", tc.List)
		v1.GET("/todos/:id", tc.Get)
		v1.POST("/todos", tc.Create)
		v1.PUT("/todos/:id", tc.Update)
		v1.DELETE("/todos/:id", tc.Delete)
	}

	err = r.Run(fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port))

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
