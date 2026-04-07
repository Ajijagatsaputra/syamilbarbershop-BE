package main

import (
	"log"
	"os"
	"syamilbarbershop-be/internal/config"
	"syamilbarbershop-be/internal/database"
	"syamilbarbershop-be/internal/routes"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()
	database.SeedData()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
