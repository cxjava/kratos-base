// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-base/internal/data/ent/address"
	"kratos-base/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks     []Hook
	mutation  *AddressMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AddressUpdate builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetStreet sets the "street" field.
func (au *AddressUpdate) SetStreet(i int32) *AddressUpdate {
	au.mutation.ResetStreet()
	au.mutation.SetStreet(i)
	return au
}

// AddStreet adds i to the "street" field.
func (au *AddressUpdate) AddStreet(i int32) *AddressUpdate {
	au.mutation.AddStreet(i)
	return au
}

// SetJoinTime sets the "joinTime" field.
func (au *AddressUpdate) SetJoinTime(t time.Time) *AddressUpdate {
	au.mutation.SetJoinTime(t)
	return au
}

// SetHost sets the "host" field.
func (au *AddressUpdate) SetHost(s string) *AddressUpdate {
	au.mutation.SetHost(s)
	return au
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AddressUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AddressUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeInt32, value)
	}
	if value, ok := au.mutation.AddedStreet(); ok {
		_spec.AddField(address.FieldStreet, field.TypeInt32, value)
	}
	if value, ok := au.mutation.JoinTime(); ok {
		_spec.SetField(address.FieldJoinTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.Host(); ok {
		_spec.SetField(address.FieldHost, field.TypeString, value)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AddressMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStreet sets the "street" field.
func (auo *AddressUpdateOne) SetStreet(i int32) *AddressUpdateOne {
	auo.mutation.ResetStreet()
	auo.mutation.SetStreet(i)
	return auo
}

// AddStreet adds i to the "street" field.
func (auo *AddressUpdateOne) AddStreet(i int32) *AddressUpdateOne {
	auo.mutation.AddStreet(i)
	return auo
}

// SetJoinTime sets the "joinTime" field.
func (auo *AddressUpdateOne) SetJoinTime(t time.Time) *AddressUpdateOne {
	auo.mutation.SetJoinTime(t)
	return auo
}

// SetHost sets the "host" field.
func (auo *AddressUpdateOne) SetHost(s string) *AddressUpdateOne {
	auo.mutation.SetHost(s)
	return auo
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (auo *AddressUpdateOne) Where(ps ...predicate.Address) *AddressUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AddressUpdateOne) Select(field string, fields ...string) *AddressUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Address entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AddressUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AddressUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Address.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, address.FieldID)
		for _, f := range fields {
			if !address.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != address.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.AddedStreet(); ok {
		_spec.AddField(address.FieldStreet, field.TypeInt32, value)
	}
	if value, ok := auo.mutation.JoinTime(); ok {
		_spec.SetField(address.FieldJoinTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Host(); ok {
		_spec.SetField(address.FieldHost, field.TypeString, value)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
