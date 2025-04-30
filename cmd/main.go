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
	// pg15, err := pkg.ConnectPg15()
	// if err != nil {
	// 	log.Printf("[ERROR] Unable to create connection: %v\n", err)
	// 	os.Exit(1)
	// }
	// graceful shutdown
	// server jalan di goroutine
	// goroutine main handling shutdown
	defer func() {
		log.Println("Closing DB...")
		pg.Close()
	}()

	// var hash pkg.HashConfig
	// hash.UseDefaultConfig()
	// password := "fazztrack"
	// hashedPassword, _ := hash.GenHashedPassword(password)
	// log.Println("[DEBUG] password: ", password)
	// log.Println("[DEBUG] hash: ", hashedPassword)

	router := routes.InitRouter(pg)
	// jalankan service
	router.Run("127.0.0.1:8080")
	// router.Run(":8080")
}
