package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Katsushi21/travelone"
	"github.com/Katsushi21/travelone/ent"
)

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return travelone.NewExecutableSchema(travelone.Config{
		Resolvers: &Resolver{client},
	})
}
