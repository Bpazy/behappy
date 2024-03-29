// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Bpazy/behappy/ent/hero"
	"github.com/Bpazy/behappy/ent/predicate"
)

// HeroDelete is the builder for deleting a Hero entity.
type HeroDelete struct {
	config
	hooks    []Hook
	mutation *HeroMutation
}

// Where appends a list predicates to the HeroDelete builder.
func (hd *HeroDelete) Where(ps ...predicate.Hero) *HeroDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HeroDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hd.hooks) == 0 {
		affected, err = hd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HeroMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hd.mutation = mutation
			affected, err = hd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hd.hooks) - 1; i >= 0; i-- {
			if hd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HeroDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HeroDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hero.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hero.FieldID,
			},
		},
	}
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HeroDeleteOne is the builder for deleting a single Hero entity.
type HeroDeleteOne struct {
	hd *HeroDelete
}

// Exec executes the deletion query.
func (hdo *HeroDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hero.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HeroDeleteOne) ExecX(ctx context.Context) {
	hdo.hd.ExecX(ctx)
}
