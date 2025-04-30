package routes

import (
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	pingRepo := repositories.NewPingRepository(db)
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepo(db)

	addPingRouter(router, pingRepo)
	addUserRouter(router, userRepo)
	addAuthRouter(router, authRepo)

	return router
}
