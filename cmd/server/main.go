package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goproject/internal/database/db"
	"goproject/internal/routes"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.GetDBConnection().Disconnect()

	router := gin.Default()
	routes.SetupRoutes(router, db.GetDBConnection())
	if err := router.Run(":5005"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
