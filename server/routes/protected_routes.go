package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"server/controllers"
	"server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movies/:imdb_id", controllers.GetMovieById(client))
	router.POST("/movies/", controllers.AddMovie(client))
	router.GET("/movies/recommended", controllers.GetRecommendedMovies(client))
	router.PATCH("/review/:imdb_id", controllers.AdminReviewUpdate(client))
}
