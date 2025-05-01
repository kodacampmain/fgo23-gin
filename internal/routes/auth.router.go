package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func addAuthRouter(router *gin.Engine, authRepo *repositories.AuthRepo) {
	authRouter := router.Group("/auth")
	authHandler := handlers.NewAuthHandler(authRepo)
	// endpoint & resource
	// /ping => protocol://hostname/ping => http://localhost:port/ping
	authRouter.POST("", authHandler.Login)
	authRouter.POST("/new", authHandler.Register)
	authRouter.GET("/verify", authHandler.VerifyToken)
}
