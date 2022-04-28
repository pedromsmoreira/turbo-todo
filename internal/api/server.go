package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pedromsmoreira/turbo-todo/internal/api/todos"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
	"github.com/pedromsmoreira/turbo-todo/internal/api/healthchecks"
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
	err := s.Router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	healthchecks.Routes(s.Router)
	todos.Routes(s.Router)

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
