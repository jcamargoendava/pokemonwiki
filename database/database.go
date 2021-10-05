package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBConnection struct {
	client *mongo.Client
}

func NewConnection(ctx context.Context, name string) (database *mongo.Database, conn *DBConnection, err error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/?compressors=disabled&gssapiServiceName=mongodb")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err
	}

	conn = &DBConnection{
		client: client,
	}
	return client.Database(name), conn, nil
}

func (conn *DBConnection) Close(ctx context.Context) {
	err := conn.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection closed")
}
