// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Katsushi21/traveling_alone/ent/marker"
	"github.com/Katsushi21/traveling_alone/ent/post"
	"github.com/google/uuid"
)

// MarkerCreate is the builder for creating a Marker entity.
type MarkerCreate struct {
	config
	mutation *MarkerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MarkerCreate) SetCreatedAt(t time.Time) *MarkerCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MarkerCreate) SetNillableCreatedAt(t *time.Time) *MarkerCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MarkerCreate) SetUpdatedAt(t time.Time) *MarkerCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MarkerCreate) SetNillableUpdatedAt(t *time.Time) *MarkerCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetTitle sets the "title" field.
func (mc *MarkerCreate) SetTitle(s string) *MarkerCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetLat sets the "lat" field.
func (mc *MarkerCreate) SetLat(s string) *MarkerCreate {
	mc.mutation.SetLat(s)
	return mc
}

// SetLng sets the "lng" field.
func (mc *MarkerCreate) SetLng(s string) *MarkerCreate {
	mc.mutation.SetLng(s)
	return mc
}

// SetPostID sets the "post_id" field.
func (mc *MarkerCreate) SetPostID(u uuid.UUID) *MarkerCreate {
	mc.mutation.SetPostID(u)
	return mc
}

// SetID sets the "id" field.
func (mc *MarkerCreate) SetID(u uuid.UUID) *MarkerCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MarkerCreate) SetNillableID(u *uuid.UUID) *MarkerCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetPost sets the "post" edge to the Post entity.
func (mc *MarkerCreate) SetPost(p *Post) *MarkerCreate {
	return mc.SetPostID(p.ID)
}

// Mutation returns the MarkerMutation object of the builder.
func (mc *MarkerCreate) Mutation() *MarkerMutation {
	return mc.mutation
}

// Save creates the Marker in the database.
func (mc *MarkerCreate) Save(ctx context.Context) (*Marker, error) {
	var (
		err  error
		node *Marker
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MarkerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Marker)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MarkerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MarkerCreate) SaveX(ctx context.Context) *Marker {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MarkerCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MarkerCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MarkerCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := marker.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := marker.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := marker.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MarkerCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Marker.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Marker.updated_at"`)}
	}
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Marker.title"`)}
	}
	if _, ok := mc.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "Marker.lat"`)}
	}
	if _, ok := mc.mutation.Lng(); !ok {
		return &ValidationError{Name: "lng", err: errors.New(`ent: missing required field "Marker.lng"`)}
	}
	if _, ok := mc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Marker.post_id"`)}
	}
	if _, ok := mc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New(`ent: missing required edge "Marker.post"`)}
	}
	return nil
}

func (mc *MarkerCreate) sqlSave(ctx context.Context) (*Marker, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (mc *MarkerCreate) createSpec() (*Marker, *sqlgraph.CreateSpec) {
	var (
		_node = &Marker{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: marker.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: marker.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: marker.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: marker.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: marker.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.Lat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: marker.FieldLat,
		})
		_node.Lat = value
	}
	if value, ok := mc.mutation.Lng(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: marker.FieldLng,
		})
		_node.Lng = value
	}
	if nodes := mc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   marker.PostTable,
			Columns: []string{marker.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MarkerCreateBulk is the builder for creating many Marker entities in bulk.
type MarkerCreateBulk struct {
	config
	builders []*MarkerCreate
}

// Save creates the Marker entities in the database.
func (mcb *MarkerCreateBulk) Save(ctx context.Context) ([]*Marker, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Marker, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MarkerMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MarkerCreateBulk) SaveX(ctx context.Context) []*Marker {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MarkerCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MarkerCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
