//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var err error
	var ex *entgql.Extension
	if ex, err = entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graph/ent.graphqls"),
		entgql.WithConfigPath("./gqlgen.yml"),
	); err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err = entc.Generate(
		"./ent/schema",
		&gen.Config{},
		[]entc.Option{
			entc.Extensions(ex),
		}...,
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
