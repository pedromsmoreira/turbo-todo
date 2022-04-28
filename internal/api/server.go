package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pedromsmoreira/turbo-todo/internal/api/todo"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
	"github.com/pedromsmoreira/turbo-todo/internal/api/healthcheck"
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

	healthcheck.Routes(s.Router)
	todo.Routes(s.Router)

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
