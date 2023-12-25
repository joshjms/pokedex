package database

import (
	"encoding/json"
	"os"
	"pokedex/src/models"
)

func New() (*[]models.Pokemon, error) {
	pokedexJson, err := os.ReadFile("src/data/pokedex.json")
	if err != nil {
		return nil, err
	}

	var database []models.Pokemon

	err = json.Unmarshal(pokedexJson, &database)

	if err != nil {
		return nil, err
	}

	return &database, nil
}
