package main

import (
	"pokedex/src/database"
	"pokedex/src/routes"
)

func main() {
	db, err := database.New()
	if err != nil {
		panic(err)
	}

	e := routes.Init(db)
	e.Logger.Fatal(e.Start(":8080"))
}
