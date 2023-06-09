// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-base/app/catalog/service/internal/data/ent/beer"
	"kratos-base/app/catalog/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BeerDelete is the builder for deleting a Beer entity.
type BeerDelete struct {
	config
	hooks    []Hook
	mutation *BeerMutation
}

// Where appends a list predicates to the BeerDelete builder.
func (bd *BeerDelete) Where(ps ...predicate.Beer) *BeerDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BeerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BeerDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BeerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(beer.Table, sqlgraph.NewFieldSpec(beer.FieldID, field.TypeInt64))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BeerDeleteOne is the builder for deleting a single Beer entity.
type BeerDeleteOne struct {
	bd *BeerDelete
}

// Where appends a list predicates to the BeerDelete builder.
func (bdo *BeerDeleteOne) Where(ps ...predicate.Beer) *BeerDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BeerDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{beer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BeerDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
