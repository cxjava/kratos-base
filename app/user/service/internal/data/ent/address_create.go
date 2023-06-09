// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-base/app/user/service/internal/data/ent/address"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetStreet sets the "street" field.
func (ac *AddressCreate) SetStreet(i int32) *AddressCreate {
	ac.mutation.SetStreet(i)
	return ac
}

// SetJoinTime sets the "joinTime" field.
func (ac *AddressCreate) SetJoinTime(t time.Time) *AddressCreate {
	ac.mutation.SetJoinTime(t)
	return ac
}

// SetHost sets the "host" field.
func (ac *AddressCreate) SetHost(s string) *AddressCreate {
	ac.mutation.SetHost(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AddressCreate) SetID(i int64) *AddressCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.Street(); !ok {
		return &ValidationError{Name: "street", err: errors.New(`ent: missing required field "Address.street"`)}
	}
	if _, ok := ac.mutation.JoinTime(); !ok {
		return &ValidationError{Name: "joinTime", err: errors.New(`ent: missing required field "Address.joinTime"`)}
	}
	if _, ok := ac.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Address.host"`)}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(address.Table, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeInt32, value)
		_node.Street = value
	}
	if value, ok := ac.mutation.JoinTime(); ok {
		_spec.SetField(address.FieldJoinTime, field.TypeTime, value)
		_node.JoinTime = value
	}
	if value, ok := ac.mutation.Host(); ok {
		_spec.SetField(address.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Address.Create().
//		SetStreet(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AddressUpsert) {
//			SetStreet(v+v).
//		}).
//		Exec(ctx)
func (ac *AddressCreate) OnConflict(opts ...sql.ConflictOption) *AddressUpsertOne {
	ac.conflict = opts
	return &AddressUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Address.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AddressCreate) OnConflictColumns(columns ...string) *AddressUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AddressUpsertOne{
		create: ac,
	}
}

type (
	// AddressUpsertOne is the builder for "upsert"-ing
	//  one Address node.
	AddressUpsertOne struct {
		create *AddressCreate
	}

	// AddressUpsert is the "OnConflict" setter.
	AddressUpsert struct {
		*sql.UpdateSet
	}
)

// SetStreet sets the "street" field.
func (u *AddressUpsert) SetStreet(v int32) *AddressUpsert {
	u.Set(address.FieldStreet, v)
	return u
}

// UpdateStreet sets the "street" field to the value that was provided on create.
func (u *AddressUpsert) UpdateStreet() *AddressUpsert {
	u.SetExcluded(address.FieldStreet)
	return u
}

// AddStreet adds v to the "street" field.
func (u *AddressUpsert) AddStreet(v int32) *AddressUpsert {
	u.Add(address.FieldStreet, v)
	return u
}

// SetJoinTime sets the "joinTime" field.
func (u *AddressUpsert) SetJoinTime(v time.Time) *AddressUpsert {
	u.Set(address.FieldJoinTime, v)
	return u
}

// UpdateJoinTime sets the "joinTime" field to the value that was provided on create.
func (u *AddressUpsert) UpdateJoinTime() *AddressUpsert {
	u.SetExcluded(address.FieldJoinTime)
	return u
}

// SetHost sets the "host" field.
func (u *AddressUpsert) SetHost(v string) *AddressUpsert {
	u.Set(address.FieldHost, v)
	return u
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *AddressUpsert) UpdateHost() *AddressUpsert {
	u.SetExcluded(address.FieldHost)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Address.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(address.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AddressUpsertOne) UpdateNewValues() *AddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(address.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Address.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AddressUpsertOne) Ignore() *AddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AddressUpsertOne) DoNothing() *AddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AddressCreate.OnConflict
// documentation for more info.
func (u *AddressUpsertOne) Update(set func(*AddressUpsert)) *AddressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AddressUpsert{UpdateSet: update})
	}))
	return u
}

// SetStreet sets the "street" field.
func (u *AddressUpsertOne) SetStreet(v int32) *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.SetStreet(v)
	})
}

// AddStreet adds v to the "street" field.
func (u *AddressUpsertOne) AddStreet(v int32) *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.AddStreet(v)
	})
}

// UpdateStreet sets the "street" field to the value that was provided on create.
func (u *AddressUpsertOne) UpdateStreet() *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateStreet()
	})
}

// SetJoinTime sets the "joinTime" field.
func (u *AddressUpsertOne) SetJoinTime(v time.Time) *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.SetJoinTime(v)
	})
}

// UpdateJoinTime sets the "joinTime" field to the value that was provided on create.
func (u *AddressUpsertOne) UpdateJoinTime() *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateJoinTime()
	})
}

// SetHost sets the "host" field.
func (u *AddressUpsertOne) SetHost(v string) *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.SetHost(v)
	})
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *AddressUpsertOne) UpdateHost() *AddressUpsertOne {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateHost()
	})
}

// Exec executes the query.
func (u *AddressUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AddressCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AddressUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AddressUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AddressUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
	conflict []sql.ConflictOption
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Address.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AddressUpsert) {
//			SetStreet(v+v).
//		}).
//		Exec(ctx)
func (acb *AddressCreateBulk) OnConflict(opts ...sql.ConflictOption) *AddressUpsertBulk {
	acb.conflict = opts
	return &AddressUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Address.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AddressCreateBulk) OnConflictColumns(columns ...string) *AddressUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AddressUpsertBulk{
		create: acb,
	}
}

// AddressUpsertBulk is the builder for "upsert"-ing
// a bulk of Address nodes.
type AddressUpsertBulk struct {
	create *AddressCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Address.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(address.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AddressUpsertBulk) UpdateNewValues() *AddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(address.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Address.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AddressUpsertBulk) Ignore() *AddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AddressUpsertBulk) DoNothing() *AddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AddressCreateBulk.OnConflict
// documentation for more info.
func (u *AddressUpsertBulk) Update(set func(*AddressUpsert)) *AddressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AddressUpsert{UpdateSet: update})
	}))
	return u
}

// SetStreet sets the "street" field.
func (u *AddressUpsertBulk) SetStreet(v int32) *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.SetStreet(v)
	})
}

// AddStreet adds v to the "street" field.
func (u *AddressUpsertBulk) AddStreet(v int32) *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.AddStreet(v)
	})
}

// UpdateStreet sets the "street" field to the value that was provided on create.
func (u *AddressUpsertBulk) UpdateStreet() *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateStreet()
	})
}

// SetJoinTime sets the "joinTime" field.
func (u *AddressUpsertBulk) SetJoinTime(v time.Time) *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.SetJoinTime(v)
	})
}

// UpdateJoinTime sets the "joinTime" field to the value that was provided on create.
func (u *AddressUpsertBulk) UpdateJoinTime() *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateJoinTime()
	})
}

// SetHost sets the "host" field.
func (u *AddressUpsertBulk) SetHost(v string) *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.SetHost(v)
	})
}

// UpdateHost sets the "host" field to the value that was provided on create.
func (u *AddressUpsertBulk) UpdateHost() *AddressUpsertBulk {
	return u.Update(func(s *AddressUpsert) {
		s.UpdateHost()
	})
}

// Exec executes the query.
func (u *AddressUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AddressCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AddressCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AddressUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
