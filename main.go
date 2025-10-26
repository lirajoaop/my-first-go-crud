package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lirajoaop/my-first-go-crud/src/configuration/database/mongodb"
	"github.com/lirajoaop/my-first-go-crud/src/configuration/logger"
	"github.com/lirajoaop/my-first-go-crud/src/controller"
	"github.com/lirajoaop/my-first-go-crud/src/controller/routes"
	"github.com/lirajoaop/my-first-go-crud/src/model/repository"
	"github.com/lirajoaop/my-first-go-crud/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	//Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
