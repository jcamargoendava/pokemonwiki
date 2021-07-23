package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context) {
	pokemons := "Charmander"
	c.JSON(http.StatusOK, gin.H{"data": pokemons})
}
