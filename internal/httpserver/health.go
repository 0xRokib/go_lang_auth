package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func checkHealth(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"service": "auth-service",
		"time":    time.Now().UTC(),
	})

}
