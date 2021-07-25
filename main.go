package main

import (
	"fmt"

	"github.com/jcamargoendava/pokemonwiki/routes"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := mgm.SetDefaultConfig(nil, "pokemon_database", options.Client().ApplyURI("mongodb+srv://juan_endava:Welcome2021%21@cluster0.0brzk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Errorf("Error trying to connect to db")
	}
	routes.StartGin()
}
