package main

import (
	"context"
	"log"
	"net/http"

	"things-api/ent"
	"things-api/ent/migrate"
	"things-api/pkg/resolver"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:things.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	http.Handle("/",
		playground.Handler("User", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
