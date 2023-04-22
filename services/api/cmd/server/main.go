package main

import (
	"context"
	"log"
	"net/http"

	"things-api/ent"
	"things-api/ent/migrate"
	_ "things-api/ent/runtime"
	"things-api/pkg/resolver"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:things.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err = client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	router := chi.NewRouter()
	router.Use(
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		}).Handler,
	)

	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	router.Handle("/", playground.Handler("Things", "/graphql"))
	router.Handle("/graphql", srv)

	if err = http.ListenAndServe(":8081", router); err != nil {
		panic(err)
	}
}
