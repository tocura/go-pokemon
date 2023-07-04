package ports

import (
	"context"

	"github.com/tocura/go-pokemon/core/model"
)

type Database interface {
	Save(ctx context.Context, pokemonTrainer *model.PokemonTrainer) (*model.PokemonTrainer, error)
	FindByID(ctx context.Context, id string) (*model.PokemonTrainer, error)
}
