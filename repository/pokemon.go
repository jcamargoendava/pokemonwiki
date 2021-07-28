package repository

import (
	"context"
	"fmt"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pokemon struct {
	DB             *mongo.Database
	CollectionName string
}

func NewPokemon(db *mongo.Database, collectionName string) *Pokemon {
	return &Pokemon{
		DB:             db,
		CollectionName: collectionName,
	}
}

func (pk *Pokemon) SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (error, error) {
	collection := pk.DB.Collection(pk.CollectionName)
	insertedPokemon, err := collection.InsertOne(ctx, pkModel)
	if err != nil {
		return nil, fmt.Errorf("Error trying to insert")
	}
	fmt.Print(insertedPokemon)
	return nil, nil
}

// func GetPokemon(id string, pkModel pokemonModel.Pokemon) pokemonModel.Pokemon {

// }

// func GetPokemons(pkModel pokemonModel.Pokemon) (pokemonModel.Pokemon, error) {
// }
