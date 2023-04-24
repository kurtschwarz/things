package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"fmt"
	"things-api/ent"
	"things-api/graph/generated"
	"things-api/pkg/util"

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

// ID is the resolver for the id field.
func (r *assetResolver) ID(ctx context.Context, obj *ent.Asset) (string, error) {
	return obj.ID.String(), nil
}

// ParentID is the resolver for the parentID field.
func (r *assetResolver) ParentID(ctx context.Context, obj *ent.Asset) (*string, error) {
	return util.StringPtr(obj.ParentID.String()), nil
}

// LocationID is the resolver for the locationID field.
func (r *assetResolver) LocationID(ctx context.Context, obj *ent.Asset) (*string, error) {
	return util.StringPtr(obj.LocationID.String()), nil
}

// ID is the resolver for the id field.
func (r *assetTagResolver) ID(ctx context.Context, obj *ent.AssetTag) (string, error) {
	return obj.ID.String(), nil
}

// AssetID is the resolver for the assetID field.
func (r *assetTagResolver) AssetID(ctx context.Context, obj *ent.AssetTag) (string, error) {
	return obj.AssetID.String(), nil
}

// TagID is the resolver for the tagID field.
func (r *assetTagResolver) TagID(ctx context.Context, obj *ent.AssetTag) (string, error) {
	return obj.TagID.String(), nil
}

// ID is the resolver for the id field.
func (r *locationResolver) ID(ctx context.Context, obj *ent.Location) (string, error) {
	return obj.ID.String(), nil
}

// ParentID is the resolver for the parentID field.
func (r *locationResolver) ParentID(ctx context.Context, obj *ent.Location) (*string, error) {
	if obj.ParentID == nil {
		return nil, nil
	}

	return util.StringPtr(obj.ParentID.String()), nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	return r.client.Noder(ctx, uuid.MustParse(id))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	return r.client.Noders(ctx, util.StringsToUUIDs(ids))
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

// ID is the resolver for the id field.
func (r *tagResolver) ID(ctx context.Context, obj *ent.Tag) (string, error) {
	return obj.ID.String(), nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// ParentID is the resolver for the parentID field.
func (r *userResolver) ParentID(ctx context.Context, obj *ent.User) (*string, error) {
	return util.StringPtr(obj.ParentID.String()), nil
}

// ID is the resolver for the id field.
func (r *assetTagWhereInputResolver) ID(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *assetTagWhereInputResolver) IDNeq(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *assetTagWhereInputResolver) IDIn(ctx context.Context, obj *ent.AssetTagWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *assetTagWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.AssetTagWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *assetTagWhereInputResolver) IDGt(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *assetTagWhereInputResolver) IDGte(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *assetTagWhereInputResolver) IDLt(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *assetTagWhereInputResolver) IDLte(ctx context.Context, obj *ent.AssetTagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ID is the resolver for the id field.
func (r *assetWhereInputResolver) ID(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *assetWhereInputResolver) IDNeq(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *assetWhereInputResolver) IDIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *assetWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *assetWhereInputResolver) IDGt(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *assetWhereInputResolver) IDGte(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *assetWhereInputResolver) IDLt(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *assetWhereInputResolver) IDLte(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ParentID is the resolver for the parentID field.
func (r *assetWhereInputResolver) ParentID(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ParentIDNeq is the resolver for the parentIDNEQ field.
func (r *assetWhereInputResolver) ParentIDNeq(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentIDNeq - parentIDNEQ"))
}

// ParentIDIn is the resolver for the parentIDIn field.
func (r *assetWhereInputResolver) ParentIDIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDIn - parentIDIn"))
}

// ParentIDNotIn is the resolver for the parentIDNotIn field.
func (r *assetWhereInputResolver) ParentIDNotIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDNotIn - parentIDNotIn"))
}

// LocationID is the resolver for the locationID field.
func (r *assetWhereInputResolver) LocationID(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: LocationID - locationID"))
}

// LocationIDNeq is the resolver for the locationIDNEQ field.
func (r *assetWhereInputResolver) LocationIDNeq(ctx context.Context, obj *ent.AssetWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: LocationIDNeq - locationIDNEQ"))
}

// LocationIDIn is the resolver for the locationIDIn field.
func (r *assetWhereInputResolver) LocationIDIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: LocationIDIn - locationIDIn"))
}

// LocationIDNotIn is the resolver for the locationIDNotIn field.
func (r *assetWhereInputResolver) LocationIDNotIn(ctx context.Context, obj *ent.AssetWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: LocationIDNotIn - locationIDNotIn"))
}

// ParentID is the resolver for the parentID field.
func (r *createAssetInputResolver) ParentID(ctx context.Context, obj *ent.CreateAssetInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ChildIDs is the resolver for the childIDs field.
func (r *createAssetInputResolver) ChildIDs(ctx context.Context, obj *ent.CreateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: ChildIDs - childIDs"))
}

// LocationID is the resolver for the locationID field.
func (r *createAssetInputResolver) LocationID(ctx context.Context, obj *ent.CreateAssetInput, data *string) error {
	panic(fmt.Errorf("not implemented: LocationID - locationID"))
}

// TagIDs is the resolver for the tagIDs field.
func (r *createAssetInputResolver) TagIDs(ctx context.Context, obj *ent.CreateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: TagIDs - tagIDs"))
}

// ParentID is the resolver for the parentID field.
func (r *createLocationInputResolver) ParentID(ctx context.Context, obj *ent.CreateLocationInput, data *string) error {
	var uid uuid.UUID
	var err error

	if data != nil {
		if uid, err = uuid.Parse(*data); err != nil {
			return err
		}

		obj.ParentID = &uid
	}

	return nil
}

// ChildIDs is the resolver for the childIDs field.
func (r *createLocationInputResolver) ChildIDs(ctx context.Context, obj *ent.CreateLocationInput, data []string) error {
	panic(fmt.Errorf("not implemented: ChildIDs - childIDs"))
}

// AssetIDs is the resolver for the assetIDs field.
func (r *createTagInputResolver) AssetIDs(ctx context.Context, obj *ent.CreateTagInput, data []string) error {
	panic(fmt.Errorf("not implemented: AssetIDs - assetIDs"))
}

// ParentID is the resolver for the parentID field.
func (r *createUserInputResolver) ParentID(ctx context.Context, obj *ent.CreateUserInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ChildIDs is the resolver for the childIDs field.
func (r *createUserInputResolver) ChildIDs(ctx context.Context, obj *ent.CreateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: ChildIDs - childIDs"))
}

// ID is the resolver for the id field.
func (r *locationWhereInputResolver) ID(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *locationWhereInputResolver) IDNeq(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *locationWhereInputResolver) IDIn(ctx context.Context, obj *ent.LocationWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *locationWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.LocationWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *locationWhereInputResolver) IDGt(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *locationWhereInputResolver) IDGte(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *locationWhereInputResolver) IDLt(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *locationWhereInputResolver) IDLte(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ParentID is the resolver for the parentID field.
func (r *locationWhereInputResolver) ParentID(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ParentIDNeq is the resolver for the parentIDNEQ field.
func (r *locationWhereInputResolver) ParentIDNeq(ctx context.Context, obj *ent.LocationWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentIDNeq - parentIDNEQ"))
}

// ParentIDIn is the resolver for the parentIDIn field.
func (r *locationWhereInputResolver) ParentIDIn(ctx context.Context, obj *ent.LocationWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDIn - parentIDIn"))
}

// ParentIDNotIn is the resolver for the parentIDNotIn field.
func (r *locationWhereInputResolver) ParentIDNotIn(ctx context.Context, obj *ent.LocationWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDNotIn - parentIDNotIn"))
}

// ID is the resolver for the id field.
func (r *tagWhereInputResolver) ID(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *tagWhereInputResolver) IDNeq(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *tagWhereInputResolver) IDIn(ctx context.Context, obj *ent.TagWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *tagWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.TagWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *tagWhereInputResolver) IDGt(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *tagWhereInputResolver) IDGte(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *tagWhereInputResolver) IDLt(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *tagWhereInputResolver) IDLte(ctx context.Context, obj *ent.TagWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ParentID is the resolver for the parentID field.
func (r *updateAssetInputResolver) ParentID(ctx context.Context, obj *ent.UpdateAssetInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// AddChildIDs is the resolver for the addChildIDs field.
func (r *updateAssetInputResolver) AddChildIDs(ctx context.Context, obj *ent.UpdateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddChildIDs - addChildIDs"))
}

// RemoveChildIDs is the resolver for the removeChildIDs field.
func (r *updateAssetInputResolver) RemoveChildIDs(ctx context.Context, obj *ent.UpdateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveChildIDs - removeChildIDs"))
}

// LocationID is the resolver for the locationID field.
func (r *updateAssetInputResolver) LocationID(ctx context.Context, obj *ent.UpdateAssetInput, data *string) error {
	panic(fmt.Errorf("not implemented: LocationID - locationID"))
}

// AddTagIDs is the resolver for the addTagIDs field.
func (r *updateAssetInputResolver) AddTagIDs(ctx context.Context, obj *ent.UpdateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddTagIDs - addTagIDs"))
}

// RemoveTagIDs is the resolver for the removeTagIDs field.
func (r *updateAssetInputResolver) RemoveTagIDs(ctx context.Context, obj *ent.UpdateAssetInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveTagIDs - removeTagIDs"))
}

// ParentID is the resolver for the parentID field.
func (r *updateLocationInputResolver) ParentID(ctx context.Context, obj *ent.UpdateLocationInput, data *string) error {
	var uid uuid.UUID
	var err error

	if data != nil {
		if uid, err = uuid.Parse(*data); err != nil {
			return err
		}

		obj.ParentID = &uid
	}

	return nil
}

// AddChildIDs is the resolver for the addChildIDs field.
func (r *updateLocationInputResolver) AddChildIDs(ctx context.Context, obj *ent.UpdateLocationInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddChildIDs - addChildIDs"))
}

// RemoveChildIDs is the resolver for the removeChildIDs field.
func (r *updateLocationInputResolver) RemoveChildIDs(ctx context.Context, obj *ent.UpdateLocationInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveChildIDs - removeChildIDs"))
}

// AddAssetIDs is the resolver for the addAssetIDs field.
func (r *updateTagInputResolver) AddAssetIDs(ctx context.Context, obj *ent.UpdateTagInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddAssetIDs - addAssetIDs"))
}

// RemoveAssetIDs is the resolver for the removeAssetIDs field.
func (r *updateTagInputResolver) RemoveAssetIDs(ctx context.Context, obj *ent.UpdateTagInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveAssetIDs - removeAssetIDs"))
}

// ParentID is the resolver for the parentID field.
func (r *updateUserInputResolver) ParentID(ctx context.Context, obj *ent.UpdateUserInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// AddChildIDs is the resolver for the addChildIDs field.
func (r *updateUserInputResolver) AddChildIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: AddChildIDs - addChildIDs"))
}

// RemoveChildIDs is the resolver for the removeChildIDs field.
func (r *updateUserInputResolver) RemoveChildIDs(ctx context.Context, obj *ent.UpdateUserInput, data []string) error {
	panic(fmt.Errorf("not implemented: RemoveChildIDs - removeChildIDs"))
}

// ID is the resolver for the id field.
func (r *userWhereInputResolver) ID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// IDNeq is the resolver for the idNEQ field.
func (r *userWhereInputResolver) IDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDNeq - idNEQ"))
}

// IDIn is the resolver for the idIn field.
func (r *userWhereInputResolver) IDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDIn - idIn"))
}

// IDNotIn is the resolver for the idNotIn field.
func (r *userWhereInputResolver) IDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: IDNotIn - idNotIn"))
}

// IDGt is the resolver for the idGT field.
func (r *userWhereInputResolver) IDGt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGt - idGT"))
}

// IDGte is the resolver for the idGTE field.
func (r *userWhereInputResolver) IDGte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDGte - idGTE"))
}

// IDLt is the resolver for the idLT field.
func (r *userWhereInputResolver) IDLt(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLt - idLT"))
}

// IDLte is the resolver for the idLTE field.
func (r *userWhereInputResolver) IDLte(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: IDLte - idLTE"))
}

// ParentID is the resolver for the parentID field.
func (r *userWhereInputResolver) ParentID(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentID - parentID"))
}

// ParentIDNeq is the resolver for the parentIDNEQ field.
func (r *userWhereInputResolver) ParentIDNeq(ctx context.Context, obj *ent.UserWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: ParentIDNeq - parentIDNEQ"))
}

// ParentIDIn is the resolver for the parentIDIn field.
func (r *userWhereInputResolver) ParentIDIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDIn - parentIDIn"))
}

// ParentIDNotIn is the resolver for the parentIDNotIn field.
func (r *userWhereInputResolver) ParentIDNotIn(ctx context.Context, obj *ent.UserWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: ParentIDNotIn - parentIDNotIn"))
}

// Asset returns generated.AssetResolver implementation.
func (r *Resolver) Asset() generated.AssetResolver { return &assetResolver{r} }

// AssetTag returns generated.AssetTagResolver implementation.
func (r *Resolver) AssetTag() generated.AssetTagResolver { return &assetTagResolver{r} }

// Location returns generated.LocationResolver implementation.
func (r *Resolver) Location() generated.LocationResolver { return &locationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// AssetTagWhereInput returns generated.AssetTagWhereInputResolver implementation.
func (r *Resolver) AssetTagWhereInput() generated.AssetTagWhereInputResolver {
	return &assetTagWhereInputResolver{r}
}

// AssetWhereInput returns generated.AssetWhereInputResolver implementation.
func (r *Resolver) AssetWhereInput() generated.AssetWhereInputResolver {
	return &assetWhereInputResolver{r}
}

// CreateAssetInput returns generated.CreateAssetInputResolver implementation.
func (r *Resolver) CreateAssetInput() generated.CreateAssetInputResolver {
	return &createAssetInputResolver{r}
}

// CreateLocationInput returns generated.CreateLocationInputResolver implementation.
func (r *Resolver) CreateLocationInput() generated.CreateLocationInputResolver {
	return &createLocationInputResolver{r}
}

// CreateTagInput returns generated.CreateTagInputResolver implementation.
func (r *Resolver) CreateTagInput() generated.CreateTagInputResolver {
	return &createTagInputResolver{r}
}

// CreateUserInput returns generated.CreateUserInputResolver implementation.
func (r *Resolver) CreateUserInput() generated.CreateUserInputResolver {
	return &createUserInputResolver{r}
}

// LocationWhereInput returns generated.LocationWhereInputResolver implementation.
func (r *Resolver) LocationWhereInput() generated.LocationWhereInputResolver {
	return &locationWhereInputResolver{r}
}

// TagWhereInput returns generated.TagWhereInputResolver implementation.
func (r *Resolver) TagWhereInput() generated.TagWhereInputResolver { return &tagWhereInputResolver{r} }

// UpdateAssetInput returns generated.UpdateAssetInputResolver implementation.
func (r *Resolver) UpdateAssetInput() generated.UpdateAssetInputResolver {
	return &updateAssetInputResolver{r}
}

// UpdateLocationInput returns generated.UpdateLocationInputResolver implementation.
func (r *Resolver) UpdateLocationInput() generated.UpdateLocationInputResolver {
	return &updateLocationInputResolver{r}
}

// UpdateTagInput returns generated.UpdateTagInputResolver implementation.
func (r *Resolver) UpdateTagInput() generated.UpdateTagInputResolver {
	return &updateTagInputResolver{r}
}

// UpdateUserInput returns generated.UpdateUserInputResolver implementation.
func (r *Resolver) UpdateUserInput() generated.UpdateUserInputResolver {
	return &updateUserInputResolver{r}
}

// UserWhereInput returns generated.UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() generated.UserWhereInputResolver {
	return &userWhereInputResolver{r}
}

type assetResolver struct{ *Resolver }
type assetTagResolver struct{ *Resolver }
type locationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tagResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type assetTagWhereInputResolver struct{ *Resolver }
type assetWhereInputResolver struct{ *Resolver }
type createAssetInputResolver struct{ *Resolver }
type createLocationInputResolver struct{ *Resolver }
type createTagInputResolver struct{ *Resolver }
type createUserInputResolver struct{ *Resolver }
type locationWhereInputResolver struct{ *Resolver }
type tagWhereInputResolver struct{ *Resolver }
type updateAssetInputResolver struct{ *Resolver }
type updateLocationInputResolver struct{ *Resolver }
type updateTagInputResolver struct{ *Resolver }
type updateUserInputResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }
