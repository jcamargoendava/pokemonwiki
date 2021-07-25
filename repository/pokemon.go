package repository

import pokemonModel "github.com/jcamargoendava/pokemonwiki/models"

type PokemonRepository interface {
	SavePokemon(name, img string) *pokemonModel.Pokemon
	GetPokemon(name string) (*pokemonModel.Pokemon, error)
	GetPokemons() ([]*pokemonModel.Pokemon, error)
}
