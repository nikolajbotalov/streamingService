package routes

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())

	router.GET("/movies/:imdb_id", controllers.GetMovieById())
	router.POST("/movies/", controllers.AddMovie())
	router.GET("/movies/recommended", controllers.GetRecommendedMovies())
	router.PATCH("/review/:imdb_id", controllers.AdminReviewUpdate())
}
