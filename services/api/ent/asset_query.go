// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"things-api/ent/asset"
	"things-api/ent/assettag"
	"things-api/ent/location"
	"things-api/ent/predicate"
	"things-api/ent/tag"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AssetQuery is the builder for querying Asset entities.
type AssetQuery struct {
	config
	ctx                *QueryContext
	order              []asset.OrderOption
	inters             []Interceptor
	predicates         []predicate.Asset
	withParent         *AssetQuery
	withChildren       *AssetQuery
	withLocation       *LocationQuery
	withTags           *TagQuery
	withAssetTags      *AssetTagQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*Asset) error
	withNamedChildren  map[string]*AssetQuery
	withNamedTags      map[string]*TagQuery
	withNamedAssetTags map[string]*AssetTagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetQuery builder.
func (aq *AssetQuery) Where(ps ...predicate.Asset) *AssetQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AssetQuery) Limit(limit int) *AssetQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AssetQuery) Offset(offset int) *AssetQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AssetQuery) Unique(unique bool) *AssetQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AssetQuery) Order(o ...asset.OrderOption) *AssetQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryParent chains the current query on the "parent" edge.
func (aq *AssetQuery) QueryParent() *AssetQuery {
	query := (&AssetClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.ParentTable, asset.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (aq *AssetQuery) QueryChildren() *AssetQuery {
	query := (&AssetClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, asset.ChildrenTable, asset.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLocation chains the current query on the "location" edge.
func (aq *AssetQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, asset.LocationTable, asset.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTags chains the current query on the "tags" edge.
func (aq *AssetQuery) QueryTags() *TagQuery {
	query := (&TagClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, asset.TagsTable, asset.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssetTags chains the current query on the "asset_tags" edge.
func (aq *AssetQuery) QueryAssetTags() *AssetTagQuery {
	query := (&AssetTagClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, selector),
			sqlgraph.To(assettag.Table, assettag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, asset.AssetTagsTable, asset.AssetTagsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Asset entity from the query.
// Returns a *NotFoundError when no Asset was found.
func (aq *AssetQuery) First(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AssetQuery) FirstX(ctx context.Context) *Asset {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Asset ID from the query.
// Returns a *NotFoundError when no Asset ID was found.
func (aq *AssetQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AssetQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Asset entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Asset entity is found.
// Returns a *NotFoundError when no Asset entities are found.
func (aq *AssetQuery) Only(ctx context.Context) (*Asset, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asset.Label}
	default:
		return nil, &NotSingularError{asset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AssetQuery) OnlyX(ctx context.Context) *Asset {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Asset ID in the query.
// Returns a *NotSingularError when more than one Asset ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AssetQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asset.Label}
	default:
		err = &NotSingularError{asset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AssetQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Assets.
func (aq *AssetQuery) All(ctx context.Context) ([]*Asset, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Asset, *AssetQuery]()
	return withInterceptors[[]*Asset](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AssetQuery) AllX(ctx context.Context) []*Asset {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Asset IDs.
func (aq *AssetQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(asset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AssetQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AssetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AssetQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AssetQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AssetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AssetQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AssetQuery) Clone() *AssetQuery {
	if aq == nil {
		return nil
	}
	return &AssetQuery{
		config:        aq.config,
		ctx:           aq.ctx.Clone(),
		order:         append([]asset.OrderOption{}, aq.order...),
		inters:        append([]Interceptor{}, aq.inters...),
		predicates:    append([]predicate.Asset{}, aq.predicates...),
		withParent:    aq.withParent.Clone(),
		withChildren:  aq.withChildren.Clone(),
		withLocation:  aq.withLocation.Clone(),
		withTags:      aq.withTags.Clone(),
		withAssetTags: aq.withAssetTags.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithParent(opts ...func(*AssetQuery)) *AssetQuery {
	query := (&AssetClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withParent = query
	return aq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithChildren(opts ...func(*AssetQuery)) *AssetQuery {
	query := (&AssetClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withChildren = query
	return aq
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithLocation(opts ...func(*LocationQuery)) *AssetQuery {
	query := (&LocationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withLocation = query
	return aq
}

// WithTags tells the query-builder to eager-load the nodes that are connected to
// the "tags" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithTags(opts ...func(*TagQuery)) *AssetQuery {
	query := (&TagClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withTags = query
	return aq
}

// WithAssetTags tells the query-builder to eager-load the nodes that are connected to
// the "asset_tags" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithAssetTags(opts ...func(*AssetTagQuery)) *AssetQuery {
	query := (&AssetTagClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAssetTags = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ParentID uuid.UUID `json:"parent_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Asset.Query().
//		GroupBy(asset.FieldParentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AssetQuery) GroupBy(field string, fields ...string) *AssetGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AssetGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = asset.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ParentID uuid.UUID `json:"parent_id,omitempty"`
//	}
//
//	client.Asset.Query().
//		Select(asset.FieldParentID).
//		Scan(ctx, &v)
func (aq *AssetQuery) Select(fields ...string) *AssetSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AssetSelect{AssetQuery: aq}
	sbuild.label = asset.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AssetSelect configured with the given aggregations.
func (aq *AssetQuery) Aggregate(fns ...AggregateFunc) *AssetSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AssetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !asset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AssetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Asset, error) {
	var (
		nodes       = []*Asset{}
		_spec       = aq.querySpec()
		loadedTypes = [5]bool{
			aq.withParent != nil,
			aq.withChildren != nil,
			aq.withLocation != nil,
			aq.withTags != nil,
			aq.withAssetTags != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Asset).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Asset{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withParent; query != nil {
		if err := aq.loadParent(ctx, query, nodes, nil,
			func(n *Asset, e *Asset) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withChildren; query != nil {
		if err := aq.loadChildren(ctx, query, nodes,
			func(n *Asset) { n.Edges.Children = []*Asset{} },
			func(n *Asset, e *Asset) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withLocation; query != nil {
		if err := aq.loadLocation(ctx, query, nodes, nil,
			func(n *Asset, e *Location) { n.Edges.Location = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTags; query != nil {
		if err := aq.loadTags(ctx, query, nodes,
			func(n *Asset) { n.Edges.Tags = []*Tag{} },
			func(n *Asset, e *Tag) { n.Edges.Tags = append(n.Edges.Tags, e) }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAssetTags; query != nil {
		if err := aq.loadAssetTags(ctx, query, nodes,
			func(n *Asset) { n.Edges.AssetTags = []*AssetTag{} },
			func(n *Asset, e *AssetTag) { n.Edges.AssetTags = append(n.Edges.AssetTags, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedChildren {
		if err := aq.loadChildren(ctx, query, nodes,
			func(n *Asset) { n.appendNamedChildren(name) },
			func(n *Asset, e *Asset) { n.appendNamedChildren(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedTags {
		if err := aq.loadTags(ctx, query, nodes,
			func(n *Asset) { n.appendNamedTags(name) },
			func(n *Asset, e *Tag) { n.appendNamedTags(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedAssetTags {
		if err := aq.loadAssetTags(ctx, query, nodes,
			func(n *Asset) { n.appendNamedAssetTags(name) },
			func(n *Asset, e *AssetTag) { n.appendNamedAssetTags(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AssetQuery) loadParent(ctx context.Context, query *AssetQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Asset)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Asset)
	for i := range nodes {
		if nodes[i].ParentID == nil {
			continue
		}
		fk := *nodes[i].ParentID
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
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadChildren(ctx context.Context, query *AssetQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Asset)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(asset.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		if fk == nil {
			return fmt.Errorf(`foreign-key "parent_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (aq *AssetQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Location)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Asset)
	for i := range nodes {
		fk := nodes[i].LocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadTags(ctx context.Context, query *TagQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *Tag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Asset)
	nids := make(map[uuid.UUID]map[*Asset]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(asset.TagsTable)
		s.Join(joinT).On(s.C(tag.FieldID), joinT.C(asset.TagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(asset.TagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(asset.TagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Asset]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Tag](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (aq *AssetQuery) loadAssetTags(ctx context.Context, query *AssetTagQuery, nodes []*Asset, init func(*Asset), assign func(*Asset, *AssetTag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Asset)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.AssetTag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(asset.AssetTagsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssetID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "asset_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AssetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AssetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(asset.Table, asset.Columns, sqlgraph.NewFieldSpec(asset.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asset.FieldID)
		for i := range fields {
			if fields[i] != asset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withParent != nil {
			_spec.Node.AddColumnOnce(asset.FieldParentID)
		}
		if aq.withLocation != nil {
			_spec.Node.AddColumnOnce(asset.FieldLocationID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AssetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(asset.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = asset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithNamedChildren(name string, opts ...func(*AssetQuery)) *AssetQuery {
	query := (&AssetClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedChildren == nil {
		aq.withNamedChildren = make(map[string]*AssetQuery)
	}
	aq.withNamedChildren[name] = query
	return aq
}

// WithNamedTags tells the query-builder to eager-load the nodes that are connected to the "tags"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithNamedTags(name string, opts ...func(*TagQuery)) *AssetQuery {
	query := (&TagClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedTags == nil {
		aq.withNamedTags = make(map[string]*TagQuery)
	}
	aq.withNamedTags[name] = query
	return aq
}

// WithNamedAssetTags tells the query-builder to eager-load the nodes that are connected to the "asset_tags"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AssetQuery) WithNamedAssetTags(name string, opts ...func(*AssetTagQuery)) *AssetQuery {
	query := (&AssetTagClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedAssetTags == nil {
		aq.withNamedAssetTags = make(map[string]*AssetTagQuery)
	}
	aq.withNamedAssetTags[name] = query
	return aq
}

// AssetGroupBy is the group-by builder for Asset entities.
type AssetGroupBy struct {
	selector
	build *AssetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AssetGroupBy) Aggregate(fns ...AggregateFunc) *AssetGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AssetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetQuery, *AssetGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AssetGroupBy) sqlScan(ctx context.Context, root *AssetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AssetSelect is the builder for selecting fields of Asset entities.
type AssetSelect struct {
	*AssetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AssetSelect) Aggregate(fns ...AggregateFunc) *AssetSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AssetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AssetQuery, *AssetSelect](ctx, as.AssetQuery, as, as.inters, v)
}

func (as *AssetSelect) sqlScan(ctx context.Context, root *AssetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
