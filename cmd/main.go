package main

import (
	"fgo23-gin/internal/routes"
	"fgo23-gin/pkg"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	if err := pkg.Connect(); err != nil {
		log.Printf("[ERROR] Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// graceful shutdown
	// server jalan di goroutine
	// goroutine main handling shutdown
	defer func() {
		log.Println("Closing DB...")
		pkg.DB.Close()
	}()

	router := routes.InitRouter()
	// jalankan service
	router.Run("127.0.0.1:8080")
	// router.Run(":8080")
}
