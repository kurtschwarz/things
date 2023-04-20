// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"things-api/ent/asset"
	"things-api/ent/assettag"
	"things-api/ent/predicate"
	"things-api/ent/tag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AssetTagQuery is the builder for querying AssetTag entities.
type AssetTagQuery struct {
	config
	ctx        *QueryContext
	order      []assettag.OrderOption
	inters     []Interceptor
	predicates []predicate.AssetTag
	withAsset  *AssetQuery
	withTag    *TagQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*AssetTag) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetTagQuery builder.
func (atq *AssetTagQuery) Where(ps ...predicate.AssetTag) *AssetTagQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit the number of records to be returned by this query.
func (atq *AssetTagQuery) Limit(limit int) *AssetTagQuery {
	atq.ctx.Limit = &limit
	return atq
}

// Offset to start from.
func (atq *AssetTagQuery) Offset(offset int) *AssetTagQuery {
	atq.ctx.Offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *AssetTagQuery) Unique(unique bool) *AssetTagQuery {
	atq.ctx.Unique = &unique
	return atq
}

// Order specifies how the records should be ordered.
func (atq *AssetTagQuery) Order(o ...assettag.OrderOption) *AssetTagQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// QueryAsset chains the current query on the "asset" edge.
func (atq *AssetTagQuery) QueryAsset() *AssetQuery {
	query := (&AssetClient{config: atq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(assettag.Table, assettag.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, assettag.AssetTable, assettag.AssetColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTag chains the current query on the "tag" edge.
func (atq *AssetTagQuery) QueryTag() *TagQuery {
	query := (&TagClient{config: atq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(assettag.Table, assettag.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, assettag.TagTable, assettag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AssetTag entity from the query.
// Returns a *NotFoundError when no AssetTag was found.
func (atq *AssetTagQuery) First(ctx context.Context) (*AssetTag, error) {
	nodes, err := atq.Limit(1).All(setContextOp(ctx, atq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{assettag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *AssetTagQuery) FirstX(ctx context.Context) *AssetTag {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AssetTag ID from the query.
// Returns a *NotFoundError when no AssetTag ID was found.
func (atq *AssetTagQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(1).IDs(setContextOp(ctx, atq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{assettag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *AssetTagQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AssetTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AssetTag entity is found.
// Returns a *NotFoundError when no AssetTag entities are found.
func (atq *AssetTagQuery) Only(ctx context.Context) (*AssetTag, error) {
	nodes, err := atq.Limit(2).All(setContextOp(ctx, atq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{assettag.Label}
	default:
		return nil, &NotSingularError{assettag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *AssetTagQuery) OnlyX(ctx context.Context) *AssetTag {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AssetTag ID in the query.
// Returns a *NotSingularError when more than one AssetTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *AssetTagQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(2).IDs(setContextOp(ctx, atq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{assettag.Label}
	default:
		err = &NotSingularError{assettag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *AssetTagQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AssetTags.
func (atq *AssetTagQuery) All(ctx context.Context) ([]*AssetTag, error) {
	ctx = setContextOp(ctx, atq.ctx, "All")
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AssetTag, *AssetTagQuery]()
	return withInterceptors[[]*AssetTag](ctx, atq, qr, atq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atq *AssetTagQuery) AllX(ctx context.Context) []*AssetTag {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AssetTag IDs.
func (atq *AssetTagQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if atq.ctx.Unique == nil && atq.path != nil {
		atq.Unique(true)
	}
	ctx = setContextOp(ctx, atq.ctx, "IDs")
	if err = atq.Select(assettag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *AssetTagQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *AssetTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, atq.ctx, "Count")
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atq, querierCount[*AssetTagQuery](), atq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atq *AssetTagQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *AssetTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, atq.ctx, "Exist")
	switch _, err := atq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *AssetTagQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *AssetTagQuery) Clone() *AssetTagQuery {
	if atq == nil {
		return nil
	}
	return &AssetTagQuery{
		config:     atq.config,
		ctx:        atq.ctx.Clone(),
		order:      append([]assettag.OrderOption{}, atq.order...),
		inters:     append([]Interceptor{}, atq.inters...),
		predicates: append([]predicate.AssetTag{}, atq.predicates...),
		withAsset:  atq.withAsset.Clone(),
		withTag:    atq.withTag.Clone(),
		// clone intermediate query.
		sql:  atq.sql.Clone(),
		path: atq.path,
	}
}

// WithAsset tells the query-builder to eager-load the nodes that are connected to
// the "asset" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *AssetTagQuery) WithAsset(opts ...func(*AssetQuery)) *AssetTagQuery {
	query := (&AssetClient{config: atq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atq.withAsset = query
	return atq
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *AssetTagQuery) WithTag(opts ...func(*TagQuery)) *AssetTagQuery {
	query := (&TagClient{config: atq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	atq.withTag = query
	return atq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AssetID uuid.UUID `json:"asset_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AssetTag.Query().
//		GroupBy(assettag.FieldAssetID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atq *AssetTagQuery) GroupBy(field string, fields ...string) *AssetTagGroupBy {
	atq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AssetTagGroupBy{build: atq}
	grbuild.flds = &atq.ctx.Fields
	grbuild.label = assettag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AssetID uuid.UUID `json:"asset_id,omitempty"`
//	}
//
//	client.AssetTag.Query().
//		Select(assettag.FieldAssetID).
//		Scan(ctx, &v)
func (atq *AssetTagQuery) Select(fields ...string) *AssetTagSelect {
	atq.ctx.Fields = append(atq.ctx.Fields, fields...)
	sbuild := &AssetTagSelect{AssetTagQuery: atq}
	sbuild.label = assettag.Label
	sbuild.flds, sbuild.scan = &atq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AssetTagSelect configured with the given aggregations.
func (atq *AssetTagQuery) Aggregate(fns ...AggregateFunc) *AssetTagSelect {
	return atq.Select().Aggregate(fns...)
}

func (atq *AssetTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atq); err != nil {
				return err
			}
		}
	}
	for _, f := range atq.ctx.Fields {
		if !assettag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *AssetTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AssetTag, error) {
	var (
		nodes       = []*AssetTag{}
		_spec       = atq.querySpec()
		loadedTypes = [2]bool{
			atq.withAsset != nil,
			atq.withTag != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AssetTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AssetTag{config: atq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(atq.modifiers) > 0 {
		_spec.Modifiers = atq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atq.withAsset; query != nil {
		if err := atq.loadAsset(ctx, query, nodes, nil,
			func(n *AssetTag, e *Asset) { n.Edges.Asset = e }); err != nil {
			return nil, err
		}
	}
	if query := atq.withTag; query != nil {
		if err := atq.loadTag(ctx, query, nodes, nil,
			func(n *AssetTag, e *Tag) { n.Edges.Tag = e }); err != nil {
			return nil, err
		}
	}
	for i := range atq.loadTotal {
		if err := atq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atq *AssetTagQuery) loadAsset(ctx context.Context, query *AssetQuery, nodes []*AssetTag, init func(*AssetTag), assign func(*AssetTag, *Asset)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AssetTag)
	for i := range nodes {
		fk := nodes[i].AssetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(asset.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (atq *AssetTagQuery) loadTag(ctx context.Context, query *TagQuery, nodes []*AssetTag, init func(*AssetTag), assign func(*AssetTag, *Tag)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AssetTag)
	for i := range nodes {
		fk := nodes[i].TagID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tag.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tag_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (atq *AssetTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	if len(atq.modifiers) > 0 {
		_spec.Modifiers = atq.modifiers
	}
	_spec.Node.Columns = atq.ctx.Fields
	if len(atq.ctx.Fields) > 0 {
		_spec.Unique = atq.ctx.Unique != nil && *atq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *AssetTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(assettag.Table, assettag.Columns, sqlgraph.NewFieldSpec(assettag.FieldID, field.TypeUUID))
	_spec.From = atq.sql
	if unique := atq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if atq.path != nil {
		_spec.Unique = true
	}
	if fields := atq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assettag.FieldID)
		for i := range fields {
			if fields[i] != assettag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if atq.withAsset != nil {
			_spec.Node.AddColumnOnce(assettag.FieldAssetID)
		}
		if atq.withTag != nil {
			_spec.Node.AddColumnOnce(assettag.FieldTagID)
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *AssetTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(assettag.Table)
	columns := atq.ctx.Fields
	if len(columns) == 0 {
		columns = assettag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.ctx.Unique != nil && *atq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AssetTagGroupBy is the group-by builder for AssetTag entities.
type AssetTagGroupBy struct {
	selector
	build *AssetTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *AssetTagGroupBy) Aggregate(fns ...AggregateFunc) *AssetTagGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the selector query and scans the result into the given value.
func (atgb *AssetTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, atgb.build.ctx, "GroupBy")
	if err := atgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetTagQuery, *AssetTagGroupBy](ctx, atgb.build, atgb, atgb.build.inters, v)
}

func (atgb *AssetTagGroupBy) sqlScan(ctx context.Context, root *AssetTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atgb.flds)+len(atgb.fns))
		for _, f := range *atgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AssetTagSelect is the builder for selecting fields of AssetTag entities.
type AssetTagSelect struct {
	*AssetTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ats *AssetTagSelect) Aggregate(fns ...AggregateFunc) *AssetTagSelect {
	ats.fns = append(ats.fns, fns...)
	return ats
}

// Scan applies the selector query and scans the result into the given value.
func (ats *AssetTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ats.ctx, "Select")
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetTagQuery, *AssetTagSelect](ctx, ats.AssetTagQuery, ats, ats.inters, v)
}

func (ats *AssetTagSelect) sqlScan(ctx context.Context, root *AssetTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ats.fns))
	for _, fn := range ats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
