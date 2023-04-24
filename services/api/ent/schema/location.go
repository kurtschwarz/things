package schema

import (
	"things-api/ent/schema/mixin"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Location struct {
	ent.Schema
}

func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("parent_id", uuid.UUID{}).Optional().Nillable(),
		field.Text("name").Optional(),
	}
}

func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Location.Type).From("parent").Unique().Field("parent_id"),
	}
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.SoftDeleteMixin{},
	}
}
