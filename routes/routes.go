package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pokemonController "github.com/jcamargoendava/pokemonwiki/controllers"
)

func ApiMiddleware(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Set("databaseConn", db)
		c.Set("ctx", ctx)
		c.Next()
	}
}

func StartGin(ctx context.Context) {
	router := gin.Default()
	router.Use(ApiMiddleware(ctx))
	api := router.Group("/api")
	{
		api.GET("are-you-alive", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		})
		api.GET("/pokemons", pokemonController.GetPokemons)
		api.GET("/pokemon/:name", pokemonController.GetPokemon)
		api.POST("/pokemon", pokemonController.CreatePokemon)
		api.PUT("/pokemon/:id", pokemonController.UpdatePokemon)
		api.DELETE("/pokemon/:id", pokemonController.DeletePokemon)

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
