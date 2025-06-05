package main

import (
	"log"
	"github.com/abik1221/Kuraz_To_Do_Test/config"
	"github.com/abik1221/Kuraz_To_Do_Test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	r := gin.Default()
	routes.RegisterTaskRoutes(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
