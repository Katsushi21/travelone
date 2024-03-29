// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Katsushi21/travelone/ent/account"
	"github.com/Katsushi21/travelone/ent/mute"
	"github.com/Katsushi21/travelone/ent/predicate"
	"github.com/google/uuid"
)

// MuteUpdate is the builder for updating Mute entities.
type MuteUpdate struct {
	config
	hooks    []Hook
	mutation *MuteMutation
}

// Where appends a list predicates to the MuteUpdate builder.
func (mu *MuteUpdate) Where(ps ...predicate.Mute) *MuteUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MuteUpdate) SetUpdatedAt(t time.Time) *MuteUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetAccountID sets the "account_id" field.
func (mu *MuteUpdate) SetAccountID(u uuid.UUID) *MuteUpdate {
	mu.mutation.SetAccountID(u)
	return mu
}

// SetMuteID sets the "mute_id" field.
func (mu *MuteUpdate) SetMuteID(u uuid.UUID) *MuteUpdate {
	mu.mutation.SetMuteID(u)
	return mu
}

// SetAccount sets the "account" edge to the Account entity.
func (mu *MuteUpdate) SetAccount(a *Account) *MuteUpdate {
	return mu.SetAccountID(a.ID)
}

// SetMute sets the "mute" edge to the Account entity.
func (mu *MuteUpdate) SetMute(a *Account) *MuteUpdate {
	return mu.SetMuteID(a.ID)
}

// Mutation returns the MuteMutation object of the builder.
func (mu *MuteUpdate) Mutation() *MuteMutation {
	return mu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (mu *MuteUpdate) ClearAccount() *MuteUpdate {
	mu.mutation.ClearAccount()
	return mu
}

// ClearMute clears the "mute" edge to the Account entity.
func (mu *MuteUpdate) ClearMute() *MuteUpdate {
	mu.mutation.ClearMute()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MuteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MuteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MuteUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MuteUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MuteUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MuteUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := mute.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MuteUpdate) check() error {
	if _, ok := mu.mutation.AccountID(); mu.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mute.account"`)
	}
	if _, ok := mu.mutation.MuteID(); mu.mutation.MuteCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mute.mute"`)
	}
	return nil
}

func (mu *MuteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mute.Table,
			Columns: mute.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mute.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mute.FieldUpdatedAt,
		})
	}
	if mu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.AccountTable,
			Columns: []string{mute.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.AccountTable,
			Columns: []string{mute.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MuteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.MuteTable,
			Columns: []string{mute.MuteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MuteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.MuteTable,
			Columns: []string{mute.MuteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MuteUpdateOne is the builder for updating a single Mute entity.
type MuteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MuteMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MuteUpdateOne) SetUpdatedAt(t time.Time) *MuteUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetAccountID sets the "account_id" field.
func (muo *MuteUpdateOne) SetAccountID(u uuid.UUID) *MuteUpdateOne {
	muo.mutation.SetAccountID(u)
	return muo
}

// SetMuteID sets the "mute_id" field.
func (muo *MuteUpdateOne) SetMuteID(u uuid.UUID) *MuteUpdateOne {
	muo.mutation.SetMuteID(u)
	return muo
}

// SetAccount sets the "account" edge to the Account entity.
func (muo *MuteUpdateOne) SetAccount(a *Account) *MuteUpdateOne {
	return muo.SetAccountID(a.ID)
}

// SetMute sets the "mute" edge to the Account entity.
func (muo *MuteUpdateOne) SetMute(a *Account) *MuteUpdateOne {
	return muo.SetMuteID(a.ID)
}

// Mutation returns the MuteMutation object of the builder.
func (muo *MuteUpdateOne) Mutation() *MuteMutation {
	return muo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (muo *MuteUpdateOne) ClearAccount() *MuteUpdateOne {
	muo.mutation.ClearAccount()
	return muo
}

// ClearMute clears the "mute" edge to the Account entity.
func (muo *MuteUpdateOne) ClearMute() *MuteUpdateOne {
	muo.mutation.ClearMute()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MuteUpdateOne) Select(field string, fields ...string) *MuteUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mute entity.
func (muo *MuteUpdateOne) Save(ctx context.Context) (*Mute, error) {
	var (
		err  error
		node *Mute
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MuteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Mute)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MuteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MuteUpdateOne) SaveX(ctx context.Context) *Mute {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MuteUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MuteUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MuteUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := mute.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MuteUpdateOne) check() error {
	if _, ok := muo.mutation.AccountID(); muo.mutation.AccountCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mute.account"`)
	}
	if _, ok := muo.mutation.MuteID(); muo.mutation.MuteCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mute.mute"`)
	}
	return nil
}

func (muo *MuteUpdateOne) sqlSave(ctx context.Context) (_node *Mute, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mute.Table,
			Columns: mute.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mute.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mute.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mute.FieldID)
		for _, f := range fields {
			if !mute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mute.FieldUpdatedAt,
		})
	}
	if muo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.AccountTable,
			Columns: []string{mute.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.AccountTable,
			Columns: []string{mute.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MuteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.MuteTable,
			Columns: []string{mute.MuteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MuteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   mute.MuteTable,
			Columns: []string{mute.MuteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mute{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
