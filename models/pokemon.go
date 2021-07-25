package models

import "github.com/kamva/mgm/v3"

type Pokemon struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `json:"name" bson:"name"`
	Img  string `json:"img" bson:"img"`
}

func NewPokemon(name string, img string) *Pokemon {
	return &Pokemon{
		Name: name,
		Img:  img,
	}
}
