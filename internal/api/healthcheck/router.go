package healthcheck

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.GET("/ping", ping)
	r.GET("/status", status)
}
