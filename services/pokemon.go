package services

import (
	"fmt"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/kamva/mgm/v3"
	"github.com/mtslzr/pokeapi-go"
)

type Pokemon struct {
	ID       int
	Name     string
	Category string
	Img      string
	Gender   string
}

func SavePokemon(name, img string) *pokemonModel.Pokemon {
	return &pokemonModel.Pokemon{
		Name: name,
		Img:  img,
	}
}

func RetrieveAllPokemons() []Pokemon {
	pokemonsFound, err := pokeapi.Resource("pokemon")
	var pokemons = []Pokemon{}
	if err != nil {
		fmt.Errorf("Error trying to get pokemons")
	}
	for _, pokemonFound := range pokemonsFound.Results {
		pkmn, err := pokeapi.Pokemon(pokemonFound.Name)
		if err != nil {
			fmt.Errorf("Error trying to get pokemon %s", pokemonFound.Name)
		}
		pokemon := Pokemon{ID: pkmn.ID, Name: pkmn.Name, Category: "", Img: pkmn.Sprites.FrontDefault, Gender: ""}
		pokemons = append(pokemons, pokemon)
	}
	pkmModel := pokemonModel.NewPokemon("test", "img.png")
	mgm.Coll(pkmModel).Create(pkmModel)

	// SavePokemon("test", "img.png")
	return pokemons
}
