package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcamargoendava/pokemonwiki/repository"
	"github.com/jcamargoendava/pokemonwiki/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPokemons(c *gin.Context) {
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	fmt.Print(db, ctx)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemonsFound := pokemonService.RetrieveAllPokemons(ctx)
	c.JSON(http.StatusOK, gin.H{"data": pokemonsFound})
}
