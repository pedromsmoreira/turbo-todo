package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	maxRetryAttempts int = 10
)

func TestMain(m *testing.M) {
	err := checkAccountApiIsUp()

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestServer(t *testing.T) {
	t.Run("when service is up and running ping returns 200 OK with pong message", func(t *testing.T) {

	})
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
