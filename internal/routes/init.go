package routes

import (
	"fgo23-gin/internal/handlers"
	"fgo23-gin/internal/middlewares"
	"fgo23-gin/internal/repositories"

	_ "fgo23-gin/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	// gin engine initialization
	router := gin.Default()
	// router := gin.New()
	pingRepo := repositories.NewPingRepository(db, rdb)
	userRepo := repositories.NewUserRepository(db)
	authRepo := repositories.NewAuthRepo(db)
	orderRepo := repositories.NewOrderRepository(db)

	middleware := middlewares.InitMiddleware()

	router.Use(middleware.CORSMiddleware)
	// router.Use(middleware.Logger)
	// router.Use(middleware.Error)

	// serve static file
	router.Static("/img", "public/img")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addPingRouter(router, pingRepo)
	addUserRouter(router, userRepo, middleware)
	addAuthRouter(router, authRepo)

	orderHandler := handlers.NewOrderHandler(orderRepo)
	router.POST("/order", middleware.VerifyToken, orderHandler.CreateTransaction)

	return router
}
