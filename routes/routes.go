package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pokemonController "github.com/jcamargoendava/pokemonwiki/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApiMiddleware(ctx context.Context, db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Set("ctx", ctx)
		c.Next()
	}
}

func StartGin(ctx context.Context, db *mongo.Database) {
	router := gin.Default()
	router.Use(ApiMiddleware(ctx, db))
	api := router.Group("/api")
	{
		api.GET("are-you-alive", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		})
		api.GET("/pokemons", pokemonController.GetPokemons)
		api.GET("/master/:id", pokemonController.GetMaster)
		api.POST("/master", pokemonController.CreateMaster)
		api.PUT("/master/:id", pokemonController.UpdateMaster)
		api.DELETE("/master/:id", pokemonController.DeleteMaster)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8081")
}
