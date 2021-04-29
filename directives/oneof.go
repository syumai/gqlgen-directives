package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func OneOf(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	rc := graphql.GetFieldContext(ctx)
	log.Printf("input object logging: %v, %s, %T, %+v", rc.Path(), rc.Field.Name, obj, obj)
	return next(ctx)
}
