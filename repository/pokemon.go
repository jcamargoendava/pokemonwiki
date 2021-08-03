package repository

import (
	"context"

	pokemonModel "github.com/jcamargoendava/pokemonwiki/models"
	"github.com/mtslzr/pokeapi-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (pk *Pokemon) GetPokemon(ctx context.Context, name string) (pokemonModel.Pokemon, error) {
	var pokemon pokemonModel.Pokemon
	pokemonFound, err := pokeapi.Pokemon(name)
	if err != nil {
		collection := pk.DB.Collection(pk.CollectionName)
		errDB := collection.FindOne(ctx, bson.M{"name": name}).Decode(&pokemon)
		return pokemon, errDB
	}
	return pokemonModel.Pokemon{
		PokemonID: pokemonFound.ID,
		Name:      pokemonFound.Name,
		Img:       pokemonFound.Sprites.FrontDefault,
	}, err
}

func (pk *Pokemon) SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (*mongo.InsertOneResult, error) {
	collection := pk.DB.Collection(pk.CollectionName)
	insertedPokemon, err := collection.InsertOne(ctx, pkModel)
	return insertedPokemon, err
}

func (pk *Pokemon) UpdatePokemon(ctx context.Context, id string, pkModel *pokemonModel.Pokemon) (pokemonModel.Pokemon, error) {
	var pokemon pokemonModel.Pokemon
	collection := pk.DB.Collection(pk.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.D{{"$set", bson.D{
		{"name", pkModel.Name},
		{"img", pkModel.Img},
	}}}).Decode(&pokemon)
	return pokemon, err
}

func (pk *Pokemon) DeletePokemon(ctx context.Context, id string) error {
	collection := pk.DB.Collection(pk.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
