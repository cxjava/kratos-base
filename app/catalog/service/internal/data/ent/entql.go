// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-base/app/catalog/service/internal/data/ent/beer"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   beer.Table,
			Columns: beer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: beer.FieldID,
			},
		},
		Type: "Beer",
		Fields: map[string]*sqlgraph.FieldSpec{
			beer.FieldName: {Type: field.TypeString, Column: beer.FieldName},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (bq *BeerQuery) addPredicate(pred func(s *sql.Selector)) {
	bq.predicates = append(bq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the BeerQuery builder.
func (bq *BeerQuery) Filter() *BeerFilter {
	return &BeerFilter{config: bq.config, predicateAdder: bq}
}

// addPredicate implements the predicateAdder interface.
func (m *BeerMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the BeerMutation builder.
func (m *BeerMutation) Filter() *BeerFilter {
	return &BeerFilter{config: m.config, predicateAdder: m}
}

// BeerFilter provides a generic filtering capability at runtime for BeerQuery.
type BeerFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *BeerFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int64 predicate on the id field.
func (f *BeerFilter) WhereID(p entql.Int64P) {
	f.Where(p.Field(beer.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *BeerFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(beer.FieldName))
}