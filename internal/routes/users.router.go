package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/middlewares"
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
)

func addUserRouter(router *gin.Engine, userRepo *repositories.UserRepository, mdw *middlewares.Middleware) {
	userRouter := router.Group("/users")
	userHandler := handlers.NewUserHandler(userRepo)

	// definisikan rute dengan params id
	userRouter.GET("/:id", userHandler.GetEmployeeById)

	// /users?name=nana
	userRouter.GET("", userHandler.GetUsers)

	// /users
	userRouter.POST("", mdw.VerifyToken, mdw.AccessGate("user", "admin"), userHandler.AddEmployee)

	userRouter.PATCH("", mdw.VerifyToken, mdw.AccessGate("user"), userHandler.EditStudents)
}
