// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Bpazy/behappy/ent/migrate"

	"github.com/Bpazy/behappy/ent/hero"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Hero is the client for interacting with the Hero builders.
	Hero *HeroClient
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
	c.Hero = NewHeroClient(c.config)
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
		Hero:   NewHeroClient(cfg),
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
		Hero:   NewHeroClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Hero.
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
	c.Hero.Use(hooks...)
}

// HeroClient is a client for the Hero schema.
type HeroClient struct {
	config
}

// NewHeroClient returns a client for the Hero from the given config.
func NewHeroClient(c config) *HeroClient {
	return &HeroClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hero.Hooks(f(g(h())))`.
func (c *HeroClient) Use(hooks ...Hook) {
	c.hooks.Hero = append(c.hooks.Hero, hooks...)
}

// Create returns a builder for creating a Hero entity.
func (c *HeroClient) Create() *HeroCreate {
	mutation := newHeroMutation(c.config, OpCreate)
	return &HeroCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Hero entities.
func (c *HeroClient) CreateBulk(builders ...*HeroCreate) *HeroCreateBulk {
	return &HeroCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Hero.
func (c *HeroClient) Update() *HeroUpdate {
	mutation := newHeroMutation(c.config, OpUpdate)
	return &HeroUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HeroClient) UpdateOne(h *Hero) *HeroUpdateOne {
	mutation := newHeroMutation(c.config, OpUpdateOne, withHero(h))
	return &HeroUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HeroClient) UpdateOneID(id int) *HeroUpdateOne {
	mutation := newHeroMutation(c.config, OpUpdateOne, withHeroID(id))
	return &HeroUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Hero.
func (c *HeroClient) Delete() *HeroDelete {
	mutation := newHeroMutation(c.config, OpDelete)
	return &HeroDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HeroClient) DeleteOne(h *Hero) *HeroDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HeroClient) DeleteOneID(id int) *HeroDeleteOne {
	builder := c.Delete().Where(hero.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HeroDeleteOne{builder}
}

// Query returns a query builder for Hero.
func (c *HeroClient) Query() *HeroQuery {
	return &HeroQuery{
		config: c.config,
	}
}

// Get returns a Hero entity by its id.
func (c *HeroClient) Get(ctx context.Context, id int) (*Hero, error) {
	return c.Query().Where(hero.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HeroClient) GetX(ctx context.Context, id int) *Hero {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *HeroClient) Hooks() []Hook {
	return c.hooks.Hero
}
