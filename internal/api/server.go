package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
	"github.com/pedromsmoreira/turbo-todo/internal/api/healthcheck"
	"github.com/pedromsmoreira/turbo-todo/internal/api/todo"
)

type Server struct {
	Cfg    *configs.Config
	Router *gin.Engine
	server *http.Server
}

func NewServer(cfg *configs.Config) *Server {
	return &Server{
		Cfg:    cfg,
		Router: gin.Default(),
	}
}

func (s *Server) Start() error {
	s.Router.SetTrustedProxies(nil)
	s.Router.GET("/ping", healthcheck.Ping)
	s.Router.GET("/status", healthcheck.Status)

	todorepo := todo.NewInMemoryTodoRepository()
	todosvc := todo.NewTodoService(todorepo)
	tc := todo.NewTodoController(todosvc)

	v1 := s.Router.Group("/v1")
	{
		v1.GET("/todos", tc.List)
		v1.GET("/todos/:id", tc.Get)
		v1.POST("/todos", tc.Create)
		v1.PUT("/todos/:id", tc.Update)
		v1.DELETE("/todos/:id", tc.Delete)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Cfg.Server.Host, s.Cfg.Server.Port),
		Handler: s.Router,
	}

	s.server = srv

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
