package model

import "github.com/tocura/go-pokemon/core/enum"

type PokemonTrainer struct {
	ID          string
	Name        string
	Gender      enum.Gender
	Birthdate   string
	GameVersion enum.GameVersion
	Pokemons    Pokemons
}
