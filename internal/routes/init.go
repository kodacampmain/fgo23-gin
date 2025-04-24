package routes

import (
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin engine initialization
	router := gin.Default()

	repositories.NewPingRepository()
	repositories.NewUserRepository()

	addPingRouter(router)
	addUserRouter(router)

	return router
}
