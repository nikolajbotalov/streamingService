package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/database"
	"server/routes"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	var client *mongo.Client = database.Connect()

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	err := router.Run(":2000")
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
