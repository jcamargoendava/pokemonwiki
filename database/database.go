package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnection struct {
	client *mongo.Client
}

func NewConnection(ctx context.Context, name string) (database *mongo.Database, conn *DBConnection) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	conn = &DBConnection{
		client: client,
	}
	return client.Database(name), conn
}

func (conn *DBConnection) Close(ctx context.Context) {
	err := conn.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection closed")
}
