package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/cahyasetya/user-service/core/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Mongo *mongo.Client
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	coll := controller.Mongo.Database("users").Collection("users")
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (controller *UserController) GetUser(c *gin.Context) {
	name := c.Query("name")
	var user entities.User
	filter := bson.D{{"name", name}}

	coll := controller.Mongo.Database("users").Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Printf("error: %v", err)
	}
	c.IndentedJSON(http.StatusOK, user)
}
