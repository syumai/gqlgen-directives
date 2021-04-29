//go:generate go run github.com/99designs/gqlgen
package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/syumai/gqlgen-directives/example/generated"
	"github.com/syumai/gqlgen-directives/example/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateFruit(ctx context.Context, input model.FruitInput) (model.Fruit, error) {
	maybeID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	id := maybeID.String()
	if input.Apple != nil {
		return &model.Apple{ID: id, Color: input.Apple.Color}, nil
	}
	if input.Banana != nil {
		return &model.Banana{ID: id, Color: input.Banana.Color}, nil
	}
	if input.Grape != nil {
		return &model.Grape{ID: id, Color: input.Grape.Color}, nil
	}
	return nil, fmt.Errorf("input must be given")
}

func (r *queryResolver) Fruits(ctx context.Context) ([]model.Fruit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
