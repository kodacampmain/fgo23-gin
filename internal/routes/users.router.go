package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func addUserRouter(router *gin.Engine, userRepo *repositories.UserRepository) {
	userRouter := router.Group("/users")
	userHandler := handlers.NewUserHandler(userRepo)

	// definisikan rute dengan params id
	userRouter.GET("/:id", userHandler.GetEmployeeById)

	// /users?name=nana
	userRouter.GET("", userHandler.GetUsers)

	// /users
	userRouter.POST("", userHandler.AddEmployee)

}
