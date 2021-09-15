package repository

import (
	"context"
	"fmt"

	"github.com/jcamargoendava/pokemonwiki/database"
	masterModel "github.com/jcamargoendava/pokemonwiki/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Master struct {
	DB             *mongo.Database
	CollectionName string
}

func NewMaster(collectionName string) *Master {
	return &Master{
		CollectionName: collectionName,
	}
}

func (m *Master) GetMaster(ctx context.Context, id string) (masterModel.Master, error) {
	var masterFound masterModel.Master
	db, conn := database.NewConnection(ctx, "pokemon_database")
	collection := db.Collection(m.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&masterFound)
	conn.Close(ctx)
	if err != nil {
		return masterModel.Master{}, fmt.Errorf("Error trying to get master")
	}
	return masterFound, nil
}

func (m *Master) SaveMaster(ctx context.Context, mModel *masterModel.Master) (*mongo.InsertOneResult, error) {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	collection := db.Collection(m.CollectionName)
	insertedMaster, err := collection.InsertOne(ctx, mModel)
	conn.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error trying to insert")
	}
	fmt.Print(insertedMaster)
	return insertedMaster, nil
}

func (m *Master) UpdateMaster(ctx context.Context, id string, mModel *masterModel.Master) (*masterModel.Master, error) {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	var masterFound masterModel.Master
	collection := db.Collection(m.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.D{{"$set", bson.D{
		{"name", mModel.Name},
		{"gender", mModel.Gender},
		{"age", mModel.Age},
		{"pokemons", mModel.Pokemons}}}}).Decode(&masterFound)
	conn.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error trying to get master")
	}
	return &masterFound, nil
}

func (m *Master) DeleteMaster(ctx context.Context, id string) error {
	db, conn := database.NewConnection(ctx, "pokemon_database")
	collection := db.Collection(m.CollectionName)
	objID, _ := primitive.ObjectIDFromHex(id)
	deletedMaster, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	conn.Close(ctx)
	if err != nil {
		return fmt.Errorf("Error trying to get master")
	}
	fmt.Print(deletedMaster)
	return nil
}
