package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Telekurysh/ozon_test/pkg/graphql/generated"
	"net/http"
)

func NewGraphQLHandler() http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))
	return srv
}

func NewPlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL Playground", "/query")
}
