package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pedromsmoreira/turbo-todo/internal/api/configs"
	"github.com/pedromsmoreira/turbo-todo/internal/api/schema"
)

var (
	maxRetryAttempts int = 10
)

func TestMain(m *testing.M) {
	cfg := &configs.Config{
		Database: *(&configs.Database{
			Host:     "localhost:26257",
			Username: "admin",
			Password: "password",
		}),
		Server: *(&configs.Server{
			Host: "localhost",
			Port: "5000",
		}),
		Messaging: *(&configs.Messaging{
			Host: "localhost",
			Port: "4222",
		}),
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

func checkAccountApiIsUp() error {
	attempts := 0
	url := getHost()

	for {
		resp, err := http.Get(fmt.Sprintf("%v/v1/health", url))
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Api started but not up. StatusCode: %v | Attemp: %v", resp.StatusCode, attempts)
			attempts++
			time.Sleep(5 * time.Second)
			continue
		}

		if attempts == maxRetryAttempts {
			return fmt.Errorf("reached maximum attemps: %v", maxRetryAttempts)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		h := &health{}
		if err = json.Unmarshal(body, h); err != nil {
			return err
		}

		if h.Status != "up" {
			fmt.Printf("Api started but not up. Status: %v | Attemp: %v", h.Status, attempts)
			attempts++
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}

	return nil
}

func getHost() string {
	host, exists := os.LookupEnv("HOST")
	if host == "" && !exists {
		return "http://localhost:5000"
	}

	return host
}

type health struct {
	Status string `json:"status"`
}
