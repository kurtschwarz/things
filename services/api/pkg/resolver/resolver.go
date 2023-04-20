package resolver

import (
	"things-api/ent"
	"things-api/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{
				client,
			},
		},
	)
}
