// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/migrate"

	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/like"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/tweet"
	"github.com/xlzhangkeke/entdemo/examples/tweet/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Like is the client for interacting with the Like builders.
	Like *LikeClient
	// Tweet is the client for interacting with the Tweet builders.
	Tweet *TweetClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Like = NewLikeClient(c.config)
	c.Tweet = NewTweetClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Like:   NewLikeClient(cfg),
		Tweet:  NewTweetClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Like:   NewLikeClient(cfg),
		Tweet:  NewTweetClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Like.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Like.Use(hooks...)
	c.Tweet.Use(hooks...)
	c.User.Use(hooks...)
}

// LikeClient is a client for the Like schema.
type LikeClient struct {
	config
}

// NewLikeClient returns a client for the Like from the given config.
func NewLikeClient(c config) *LikeClient {
	return &LikeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `like.Hooks(f(g(h())))`.
func (c *LikeClient) Use(hooks ...Hook) {
	c.hooks.Like = append(c.hooks.Like, hooks...)
}

// Create returns a builder for creating a Like entity.
func (c *LikeClient) Create() *LikeCreate {
	mutation := newLikeMutation(c.config, OpCreate)
	return &LikeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Like entities.
func (c *LikeClient) CreateBulk(builders ...*LikeCreate) *LikeCreateBulk {
	return &LikeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Like.
func (c *LikeClient) Update() *LikeUpdate {
	mutation := newLikeMutation(c.config, OpUpdate)
	return &LikeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LikeClient) UpdateOne(l *Like) *LikeUpdateOne {
	mutation := newLikeMutation(c.config, OpUpdateOne)
	mutation.user = &l.UserID
	mutation.tweet = &l.TweetID
	return &LikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Like.
func (c *LikeClient) Delete() *LikeDelete {
	mutation := newLikeMutation(c.config, OpDelete)
	return &LikeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Query returns a query builder for Like.
func (c *LikeClient) Query() *LikeQuery {
	return &LikeQuery{
		config: c.config,
	}
}

// QueryUser queries the user edge of a Like.
func (c *LikeClient) QueryUser(l *Like) *UserQuery {
	return c.Query().
		Where(like.UserID(l.UserID), like.TweetID(l.TweetID)).
		QueryUser()
}

// QueryTweet queries the tweet edge of a Like.
func (c *LikeClient) QueryTweet(l *Like) *TweetQuery {
	return c.Query().
		Where(like.UserID(l.UserID), like.TweetID(l.TweetID)).
		QueryTweet()
}

// Hooks returns the client hooks.
func (c *LikeClient) Hooks() []Hook {
	return c.hooks.Like
}

// TweetClient is a client for the Tweet schema.
type TweetClient struct {
	config
}

// NewTweetClient returns a client for the Tweet from the given config.
func NewTweetClient(c config) *TweetClient {
	return &TweetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tweet.Hooks(f(g(h())))`.
func (c *TweetClient) Use(hooks ...Hook) {
	c.hooks.Tweet = append(c.hooks.Tweet, hooks...)
}

// Create returns a builder for creating a Tweet entity.
func (c *TweetClient) Create() *TweetCreate {
	mutation := newTweetMutation(c.config, OpCreate)
	return &TweetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tweet entities.
func (c *TweetClient) CreateBulk(builders ...*TweetCreate) *TweetCreateBulk {
	return &TweetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tweet.
func (c *TweetClient) Update() *TweetUpdate {
	mutation := newTweetMutation(c.config, OpUpdate)
	return &TweetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TweetClient) UpdateOne(t *Tweet) *TweetUpdateOne {
	mutation := newTweetMutation(c.config, OpUpdateOne, withTweet(t))
	return &TweetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TweetClient) UpdateOneID(id int) *TweetUpdateOne {
	mutation := newTweetMutation(c.config, OpUpdateOne, withTweetID(id))
	return &TweetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tweet.
func (c *TweetClient) Delete() *TweetDelete {
	mutation := newTweetMutation(c.config, OpDelete)
	return &TweetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TweetClient) DeleteOne(t *Tweet) *TweetDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TweetClient) DeleteOneID(id int) *TweetDeleteOne {
	builder := c.Delete().Where(tweet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TweetDeleteOne{builder}
}

// Query returns a query builder for Tweet.
func (c *TweetClient) Query() *TweetQuery {
	return &TweetQuery{
		config: c.config,
	}
}

// Get returns a Tweet entity by its id.
func (c *TweetClient) Get(ctx context.Context, id int) (*Tweet, error) {
	return c.Query().Where(tweet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TweetClient) GetX(ctx context.Context, id int) *Tweet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLikedUsers queries the liked_users edge of a Tweet.
func (c *TweetClient) QueryLikedUsers(t *Tweet) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tweet.LikedUsersTable, tweet.LikedUsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikes queries the likes edge of a Tweet.
func (c *TweetClient) QueryLikes(t *Tweet) *LikeQuery {
	query := &LikeQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tweet.Table, tweet.FieldID, id),
			sqlgraph.To(like.Table, like.TweetColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, tweet.LikesTable, tweet.LikesColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TweetClient) Hooks() []Hook {
	return c.hooks.Tweet
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLikedTweets queries the liked_tweets edge of a User.
func (c *UserClient) QueryLikedTweets(u *User) *TweetQuery {
	query := &TweetQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.LikedTweetsTable, user.LikedTweetsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikes queries the likes edge of a User.
func (c *UserClient) QueryLikes(u *User) *LikeQuery {
	query := &LikeQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(like.Table, like.UserColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, user.LikesTable, user.LikesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}