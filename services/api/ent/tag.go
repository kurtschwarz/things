// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"things-api/ent/tag"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Asset holds the value of the asset edge.
	Asset []*Asset `json:"asset,omitempty"`
	// AssetTag holds the value of the asset_tag edge.
	AssetTag []*AssetTag `json:"asset_tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedAsset    map[string][]*Asset
	namedAssetTag map[string][]*AssetTag
}

// AssetOrErr returns the Asset value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) AssetOrErr() ([]*Asset, error) {
	if e.loadedTypes[0] {
		return e.Asset, nil
	}
	return nil, &NotLoadedError{edge: "asset"}
}

// AssetTagOrErr returns the AssetTag value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) AssetTagOrErr() ([]*AssetTag, error) {
	if e.loadedTypes[1] {
		return e.AssetTag, nil
	}
	return nil, &NotLoadedError{edge: "asset_tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldName:
			values[i] = new(sql.NullString)
		case tag.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryAsset queries the "asset" edge of the Tag entity.
func (t *Tag) QueryAsset() *AssetQuery {
	return NewTagClient(t.config).QueryAsset(t)
}

// QueryAssetTag queries the "asset_tag" edge of the Tag entity.
func (t *Tag) QueryAssetTag() *AssetTagQuery {
	return NewTagClient(t.config).QueryAssetTag(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedAsset returns the Asset named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Tag) NamedAsset(name string) ([]*Asset, error) {
	if t.Edges.namedAsset == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedAsset[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Tag) appendNamedAsset(name string, edges ...*Asset) {
	if t.Edges.namedAsset == nil {
		t.Edges.namedAsset = make(map[string][]*Asset)
	}
	if len(edges) == 0 {
		t.Edges.namedAsset[name] = []*Asset{}
	} else {
		t.Edges.namedAsset[name] = append(t.Edges.namedAsset[name], edges...)
	}
}

// NamedAssetTag returns the AssetTag named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Tag) NamedAssetTag(name string) ([]*AssetTag, error) {
	if t.Edges.namedAssetTag == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedAssetTag[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Tag) appendNamedAssetTag(name string, edges ...*AssetTag) {
	if t.Edges.namedAssetTag == nil {
		t.Edges.namedAssetTag = make(map[string][]*AssetTag)
	}
	if len(edges) == 0 {
		t.Edges.namedAssetTag[name] = []*AssetTag{}
	} else {
		t.Edges.namedAssetTag[name] = append(t.Edges.namedAssetTag[name], edges...)
	}
}

// Tags is a parsable slice of Tag.
type Tags []*Tag
