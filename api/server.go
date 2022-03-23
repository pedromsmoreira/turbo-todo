package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedromsmoreira/turbo-todo/api/controllers"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/v1/ping", controllers.Ping)

	r.Run()
}
