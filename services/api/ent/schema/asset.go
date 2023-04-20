package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Asset struct {
	ent.Schema
}

func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("parent_id", uuid.UUID{}).Optional(),
		field.UUID("location_id", uuid.UUID{}).Optional(),
		field.Text("name").Optional(),
	}
}

func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Asset.Type).From("parent").Unique().Field("parent_id"),
		edge.To("location", Location.Type).Unique().Field("location_id"),
		edge.To("tags", Tag.Type).Through("asset_tags", AssetTag.Type),
	}
}

func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
