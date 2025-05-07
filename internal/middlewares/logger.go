package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Logger(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	elapsedTime := time.Since(start)
	log.Println(elapsedTime)
}
