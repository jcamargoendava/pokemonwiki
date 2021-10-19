package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcamargoendava/pokemonwiki/dto"
	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/jcamargoendava/pokemonwiki/repository"
	"github.com/jcamargoendava/pokemonwiki/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPokemon godoc
// @Description Gets a Pokemon using the name from the path parameter
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} pokemonModel.Pokemon
// @Failure 404 "not found"
// @Failure 500 {object} dto.Response_500
// @Param name path string true "name"
// @Tags pokemon
// @Router /pokemon/{name} [get]
func GetPokemon(c *gin.Context) {
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon("pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemonName := c.Param("name")
	if pokemonName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name can not be blank"})
		return
	}

	//this is just so go doesn't delete the dto import when I save the file
	var x dto.Response
	fmt.Println(x)

	foundPokemon, err := pokemonService.GetPokemon(ctx, pokemonName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundPokemon})
}

// CreatePokemon godoc
// @Description Creates a new Pokemon based on the information in the body of the post request
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} dto.Response_200
// @Failure 500 {object} dto.Response_500
// @Param Body body dto.pokemon_body true "Request Body"
// @Tags pokemon
// @Router /pokemon/ [post]
func CreatePokemon(c *gin.Context) {
	var pokemon pokemonModel.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon("pokemon")
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

// UpdatePokemon godoc
// @Description Updates a Pokemon based on the id found in the path parameter with the information found in body of the put request
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} pokemonModel.Pokemon
// @Failure 500 "Internal Server Error"
// @Param id path string true "ID"
// @Param Body body dto.pokemon_body true "Request Body"
// @Tags pokemon
// @Router /pokemon/{id} [put]
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
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon("pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	foundPokemon, err := pokemonService.UpdatePokemon(ctx, id, &pokemon)
	if err != nil {
		fmt.Errorf("Error trying to get a pokemon")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": foundPokemon})
}

// DeletePokemon godoc
// @Description Deletes a Pokemon based on the id it finds in the path parameter
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} dto.Response_200
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Param id path string true "ID"
// @Tags pokemon
// @Router /pokemon/id [delete]
func DeletePokemon(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "param id is required"})
		return
	}
	ctx, _ := c.MustGet("ctx").(context.Context)
	pokemonRepo := repository.NewPokemon("pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	if err := pokemonService.DeletePokemon(ctx, id); err != nil {
		fmt.Errorf("Error trying to get a pokemon")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error ocurred"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

// GetPokemons godoc
// @Description Returns a list of all the Pokemons inside the mongoDB database
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} dto.Response_200_get_pokemons
// @Failure 404 "not found"
// @Failure 500 "Internal Server Error"
// @Tags pokemon
// @Router /pokemons/ [get]
func GetPokemons(c *gin.Context) {
	ctx, _ := c.MustGet("ctx").(context.Context)
	limit := c.Query("limit")
	if len(limit) == 0 {
		limit = "100"
	}
	offset := c.Query("offset")
	if len(offset) == 0 {
		offset = "0"
	}

	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)
	pokemonRepo := repository.NewPokemon("pokemon")
	pokemonService := services.NewPokemon(pokemonRepo)
	pokemonsFound := pokemonService.RetrieveAllPokemons(ctx, offsetInt, limitInt)
	c.JSON(http.StatusOK, gin.H{"data": pokemonsFound})
}
