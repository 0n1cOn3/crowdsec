// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/bouncer"
)

// BouncerCreate is the builder for creating a Bouncer entity.
type BouncerCreate struct {
	config
	mutation *BouncerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BouncerCreate) SetCreatedAt(t time.Time) *BouncerCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableCreatedAt(t *time.Time) *BouncerCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BouncerCreate) SetUpdatedAt(t time.Time) *BouncerCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableUpdatedAt(t *time.Time) *BouncerCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetName sets the "name" field.
func (bc *BouncerCreate) SetName(s string) *BouncerCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetAPIKey sets the "api_key" field.
func (bc *BouncerCreate) SetAPIKey(s string) *BouncerCreate {
	bc.mutation.SetAPIKey(s)
	return bc
}

// SetRevoked sets the "revoked" field.
func (bc *BouncerCreate) SetRevoked(b bool) *BouncerCreate {
	bc.mutation.SetRevoked(b)
	return bc
}

// SetIPAddress sets the "ip_address" field.
func (bc *BouncerCreate) SetIPAddress(s string) *BouncerCreate {
	bc.mutation.SetIPAddress(s)
	return bc
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableIPAddress(s *string) *BouncerCreate {
	if s != nil {
		bc.SetIPAddress(*s)
	}
	return bc
}

// SetType sets the "type" field.
func (bc *BouncerCreate) SetType(s string) *BouncerCreate {
	bc.mutation.SetType(s)
	return bc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableType(s *string) *BouncerCreate {
	if s != nil {
		bc.SetType(*s)
	}
	return bc
}

// SetVersion sets the "version" field.
func (bc *BouncerCreate) SetVersion(s string) *BouncerCreate {
	bc.mutation.SetVersion(s)
	return bc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableVersion(s *string) *BouncerCreate {
	if s != nil {
		bc.SetVersion(*s)
	}
	return bc
}

// SetUntil sets the "until" field.
func (bc *BouncerCreate) SetUntil(t time.Time) *BouncerCreate {
	bc.mutation.SetUntil(t)
	return bc
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableUntil(t *time.Time) *BouncerCreate {
	if t != nil {
		bc.SetUntil(*t)
	}
	return bc
}

// SetLastPull sets the "last_pull" field.
func (bc *BouncerCreate) SetLastPull(t time.Time) *BouncerCreate {
	bc.mutation.SetLastPull(t)
	return bc
}

// SetNillableLastPull sets the "last_pull" field if the given value is not nil.
func (bc *BouncerCreate) SetNillableLastPull(t *time.Time) *BouncerCreate {
	if t != nil {
		bc.SetLastPull(*t)
	}
	return bc
}

// Mutation returns the BouncerMutation object of the builder.
func (bc *BouncerCreate) Mutation() *BouncerMutation {
	return bc.mutation
}

// Save creates the Bouncer in the database.
func (bc *BouncerCreate) Save(ctx context.Context) (*Bouncer, error) {
	var (
		err  error
		node *Bouncer
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BouncerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BouncerCreate) SaveX(ctx context.Context) *Bouncer {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BouncerCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BouncerCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BouncerCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := bouncer.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := bouncer.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.IPAddress(); !ok {
		v := bouncer.DefaultIPAddress
		bc.mutation.SetIPAddress(v)
	}
	if _, ok := bc.mutation.Until(); !ok {
		v := bouncer.DefaultUntil()
		bc.mutation.SetUntil(v)
	}
	if _, ok := bc.mutation.LastPull(); !ok {
		v := bouncer.DefaultLastPull()
		bc.mutation.SetLastPull(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BouncerCreate) check() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := bc.mutation.APIKey(); !ok {
		return &ValidationError{Name: "api_key", err: errors.New(`ent: missing required field "api_key"`)}
	}
	if _, ok := bc.mutation.Revoked(); !ok {
		return &ValidationError{Name: "revoked", err: errors.New(`ent: missing required field "revoked"`)}
	}
	if _, ok := bc.mutation.LastPull(); !ok {
		return &ValidationError{Name: "last_pull", err: errors.New(`ent: missing required field "last_pull"`)}
	}
	return nil
}

func (bc *BouncerCreate) sqlSave(ctx context.Context) (*Bouncer, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BouncerCreate) createSpec() (*Bouncer, *sqlgraph.CreateSpec) {
	var (
		_node = &Bouncer{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bouncer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bouncer.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bouncer.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bouncer.FieldUpdatedAt,
		})
		_node.UpdatedAt = &value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bouncer.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.APIKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bouncer.FieldAPIKey,
		})
		_node.APIKey = value
	}
	if value, ok := bc.mutation.Revoked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: bouncer.FieldRevoked,
		})
		_node.Revoked = value
	}
	if value, ok := bc.mutation.IPAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bouncer.FieldIPAddress,
		})
		_node.IPAddress = value
	}
	if value, ok := bc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bouncer.FieldType,
		})
		_node.Type = value
	}
	if value, ok := bc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bouncer.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := bc.mutation.Until(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bouncer.FieldUntil,
		})
		_node.Until = value
	}
	if value, ok := bc.mutation.LastPull(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bouncer.FieldLastPull,
		})
		_node.LastPull = value
	}
	return _node, _spec
}

// BouncerCreateBulk is the builder for creating many Bouncer entities in bulk.
type BouncerCreateBulk struct {
	config
	builders []*BouncerCreate
}

// Save creates the Bouncer entities in the database.
func (bcb *BouncerCreateBulk) Save(ctx context.Context) ([]*Bouncer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bouncer, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BouncerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BouncerCreateBulk) SaveX(ctx context.Context) []*Bouncer {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BouncerCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BouncerCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
