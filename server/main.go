package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"server/database"
	"server/routes"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unable to load .env file")
	}

	var client *mongo.Client = database.Connect()
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Faiiled to reach server: %v", err)
	}

	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	err = router.Run(":2000")
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
