// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"things-api/ent/location"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Location is the model entity for the Location schema.
type Location struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID uuid.UUID `json:"parent_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges        LocationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Location `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Location `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedChildren map[string][]*Location
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LocationEdges) ParentOrErr() (*Location, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ChildrenOrErr() ([]*Location, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case location.FieldName:
			values[i] = new(sql.NullString)
		case location.FieldID, location.FieldParentID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case location.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case location.FieldParentID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value != nil {
				l.ParentID = *value
			}
		case location.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Location.
// This includes values selected through modifiers, order, etc.
func (l *Location) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Location entity.
func (l *Location) QueryParent() *LocationQuery {
	return NewLocationClient(l.config).QueryParent(l)
}

// QueryChildren queries the "children" edge of the Location entity.
func (l *Location) QueryChildren() *LocationQuery {
	return NewLocationClient(l.config).QueryChildren(l)
}

// Update returns a builder for updating this Location.
// Note that you need to call Location.Unwrap() before calling this method if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return NewLocationClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Location entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", l.ParentID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Location) NamedChildren(name string) ([]*Location, error) {
	if l.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Location) appendNamedChildren(name string, edges ...*Location) {
	if l.Edges.namedChildren == nil {
		l.Edges.namedChildren = make(map[string][]*Location)
	}
	if len(edges) == 0 {
		l.Edges.namedChildren[name] = []*Location{}
	} else {
		l.Edges.namedChildren[name] = append(l.Edges.namedChildren[name], edges...)
	}
}

// Locations is a parsable slice of Location.
type Locations []*Location