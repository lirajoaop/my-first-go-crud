package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lirajoaop/my-first-go-crud/src/configuration/logger"
	"github.com/lirajoaop/my-first-go-crud/src/controller/routes"
)

func main() {
	logger.Info("About to star user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
