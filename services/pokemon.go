package services

import (
	"context"
	"fmt"
	"strconv"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/mtslzr/pokeapi-go"
)

type PokemonRepository interface {
	SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (error, error)
	// GetPokemon(ctx context.Context, id string) (*pokemonModel.Pokemon, error)
	// GetPokemons(ctx context.Context) ([]*pokemonModel.Pokemon, error)
}

type Pokemon struct {
	Repo PokemonRepository
}

func NewPokemon(repo PokemonRepository) *Pokemon {
	return &Pokemon{
		Repo: repo,
	}
}

func (p *Pokemon) RetrieveAllPokemons(ctx context.Context) []pokemonModel.Pokemon {
	pokemonsFound, err := pokeapi.Resource("pokemon")
	var pokemons = []pokemonModel.Pokemon{}
	if err != nil {
		fmt.Errorf("Error trying to get pokemons")
	}
	for _, pokemonFound := range pokemonsFound.Results {
		pkmn, err := pokeapi.Pokemon(pokemonFound.Name)
		if err != nil {
			fmt.Errorf("Error trying to get pokemon %s", pokemonFound.Name)
		}
		fmt.Print(pkmn)
		pokemon := pokemonModel.Pokemon{
			ID:   strconv.Itoa(pkmn.ID),
			Name: pkmn.Name,
			Img:  pkmn.Sprites.FrontDefault,
		}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}
