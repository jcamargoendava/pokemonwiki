package lib

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createConnection() {
	mgm.SetDefaultConfig(nil, "pokemon_database", options.Client().ApplyURI("mongodb+srv://juan_endava:Welcome2021%21@cluster0.0brzk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
}
