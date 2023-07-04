package ports

import (
	"context"

	"github.com/tocura/go-pokemon/core/model"
)

type PokemonTrainerService interface {
	Create(ctx context.Context, pokemonTrainer *model.PokemonTrainer) (*model.PokemonTrainer, error)
	FindByID(ctx context.Context, id string) (*model.PokemonTrainer, error)
	PatchPokemons(ctx context.Context, trainerID string, pokemons model.Pokemons) (*model.PokemonTrainer, error)
	AddPokemon(ctx context.Context, trainerID, pokemon string) (*model.PokemonTrainer, error)
	RemovePokemon(ctx context.Context, trainerID, pokemon string) (*model.PokemonTrainer, error)
}

type PokeAPIService interface {
	GetPokemonInGameVersion(ctx context.Context, gameVersion string, pokemon *model.Pokemon) (*model.Pokemon, error)
}
