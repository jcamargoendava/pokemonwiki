package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	PokemonID int                `json:"pokemon_id" bson:"pokemon_id"`
	Name      string             `json:"name" bson:"name"`
	Img       string             `json:"img" bson:"img"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
	// Master    Master             `json:"master"`
}

func NewPokemon(name string, img string) *Pokemon {
	return &Pokemon{
		Name: name,
		Img:  img,
	}
}
