// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"things-api/ent/asset"
	"things-api/ent/assettag"
	"things-api/ent/tag"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AssetTag is the model entity for the AssetTag schema.
type AssetTag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AssetID holds the value of the "asset_id" field.
	AssetID uuid.UUID `json:"asset_id,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID uuid.UUID `json:"tag_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AssetTagQuery when eager-loading is set.
	Edges        AssetTagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AssetTagEdges holds the relations/edges for other nodes in the graph.
type AssetTagEdges struct {
	// Asset holds the value of the asset edge.
	Asset *Asset `json:"asset,omitempty"`
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// AssetOrErr returns the Asset value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AssetTagEdges) AssetOrErr() (*Asset, error) {
	if e.loadedTypes[0] {
		if e.Asset == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: asset.Label}
		}
		return e.Asset, nil
	}
	return nil, &NotLoadedError{edge: "asset"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AssetTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[1] {
		if e.Tag == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AssetTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case assettag.FieldID, assettag.FieldAssetID, assettag.FieldTagID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AssetTag fields.
func (at *AssetTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assettag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				at.ID = *value
			}
		case assettag.FieldAssetID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field asset_id", values[i])
			} else if value != nil {
				at.AssetID = *value
			}
		case assettag.FieldTagID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value != nil {
				at.TagID = *value
			}
		default:
			at.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AssetTag.
// This includes values selected through modifiers, order, etc.
func (at *AssetTag) Value(name string) (ent.Value, error) {
	return at.selectValues.Get(name)
}

// QueryAsset queries the "asset" edge of the AssetTag entity.
func (at *AssetTag) QueryAsset() *AssetQuery {
	return NewAssetTagClient(at.config).QueryAsset(at)
}

// QueryTag queries the "tag" edge of the AssetTag entity.
func (at *AssetTag) QueryTag() *TagQuery {
	return NewAssetTagClient(at.config).QueryTag(at)
}

// Update returns a builder for updating this AssetTag.
// Note that you need to call AssetTag.Unwrap() before calling this method if this AssetTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (at *AssetTag) Update() *AssetTagUpdateOne {
	return NewAssetTagClient(at.config).UpdateOne(at)
}

// Unwrap unwraps the AssetTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (at *AssetTag) Unwrap() *AssetTag {
	_tx, ok := at.config.driver.(*txDriver)
	if !ok {
		panic("ent: AssetTag is not a transactional entity")
	}
	at.config.driver = _tx.drv
	return at
}

// String implements the fmt.Stringer.
func (at *AssetTag) String() string {
	var builder strings.Builder
	builder.WriteString("AssetTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", at.ID))
	builder.WriteString("asset_id=")
	builder.WriteString(fmt.Sprintf("%v", at.AssetID))
	builder.WriteString(", ")
	builder.WriteString("tag_id=")
	builder.WriteString(fmt.Sprintf("%v", at.TagID))
	builder.WriteByte(')')
	return builder.String()
}

// AssetTags is a parsable slice of AssetTag.
type AssetTags []*AssetTag