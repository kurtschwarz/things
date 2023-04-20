// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"things-api/ent/location"
	"things-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// Where appends a list predicates to the LocationUpdate builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetParentID sets the "parent_id" field.
func (lu *LocationUpdate) SetParentID(u uuid.UUID) *LocationUpdate {
	lu.mutation.SetParentID(u)
	return lu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableParentID(u *uuid.UUID) *LocationUpdate {
	if u != nil {
		lu.SetParentID(*u)
	}
	return lu
}

// ClearParentID clears the value of the "parent_id" field.
func (lu *LocationUpdate) ClearParentID() *LocationUpdate {
	lu.mutation.ClearParentID()
	return lu
}

// SetName sets the "name" field.
func (lu *LocationUpdate) SetName(s string) *LocationUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LocationUpdate) SetNillableName(s *string) *LocationUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// ClearName clears the value of the "name" field.
func (lu *LocationUpdate) ClearName() *LocationUpdate {
	lu.mutation.ClearName()
	return lu
}

// SetParent sets the "parent" edge to the Location entity.
func (lu *LocationUpdate) SetParent(l *Location) *LocationUpdate {
	return lu.SetParentID(l.ID)
}

// AddChildIDs adds the "children" edge to the Location entity by IDs.
func (lu *LocationUpdate) AddChildIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.AddChildIDs(ids...)
	return lu
}

// AddChildren adds the "children" edges to the Location entity.
func (lu *LocationUpdate) AddChildren(l ...*Location) *LocationUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.AddChildIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// ClearParent clears the "parent" edge to the Location entity.
func (lu *LocationUpdate) ClearParent() *LocationUpdate {
	lu.mutation.ClearParent()
	return lu
}

// ClearChildren clears all "children" edges to the Location entity.
func (lu *LocationUpdate) ClearChildren() *LocationUpdate {
	lu.mutation.ClearChildren()
	return lu
}

// RemoveChildIDs removes the "children" edge to Location entities by IDs.
func (lu *LocationUpdate) RemoveChildIDs(ids ...uuid.UUID) *LocationUpdate {
	lu.mutation.RemoveChildIDs(ids...)
	return lu
}

// RemoveChildren removes "children" edges to Location entities.
func (lu *LocationUpdate) RemoveChildren(l ...*Location) *LocationUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, LocationMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if lu.mutation.NameCleared() {
		_spec.ClearField(location.FieldName, field.TypeString)
	}
	if lu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !lu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LocationMutation
}

// SetParentID sets the "parent_id" field.
func (luo *LocationUpdateOne) SetParentID(u uuid.UUID) *LocationUpdateOne {
	luo.mutation.SetParentID(u)
	return luo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableParentID(u *uuid.UUID) *LocationUpdateOne {
	if u != nil {
		luo.SetParentID(*u)
	}
	return luo
}

// ClearParentID clears the value of the "parent_id" field.
func (luo *LocationUpdateOne) ClearParentID() *LocationUpdateOne {
	luo.mutation.ClearParentID()
	return luo
}

// SetName sets the "name" field.
func (luo *LocationUpdateOne) SetName(s string) *LocationUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LocationUpdateOne) SetNillableName(s *string) *LocationUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// ClearName clears the value of the "name" field.
func (luo *LocationUpdateOne) ClearName() *LocationUpdateOne {
	luo.mutation.ClearName()
	return luo
}

// SetParent sets the "parent" edge to the Location entity.
func (luo *LocationUpdateOne) SetParent(l *Location) *LocationUpdateOne {
	return luo.SetParentID(l.ID)
}

// AddChildIDs adds the "children" edge to the Location entity by IDs.
func (luo *LocationUpdateOne) AddChildIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.AddChildIDs(ids...)
	return luo
}

// AddChildren adds the "children" edges to the Location entity.
func (luo *LocationUpdateOne) AddChildren(l ...*Location) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.AddChildIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// ClearParent clears the "parent" edge to the Location entity.
func (luo *LocationUpdateOne) ClearParent() *LocationUpdateOne {
	luo.mutation.ClearParent()
	return luo
}

// ClearChildren clears all "children" edges to the Location entity.
func (luo *LocationUpdateOne) ClearChildren() *LocationUpdateOne {
	luo.mutation.ClearChildren()
	return luo
}

// RemoveChildIDs removes the "children" edge to Location entities by IDs.
func (luo *LocationUpdateOne) RemoveChildIDs(ids ...uuid.UUID) *LocationUpdateOne {
	luo.mutation.RemoveChildIDs(ids...)
	return luo
}

// RemoveChildren removes "children" edges to Location entities.
func (luo *LocationUpdateOne) RemoveChildren(l ...*Location) *LocationUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the LocationUpdate builder.
func (luo *LocationUpdateOne) Where(ps ...predicate.Location) *LocationUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LocationUpdateOne) Select(field string, fields ...string) *LocationUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Location entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	return withHooks[*Location, LocationMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (_node *Location, err error) {
	_spec := sqlgraph.NewUpdateSpec(location.Table, location.Columns, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Location.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, location.FieldID)
		for _, f := range fields {
			if !location.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != location.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
	}
	if luo.mutation.NameCleared() {
		_spec.ClearField(location.FieldName, field.TypeString)
	}
	if luo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !luo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Location{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
