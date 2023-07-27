// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/stark-sim/serializer/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/serializer/ent/collect"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Collect is the client for interacting with the Collect builders.
	Collect *CollectClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Collect = NewCollectClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		ctx:     ctx,
		config:  cfg,
		Collect: NewCollectClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Collect: NewCollectClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Collect.
//		Query().
//		Count(ctx)
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
	c.Collect.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Collect.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CollectMutation:
		return c.Collect.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CollectClient is a client for the Collect schema.
type CollectClient struct {
	config
}

// NewCollectClient returns a client for the Collect from the given config.
func NewCollectClient(c config) *CollectClient {
	return &CollectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `collect.Hooks(f(g(h())))`.
func (c *CollectClient) Use(hooks ...Hook) {
	c.hooks.Collect = append(c.hooks.Collect, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `collect.Intercept(f(g(h())))`.
func (c *CollectClient) Intercept(interceptors ...Interceptor) {
	c.inters.Collect = append(c.inters.Collect, interceptors...)
}

// Create returns a builder for creating a Collect entity.
func (c *CollectClient) Create() *CollectCreate {
	mutation := newCollectMutation(c.config, OpCreate)
	return &CollectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Collect entities.
func (c *CollectClient) CreateBulk(builders ...*CollectCreate) *CollectCreateBulk {
	return &CollectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Collect.
func (c *CollectClient) Update() *CollectUpdate {
	mutation := newCollectMutation(c.config, OpUpdate)
	return &CollectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CollectClient) UpdateOne(co *Collect) *CollectUpdateOne {
	mutation := newCollectMutation(c.config, OpUpdateOne, withCollect(co))
	return &CollectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CollectClient) UpdateOneID(id int) *CollectUpdateOne {
	mutation := newCollectMutation(c.config, OpUpdateOne, withCollectID(id))
	return &CollectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Collect.
func (c *CollectClient) Delete() *CollectDelete {
	mutation := newCollectMutation(c.config, OpDelete)
	return &CollectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CollectClient) DeleteOne(co *Collect) *CollectDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CollectClient) DeleteOneID(id int) *CollectDeleteOne {
	builder := c.Delete().Where(collect.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CollectDeleteOne{builder}
}

// Query returns a query builder for Collect.
func (c *CollectClient) Query() *CollectQuery {
	return &CollectQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCollect},
		inters: c.Interceptors(),
	}
}

// Get returns a Collect entity by its id.
func (c *CollectClient) Get(ctx context.Context, id int) (*Collect, error) {
	return c.Query().Where(collect.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CollectClient) GetX(ctx context.Context, id int) *Collect {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CollectClient) Hooks() []Hook {
	return c.hooks.Collect
}

// Interceptors returns the client interceptors.
func (c *CollectClient) Interceptors() []Interceptor {
	return c.inters.Collect
}

func (c *CollectClient) mutate(ctx context.Context, m *CollectMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CollectCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CollectUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CollectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CollectDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Collect mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Collect []ent.Hook
	}
	inters struct {
		Collect []ent.Interceptor
	}
)