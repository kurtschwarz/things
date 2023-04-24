package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"things-api/ent"
	"things-api/graph/generated"

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return nil, nil
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return nil, nil
}

// Assets is the resolver for the assets field.
func (r *queryResolver) Assets(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.AssetWhereInput) (*ent.AssetConnection, error) {
	return r.client.Asset.Query().Paginate(ctx, after, first, before, last, ent.WithAssetFilter(where.Filter))
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.LocationWhereInput) (*ent.LocationConnection, error) {
	return r.client.Location.Query().Paginate(ctx, after, first, before, last, ent.WithLocationFilter(where.Filter))
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.TagWhereInput) (*ent.TagConnection, error) {
	return r.client.Tag.Query().Paginate(ctx, after, first, before, last, ent.WithTagFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
