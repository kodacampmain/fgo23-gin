package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func addPingRouter(router *gin.Engine, pingRepo *repositories.PingRepository) {
	pingRouter := router.Group("/ping")
	pingHandler := handlers.NewPingHandler(pingRepo)
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	pingRouter.GET("", pingHandler.GetStudents)
}
