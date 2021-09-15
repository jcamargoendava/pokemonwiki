package repository

import (
	"context"

	"github.com/jcamargoendava/pokemonwiki/database"
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

func NewPokemon(collectionName string) *Pokemon {
	return &Pokemon{
		CollectionName: collectionName,
	}
}

func (pk *Pokemon) GetPokemon(ctx context.Context, name string) (pokemonModel.Pokemon, error) {
	var pokemon pokemonModel.Pokemon
	pokemonFound, err := pokeapi.Pokemon(name)
	if err != nil {
		db, conn := database.NewConnection(ctx, "pokemon_database")
		collection := db.Collection(pk.CollectionName)
		errDB := collection.FindOne(ctx, bson.M{"name": name}).Decode(&pokemon)
		conn.Close(ctx)
		return pokemon, errDB
	}
	return pokemonModel.Pokemon{
		PokemonID: pokemonFound.ID,
		Name:      pokemonFound.Name,
		Img:       pokemonFound.Sprites.FrontDefault,
	}, err
}

func (pk *Pokemon) SavePokemon(ctx context.Context, pkModel *pokemonModel.Pokemon) (*mongo.InsertOneResult, error) {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	collection := db.Collection(pk.CollectionName)
	insertedPokemon, err := collection.InsertOne(ctx, pkModel)
	conn.Close(ctx)
	return insertedPokemon, err
}

func (pk *Pokemon) UpdatePokemon(ctx context.Context, id string, pkModel *pokemonModel.Pokemon) (pokemonModel.Pokemon, error) {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	var pokemon pokemonModel.Pokemon
	collection := db.Collection(pk.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.D{{"$set", bson.D{
		{"name", pkModel.Name},
		{"img", pkModel.Img},
	}}}).Decode(&pokemon)
	conn.Close(ctx)
	return pokemon, err
}

func (pk *Pokemon) DeletePokemon(ctx context.Context, id string) error {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	collection := db.Collection(pk.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	conn.Close(ctx)
	return err
}
