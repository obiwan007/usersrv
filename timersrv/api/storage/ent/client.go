// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/migrate"

	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Timer is the client for interacting with the Timer builders.
	Timer *TimerClient
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
	c.Timer = NewTimerClient(c.config)
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Timer:  NewTimerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Timer:  NewTimerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Timer.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Timer.Use(hooks...)
}

// TimerClient is a client for the Timer schema.
type TimerClient struct {
	config
}

// NewTimerClient returns a client for the Timer from the given config.
func NewTimerClient(c config) *TimerClient {
	return &TimerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `timer.Hooks(f(g(h())))`.
func (c *TimerClient) Use(hooks ...Hook) {
	c.hooks.Timer = append(c.hooks.Timer, hooks...)
}

// Create returns a create builder for Timer.
func (c *TimerClient) Create() *TimerCreate {
	mutation := newTimerMutation(c.config, OpCreate)
	return &TimerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Timer.
func (c *TimerClient) Update() *TimerUpdate {
	mutation := newTimerMutation(c.config, OpUpdate)
	return &TimerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TimerClient) UpdateOne(t *Timer) *TimerUpdateOne {
	return c.UpdateOneID(t.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *TimerClient) UpdateOneID(id int) *TimerUpdateOne {
	mutation := newTimerMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &TimerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Timer.
func (c *TimerClient) Delete() *TimerDelete {
	mutation := newTimerMutation(c.config, OpDelete)
	return &TimerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TimerClient) DeleteOne(t *Timer) *TimerDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TimerClient) DeleteOneID(id int) *TimerDeleteOne {
	builder := c.Delete().Where(timer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TimerDeleteOne{builder}
}

// Create returns a query builder for Timer.
func (c *TimerClient) Query() *TimerQuery {
	return &TimerQuery{config: c.config}
}

// Get returns a Timer entity by its id.
func (c *TimerClient) Get(ctx context.Context, id int) (*Timer, error) {
	return c.Query().Where(timer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TimerClient) GetX(ctx context.Context, id int) *Timer {
	t, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return t
}

// Hooks returns the client hooks.
func (c *TimerClient) Hooks() []Hook {
	return c.hooks.Timer
}
