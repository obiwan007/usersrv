// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent/migrate"

	"github.com/obiwan007/usersrv/clientsrv/api/storage/ent/client"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Client is the client for interacting with the Client builders.
	Client *ClientClient
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
	c.Client = NewClientClient(c.config)
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
		Client: NewClientClient(cfg),
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
		Client: NewClientClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Client.
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
	c.Client.Use(hooks...)
}

// ClientClient is a client for the Client schema.
type ClientClient struct {
	config
}

// NewClientClient returns a client for the Client from the given config.
func NewClientClient(c config) *ClientClient {
	return &ClientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `client.Hooks(f(g(h())))`.
func (c *ClientClient) Use(hooks ...Hook) {
	c.hooks.Client = append(c.hooks.Client, hooks...)
}

// Create returns a create builder for Client.
func (c *ClientClient) Create() *ClientCreate {
	mutation := newClientMutation(c.config, OpCreate)
	return &ClientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Client.
func (c *ClientClient) Update() *ClientUpdate {
	mutation := newClientMutation(c.config, OpUpdate)
	return &ClientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClientClient) UpdateOne(cl *Client) *ClientUpdateOne {
	return c.UpdateOneID(cl.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ClientClient) UpdateOneID(id int) *ClientUpdateOne {
	mutation := newClientMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &ClientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Client.
func (c *ClientClient) Delete() *ClientDelete {
	mutation := newClientMutation(c.config, OpDelete)
	return &ClientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ClientClient) DeleteOne(cl *Client) *ClientDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ClientClient) DeleteOneID(id int) *ClientDeleteOne {
	builder := c.Delete().Where(client.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClientDeleteOne{builder}
}

// Create returns a query builder for Client.
func (c *ClientClient) Query() *ClientQuery {
	return &ClientQuery{config: c.config}
}

// Get returns a Client entity by its id.
func (c *ClientClient) Get(ctx context.Context, id int) (*Client, error) {
	return c.Query().Where(client.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClientClient) GetX(ctx context.Context, id int) *Client {
	cl, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return cl
}

// Hooks returns the client hooks.
func (c *ClientClient) Hooks() []Hook {
	return c.hooks.Client
}
