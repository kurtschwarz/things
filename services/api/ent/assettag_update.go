// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"things-api/ent/asset"
	"things-api/ent/assettag"
	"things-api/ent/predicate"
	"things-api/ent/tag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AssetTagUpdate is the builder for updating AssetTag entities.
type AssetTagUpdate struct {
	config
	hooks    []Hook
	mutation *AssetTagMutation
}

// Where appends a list predicates to the AssetTagUpdate builder.
func (atu *AssetTagUpdate) Where(ps ...predicate.AssetTag) *AssetTagUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetAssetID sets the "asset_id" field.
func (atu *AssetTagUpdate) SetAssetID(u uuid.UUID) *AssetTagUpdate {
	atu.mutation.SetAssetID(u)
	return atu
}

// SetTagID sets the "tag_id" field.
func (atu *AssetTagUpdate) SetTagID(u uuid.UUID) *AssetTagUpdate {
	atu.mutation.SetTagID(u)
	return atu
}

// SetAsset sets the "asset" edge to the Asset entity.
func (atu *AssetTagUpdate) SetAsset(a *Asset) *AssetTagUpdate {
	return atu.SetAssetID(a.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (atu *AssetTagUpdate) SetTag(t *Tag) *AssetTagUpdate {
	return atu.SetTagID(t.ID)
}

// Mutation returns the AssetTagMutation object of the builder.
func (atu *AssetTagUpdate) Mutation() *AssetTagMutation {
	return atu.mutation
}

// ClearAsset clears the "asset" edge to the Asset entity.
func (atu *AssetTagUpdate) ClearAsset() *AssetTagUpdate {
	atu.mutation.ClearAsset()
	return atu
}

// ClearTag clears the "tag" edge to the Tag entity.
func (atu *AssetTagUpdate) ClearTag() *AssetTagUpdate {
	atu.mutation.ClearTag()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AssetTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AssetTagMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AssetTagUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AssetTagUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AssetTagUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AssetTagUpdate) check() error {
	if _, ok := atu.mutation.AssetID(); atu.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetTag.asset"`)
	}
	if _, ok := atu.mutation.TagID(); atu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetTag.tag"`)
	}
	return nil
}

func (atu *AssetTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(assettag.Table, assettag.Columns, sqlgraph.NewFieldSpec(assettag.FieldID, field.TypeUUID))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.AssetTable,
			Columns: []string{assettag.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.AssetTable,
			Columns: []string{assettag.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.TagTable,
			Columns: []string{assettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.TagTable,
			Columns: []string{assettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assettag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AssetTagUpdateOne is the builder for updating a single AssetTag entity.
type AssetTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AssetTagMutation
}

// SetAssetID sets the "asset_id" field.
func (atuo *AssetTagUpdateOne) SetAssetID(u uuid.UUID) *AssetTagUpdateOne {
	atuo.mutation.SetAssetID(u)
	return atuo
}

// SetTagID sets the "tag_id" field.
func (atuo *AssetTagUpdateOne) SetTagID(u uuid.UUID) *AssetTagUpdateOne {
	atuo.mutation.SetTagID(u)
	return atuo
}

// SetAsset sets the "asset" edge to the Asset entity.
func (atuo *AssetTagUpdateOne) SetAsset(a *Asset) *AssetTagUpdateOne {
	return atuo.SetAssetID(a.ID)
}

// SetTag sets the "tag" edge to the Tag entity.
func (atuo *AssetTagUpdateOne) SetTag(t *Tag) *AssetTagUpdateOne {
	return atuo.SetTagID(t.ID)
}

// Mutation returns the AssetTagMutation object of the builder.
func (atuo *AssetTagUpdateOne) Mutation() *AssetTagMutation {
	return atuo.mutation
}

// ClearAsset clears the "asset" edge to the Asset entity.
func (atuo *AssetTagUpdateOne) ClearAsset() *AssetTagUpdateOne {
	atuo.mutation.ClearAsset()
	return atuo
}

// ClearTag clears the "tag" edge to the Tag entity.
func (atuo *AssetTagUpdateOne) ClearTag() *AssetTagUpdateOne {
	atuo.mutation.ClearTag()
	return atuo
}

// Where appends a list predicates to the AssetTagUpdate builder.
func (atuo *AssetTagUpdateOne) Where(ps ...predicate.AssetTag) *AssetTagUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AssetTagUpdateOne) Select(field string, fields ...string) *AssetTagUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AssetTag entity.
func (atuo *AssetTagUpdateOne) Save(ctx context.Context) (*AssetTag, error) {
	return withHooks[*AssetTag, AssetTagMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AssetTagUpdateOne) SaveX(ctx context.Context) *AssetTag {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AssetTagUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AssetTagUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AssetTagUpdateOne) check() error {
	if _, ok := atuo.mutation.AssetID(); atuo.mutation.AssetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetTag.asset"`)
	}
	if _, ok := atuo.mutation.TagID(); atuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AssetTag.tag"`)
	}
	return nil
}

func (atuo *AssetTagUpdateOne) sqlSave(ctx context.Context) (_node *AssetTag, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(assettag.Table, assettag.Columns, sqlgraph.NewFieldSpec(assettag.FieldID, field.TypeUUID))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AssetTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assettag.FieldID)
		for _, f := range fields {
			if !assettag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != assettag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.AssetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.AssetTable,
			Columns: []string{assettag.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.AssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.AssetTable,
			Columns: []string{assettag.AssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.TagTable,
			Columns: []string{assettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   assettag.TagTable,
			Columns: []string{assettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AssetTag{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{assettag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}