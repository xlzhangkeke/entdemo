// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/like"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/tweet"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/user"
)

// LikeCreate is the builder for creating a Like entity.
type LikeCreate struct {
	config
	mutation *LikeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (lc *LikeCreate) SetUserID(i int) *LikeCreate {
	lc.mutation.SetUserID(i)
	return lc
}

// SetTweetID sets the "tweet_id" field.
func (lc *LikeCreate) SetTweetID(i int) *LikeCreate {
	lc.mutation.SetTweetID(i)
	return lc
}

// SetLikedAt sets the "liked_at" field.
func (lc *LikeCreate) SetLikedAt(t time.Time) *LikeCreate {
	lc.mutation.SetLikedAt(t)
	return lc
}

// SetNillableLikedAt sets the "liked_at" field if the given value is not nil.
func (lc *LikeCreate) SetNillableLikedAt(t *time.Time) *LikeCreate {
	if t != nil {
		lc.SetLikedAt(*t)
	}
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LikeCreate) SetUser(u *User) *LikeCreate {
	return lc.SetUserID(u.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (lc *LikeCreate) SetTweet(t *Tweet) *LikeCreate {
	return lc.SetTweetID(t.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lc *LikeCreate) Mutation() *LikeMutation {
	return lc.mutation
}

// Save creates the Like in the database.
func (lc *LikeCreate) Save(ctx context.Context) (*Like, error) {
	var (
		err  error
		node *Like
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Like)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LikeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikeCreate) SaveX(ctx context.Context) *Like {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikeCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikeCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LikeCreate) defaults() {
	if _, ok := lc.mutation.LikedAt(); !ok {
		v := like.DefaultLikedAt()
		lc.mutation.SetLikedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikeCreate) check() error {
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Like.user_id"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "Like.tweet_id"`)}
	}
	if _, ok := lc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "liked_at", err: errors.New(`ent: missing required field "Like.liked_at"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Like.user"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "Like.tweet"`)}
	}
	return nil
}

func (lc *LikeCreate) sqlSave(ctx context.Context) (*Like, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (lc *LikeCreate) createSpec() (*Like, *sqlgraph.CreateSpec) {
	var (
		_node = &Like{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: like.Table,
		}
	)
	if value, ok := lc.mutation.LikedAt(); ok {
		_spec.SetField(like.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LikeCreateBulk is the builder for creating many Like entities in bulk.
type LikeCreateBulk struct {
	config
	builders []*LikeCreate
}

// Save creates the Like entities in the database.
func (lcb *LikeCreateBulk) Save(ctx context.Context) ([]*Like, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Like, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikeCreateBulk) SaveX(ctx context.Context) []*Like {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikeCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikeCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
