package main

import (
	"fgo23-gin/internal/routes"
	"fgo23-gin/pkg"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	pg, err := pkg.Connect()
	if err != nil {
		log.Printf("[ERROR] Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// graceful shutdown
	// server jalan di goroutine
	// goroutine main handling shutdown
	defer func() {
		log.Println("Closing DB...")
		pg.Close()
	}()

	router := routes.InitRouter(pg)
	// jalankan service
	router.Run("127.0.0.1:8080")
	// router.Run(":8080")
}
