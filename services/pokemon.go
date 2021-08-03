package services

import (
	"context"
	"fmt"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/mtslzr/pokeapi-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonRepository interface {
	GetPokemon(ctx context.Context, name string) (pokemonModel.Pokemon, error)
	SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (*mongo.InsertOneResult, error)
	UpdatePokemon(ctx context.Context, id string, pkModel *pokemonModel.Pokemon) (pokemonModel.Pokemon, error)
	DeletePokemon(ctx context.Context, id string) error
}

type Pokemon struct {
	Repo PokemonRepository
}

func NewPokemon(repo PokemonRepository) *Pokemon {
	return &Pokemon{
		Repo: repo,
	}
}

func (p *Pokemon) GetPokemon(ctx context.Context, name string) (pokemonModel.Pokemon, error) {
	pokemon, err := p.Repo.GetPokemon(ctx, name)
	return pokemon, err
}

func (p *Pokemon) SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (*mongo.InsertOneResult, error) {
	insertedPokemon, err := p.Repo.SavePokemon(ctx, pkModel)
	return insertedPokemon, err
}

func (p *Pokemon) UpdatePokemon(ctx context.Context, id string, pkModel *pokemonModel.Pokemon) (pokemonModel.Pokemon, error) {
	foundPokemon, err := p.Repo.UpdatePokemon(ctx, id, pkModel)
	return foundPokemon, err
}

func (p *Pokemon) DeletePokemon(ctx context.Context, id string) error {
	err := p.Repo.DeletePokemon(ctx, id)
	return err
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
		pokemon := pokemonModel.Pokemon{
			PokemonID: pkmn.ID,
			Name:      pkmn.Name,
			Img:       pkmn.Sprites.FrontDefault,
		}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}
