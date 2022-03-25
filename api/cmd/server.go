package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/api/internal/healthcheck"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/v1/ping", healthcheck.Ping)

	err := r.Run()

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
