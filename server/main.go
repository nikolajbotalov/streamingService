package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/routes"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	err := router.Run(":2000")
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
