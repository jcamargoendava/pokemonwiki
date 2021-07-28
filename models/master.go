package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Master struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Gender    string             `json:"gender" bson:"gender"`
	Age       string             `json:"age" bson:"age"`
	Pokemons  []string           `json:"pokemons" bson:"pokemons"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}

func NewMaster(name string, gender string, age string, pokemons []string) *Master {
	return &Master{
		Name:     name,
		Gender:   gender,
		Age:      age,
		Pokemons: pokemons,
	}
}
