package routes

import (
	"fgo23-gin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func addPingRouter(router *gin.Engine) {
	pingRouter := router.Group("/ping")
	pingHandler := handlers.NewPingHandler()
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	pingRouter.GET("", pingHandler.GetStudents)
}
