package dto

type pokemon_body struct {
	PokemonID int    `pokemon_id:""`
	Name      string `name:""`
	Img       string `img:""`
}
