package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/middlewares"
	"fgo23-gin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	pingRepo := repositories.NewPingRepository(db, rdb)
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepo(db)
	orderRepo := repositories.NewOrderRepository(db)

	middleware := middlewares.InitMiddleware()

	router.Use(middleware.CORSMiddleware)

	// serve static file
	router.Static("/img", "public/img")

	addPingRouter(router, pingRepo)
	addUserRouter(router, userRepo, middleware)
	addAuthRouter(router, authRepo)

	orderHandler := handlers.NewOrderHandler(orderRepo)
	router.POST("/order", middleware.VerifyToken, orderHandler.CreateTransaction)

	return router
}
