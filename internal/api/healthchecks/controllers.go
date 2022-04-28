package healthchecks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func status(c *gin.Context) {
	// add controls for DB status
	// status to have: up, unstable, down
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}
