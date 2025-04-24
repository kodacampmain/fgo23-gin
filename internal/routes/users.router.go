package routes

import (
	"fgo23-gin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func addUserRouter(router *gin.Engine) {
	userRouter := router.Group("/users")
	userHandler := handlers.NewUserHandler()

	// definisikan rute dengan params id
	userRouter.GET("/:id", userHandler.GetEmployeeById)

	// /users?name=nana
	userRouter.GET("", userHandler.GetUsers)

	// /users
	userRouter.POST("", userHandler.AddEmployee)

}
