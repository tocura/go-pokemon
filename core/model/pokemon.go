package model

type Pokemon struct {
	ID             string
	Name           string
	PokedexID      int
	BaseExperience int
}

type Pokemons []Pokemon
