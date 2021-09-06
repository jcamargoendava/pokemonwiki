package services

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
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

func (p *Pokemon) RetrieveAllPokemons(ctx context.Context, offset int, limit int) []pokemonModel.Pokemon {
	ch := make(chan structs.Resource)
	var wg sync.WaitGroup
	start := time.Now()
	loops := float64(limit) / float64(50)
	goRoutines := int(math.Ceil(loops))
	for i := 0; i < goRoutines; i++ {
		wg.Add(1)
		go GetPokemonRequest(offset, limit, ch, &wg)
		offset = offset + limit
	}

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
		fmt.Println(time.Since(start))
	}()
	// read from channel as they come in until its closed
	var pokemons = []pokemonModel.Pokemon{}
	for pokemonsFound := range ch {
		wg := sync.WaitGroup{}
		for _, pokemonFound := range pokemonsFound.Results {
			wg.Add(1)
			go func(pokemonName string) {
				pkmn, err := pokeapi.Pokemon(pokemonName)
				if err != nil {
					fmt.Errorf("Error trying to get pokemon %s", pokemonName)
				}
				pokemon := pokemonModel.Pokemon{
					PokemonID: pkmn.ID,
					Name:      pkmn.Name,
					Img:       pkmn.Sprites.FrontDefault,
				}
				pokemons = append(pokemons, pokemon)
				wg.Done()
			}(pokemonFound.Name)
		}
		wg.Wait()
	}
	return pokemons
}

func GetPokemonRequest(offset, limit int, ch chan<- structs.Resource, wg *sync.WaitGroup) {
	defer wg.Done()
	pokemonsFound, err := pokeapi.Resource("pokemon", offset, limit)
	if err != nil {
		fmt.Errorf("Error trying to get pokemons")
	}
	ch <- pokemonsFound
}
