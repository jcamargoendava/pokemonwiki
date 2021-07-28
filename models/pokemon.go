package models

import "time"

type Pokemon struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" bson:"name"`
	Img       string    `json:"img" bson:"img"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewPokemon(name string, img string) *Pokemon {
	return &Pokemon{
		Name: name,
		Img:  img,
	}
}
