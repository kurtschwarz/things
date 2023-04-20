package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// AssetTag holds the schema definition for the AssetTag entity.
type AssetTag struct {
	ent.Schema
}

// Fields of the AssetTag.
func (AssetTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("asset_id", uuid.UUID{}),
		field.UUID("tag_id", uuid.UUID{}),
	}
}

// Edges of the AssetTag.
func (AssetTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("asset", Asset.Type).Unique().Required().Field("asset_id"),
		edge.To("tag", Tag.Type).Unique().Required().Field("tag_id"),
	}
}

// Indexes of the UserTweet.
func (AssetTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag_id").
			Unique(),
	}
}
