package main

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/syumai/gqlgen-directives/directives"
	"github.com/syumai/gqlgen-directives/example/generated"
)

func TestDirectives(t *testing.T) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{},
		Directives: generated.DirectiveRoot{
			OneOf: directives.OneOf,
		},
	}))

	c := client.New(srv)

	t.Run("oneOf", func(t *testing.T) {
		var i interface{}
		err := c.Post(`mutation { createFruit(input: { apple: { color: RED }, banana: null, grape: null }) { id color } }`, &i)
		if err != nil {
			t.Fatal(err)
		}
	})
}
