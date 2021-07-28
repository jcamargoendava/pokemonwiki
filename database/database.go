package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, name string) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database(name)
}
