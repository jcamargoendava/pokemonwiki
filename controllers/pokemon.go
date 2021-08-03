package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/jcamargoendava/pokemonwiki/repository"
	"github.com/jcamargoendava/pokemonwiki/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPokemon(c *gin.Context) {
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemonName := c.Param("name")
	if pokemonName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name can not be blank"})
		return
	}
	foundPokemon, err := pokemonService.GetPokemon(ctx, pokemonName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundPokemon})
}

func CreatePokemon(c *gin.Context) {
	var pokemon pokemonModel.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemon.ID = primitive.NewObjectID()
	createdPokemon, err := pokemonService.SavePokemon(ctx, &pokemon)
	if err != nil {
		fmt.Errorf("Error trying to create a pokemon")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": createdPokemon})
}

func UpdatePokemon(c *gin.Context) {
	var pokemon pokemonModel.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	foundPokemon, err := pokemonService.UpdatePokemon(ctx, id, &pokemon)
	if err != nil {
		fmt.Errorf("Error trying to get a pokemon")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundPokemon})
}

func DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	if err := pokemonService.DeletePokemon(ctx, id); err != nil {
		fmt.Errorf("Error trying to get a pokemon")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

func GetPokemons(c *gin.Context) {
	db, _ := c.MustGet("databaseConn").(*mongo.Database)
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon(db, "pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemonsFound := pokemonService.RetrieveAllPokemons(ctx)
	c.JSON(http.StatusOK, gin.H{"data": pokemonsFound})
}
