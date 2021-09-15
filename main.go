package main

import (
	"context"

	"github.com/jcamargoendava/pokemonwiki/routes"
)

func main() {
	ctx := context.TODO()
	// db := database.ConnectDB(ctx, "pokemon_database")
	routes.StartGin(ctx)
}
