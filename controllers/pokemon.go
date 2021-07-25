package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pokemonService "github.com/jcamargoendava/pokemonwiki/services"
)

func GetPokemons(c *gin.Context) {
	pokemonsFound := pokemonService.RetrieveAllPokemons()
	c.JSON(http.StatusOK, gin.H{"data": pokemonsFound})
}
