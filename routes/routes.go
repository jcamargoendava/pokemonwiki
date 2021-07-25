package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pokemonController "github.com/jcamargoendava/pokemonwiki/controllers"
)

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("are-you-alive", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "running"})
		})
		api.GET("/pokemons", pokemonController.GetPokemons)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8081")
}
