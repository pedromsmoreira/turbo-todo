package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Status(c *gin.Context) {
	// add controls for DB status
	// status to have: up, unstable, down
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}
