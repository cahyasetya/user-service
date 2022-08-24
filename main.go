package main

import (
	"context"
	"time"

	"github.com/cahyasetya/user-service/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.
		Client().
		ApplyURI("mongodb+srv://<user>:<password>@users.a4msqyl.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	router := gin.Default()
	client := initDB()
	userController := controllers.UserController{Mongo: client}
	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUser)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
