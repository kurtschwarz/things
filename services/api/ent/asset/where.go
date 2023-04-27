// Code generated by ent, DO NOT EDIT.

package asset

import (
	"things-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldID, id))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldParentID, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldLocationID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldName, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldQuantity, v))
}

// ModelNumber applies equality check predicate on the "model_number" field. It's identical to ModelNumberEQ.
func ModelNumber(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldModelNumber, v))
}

// SerialNumber applies equality check predicate on the "serial_number" field. It's identical to SerialNumberEQ.
func SerialNumber(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldSerialNumber, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Asset {
	return predicate.Asset(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Asset {
	return predicate.Asset(sql.FieldNotNull(FieldParentID))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...uuid.UUID) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldLocationID, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldName, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldQuantity, v))
}

// ModelNumberEQ applies the EQ predicate on the "model_number" field.
func ModelNumberEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldModelNumber, v))
}

// ModelNumberNEQ applies the NEQ predicate on the "model_number" field.
func ModelNumberNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldModelNumber, v))
}

// ModelNumberIn applies the In predicate on the "model_number" field.
func ModelNumberIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldModelNumber, vs...))
}

// ModelNumberNotIn applies the NotIn predicate on the "model_number" field.
func ModelNumberNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldModelNumber, vs...))
}

// ModelNumberGT applies the GT predicate on the "model_number" field.
func ModelNumberGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldModelNumber, v))
}

// ModelNumberGTE applies the GTE predicate on the "model_number" field.
func ModelNumberGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldModelNumber, v))
}

// ModelNumberLT applies the LT predicate on the "model_number" field.
func ModelNumberLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldModelNumber, v))
}

// ModelNumberLTE applies the LTE predicate on the "model_number" field.
func ModelNumberLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldModelNumber, v))
}

// ModelNumberContains applies the Contains predicate on the "model_number" field.
func ModelNumberContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldModelNumber, v))
}

// ModelNumberHasPrefix applies the HasPrefix predicate on the "model_number" field.
func ModelNumberHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldModelNumber, v))
}

// ModelNumberHasSuffix applies the HasSuffix predicate on the "model_number" field.
func ModelNumberHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldModelNumber, v))
}

// ModelNumberIsNil applies the IsNil predicate on the "model_number" field.
func ModelNumberIsNil() predicate.Asset {
	return predicate.Asset(sql.FieldIsNull(FieldModelNumber))
}

// ModelNumberNotNil applies the NotNil predicate on the "model_number" field.
func ModelNumberNotNil() predicate.Asset {
	return predicate.Asset(sql.FieldNotNull(FieldModelNumber))
}

// ModelNumberEqualFold applies the EqualFold predicate on the "model_number" field.
func ModelNumberEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldModelNumber, v))
}

// ModelNumberContainsFold applies the ContainsFold predicate on the "model_number" field.
func ModelNumberContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldModelNumber, v))
}

// SerialNumberEQ applies the EQ predicate on the "serial_number" field.
func SerialNumberEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldSerialNumber, v))
}

// SerialNumberNEQ applies the NEQ predicate on the "serial_number" field.
func SerialNumberNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldSerialNumber, v))
}

// SerialNumberIn applies the In predicate on the "serial_number" field.
func SerialNumberIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldSerialNumber, vs...))
}

// SerialNumberNotIn applies the NotIn predicate on the "serial_number" field.
func SerialNumberNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldSerialNumber, vs...))
}

// SerialNumberGT applies the GT predicate on the "serial_number" field.
func SerialNumberGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldSerialNumber, v))
}

// SerialNumberGTE applies the GTE predicate on the "serial_number" field.
func SerialNumberGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldSerialNumber, v))
}

// SerialNumberLT applies the LT predicate on the "serial_number" field.
func SerialNumberLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldSerialNumber, v))
}

// SerialNumberLTE applies the LTE predicate on the "serial_number" field.
func SerialNumberLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldSerialNumber, v))
}

// SerialNumberContains applies the Contains predicate on the "serial_number" field.
func SerialNumberContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldSerialNumber, v))
}

// SerialNumberHasPrefix applies the HasPrefix predicate on the "serial_number" field.
func SerialNumberHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldSerialNumber, v))
}

// SerialNumberHasSuffix applies the HasSuffix predicate on the "serial_number" field.
func SerialNumberHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldSerialNumber, v))
}

// SerialNumberIsNil applies the IsNil predicate on the "serial_number" field.
func SerialNumberIsNil() predicate.Asset {
	return predicate.Asset(sql.FieldIsNull(FieldSerialNumber))
}

// SerialNumberNotNil applies the NotNil predicate on the "serial_number" field.
func SerialNumberNotNil() predicate.Asset {
	return predicate.Asset(sql.FieldNotNull(FieldSerialNumber))
}

// SerialNumberEqualFold applies the EqualFold predicate on the "serial_number" field.
func SerialNumberEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldSerialNumber, v))
}

// SerialNumberContainsFold applies the ContainsFold predicate on the "serial_number" field.
func SerialNumberContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldSerialNumber, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssetTags applies the HasEdge predicate on the "asset_tags" edge.
func HasAssetTags() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AssetTagsTable, AssetTagsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetTagsWith applies the HasEdge predicate on the "asset_tags" edge with a given conditions (other predicates).
func HasAssetTagsWith(preds ...predicate.AssetTag) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := newAssetTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		p(s.Not())
	})
}
