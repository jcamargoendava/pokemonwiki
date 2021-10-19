package dto

type Response struct {
	Message string `json:"message"`
}

type Response_500 struct {
	Error string `error:"mongo: no documents in result"`
}

type Response_200 struct {
	Data string `data:""`
}
type Response_200_get_pokemons struct {
	Data []struct {
		Data string `data:""`
	}
}
