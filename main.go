package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"GoLang_Project/controllers"
	"GoLang_Project/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	connectionString := "mongodb://127.0.0.1:27017/Go_DB"

	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	r := gin.Default()

	employeeCollection := mongoClient.Database("Go_DB").Collection("employees")
	employeeController := &controllers.EmployeeController{Collection: employeeCollection}

	routes.DefineEmployeeRoutes(r, employeeController)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
