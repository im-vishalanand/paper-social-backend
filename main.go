package main

import (
    "log"
    "net/http"
    "github.com/99designs/gqlgen/handler"
    "github.com/im-vishalanand/paper-social-backend/graphql"
)

func main() {
    resolver := &graphql.Resolver{}
    srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

    http.Handle("/graphql", srv)

    log.Println("🚀 GraphQL server running on http://localhost:8080/graphql")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
