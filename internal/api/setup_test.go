package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
	"github.com/pedromsmoreira/turbo-todo/internal/api/schema"
)

func TestMain(m *testing.M) {
	cfg := &configs.Config{
		Database: configs.Database{
			Host:     "localhost:26257",
			Username: "admin",
			Password: "password",
		},
		Server: configs.Server{
			Host: "localhost",
			Port: "5000",
		},
		Messaging: configs.Messaging{
			Host: "localhost",
			Port: "4222",
		},
	}
	err := schema.CreateSchema(cfg)
	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}

	server := NewServer(cfg)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	code := m.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(code)
}
