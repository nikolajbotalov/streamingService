package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello world")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movies/:imdb_id", controllers.GetMovieById())

	err := router.Run(":2000")
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
