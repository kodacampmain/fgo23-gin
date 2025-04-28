package routes

import (
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

func InitRouter(db *pgxpool.Pool, db15 *sqlx.DB) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	pingRepo := repositories.NewPingRepository(db)
	userRepo := repositories.NewUserRepository(db)

	addPingRouter(router, pingRepo)
	addUserRouter(router, userRepo)

	return router
}
