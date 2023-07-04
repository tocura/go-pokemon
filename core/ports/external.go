package ports

import (
	"context"

	"github.com/tocura/go-pokemon/core/model"
)

type PokeAPIClient interface {
	GetPokemon(ctx context.Context, pokemon string) (*model.PokeAPI, error)
}
