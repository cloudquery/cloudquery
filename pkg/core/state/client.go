// Package meta_storage interacts with core database schema and stores cloudquery metadata such as fetch summaries
package state

import (
	"context"
	"embed"
	"fmt"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/database/dsn"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/execution"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-hclog"
)

var (
	//go:embed migrations/*/*.sql
	coreMigrations embed.FS
)

type Client struct {
	dsn    string
	db     execution.QueryExecer
	closer func()
	Logger hclog.Logger
}

func NewClient(db execution.QueryExecer, logger hclog.Logger) *Client {
	return &Client{
		db:     db,
		Logger: logger,
	}
}

func NewMigratedClient(ctx context.Context, dsn string, logger hclog.Logger) (*Client, error) {
	c := &Client{
		dsn:    dsn,
		Logger: logger,
	}

	db, err := sdkdb.New(ctx, c.Logger, c.dsn)
	if err != nil {
		return nil, diag.FromError(err, diag.DATABASE)
	}
	c.db = db
	c.closer = db.Close

	// migrate CloudQuery core tables to latest version
	if err := c.migrateCore(ctx); err != nil {
		return nil, diag.FromError(err, diag.DATABASE, diag.WithSummary("failed to migrate cloudquery_core tables"))
	}

	return c, nil
}

func (c *Client) Close() {
	if c.db == nil {
		return
	}
	c.closer()
	c.db, c.closer = nil, nil
}

func (c *Client) migrateCore(ctx context.Context) error {
	if err := c.db.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS cloudquery"); err != nil {
		return err
	}

	migrations, err := migrator.ReadMigrationFiles(c.Logger, coreMigrations)
	if err != nil {
		return err
	}
	newDSN, err := dsn.SetDSNElement(c.dsn, map[string]string{"search_path": "cloudquery"})
	if err != nil {
		return err
	}
	m, err := migrator.New(c.Logger, schema.Postgres, migrations, newDSN, "cloudquery_core")
	if err != nil {
		return err
	}

	defer func() {
		if err := m.Close(); err != nil {
			c.Logger.Error("failed to close migrator connection", "error", err)
		}
	}()

	if err := m.UpgradeProvider(migrator.Latest); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate cloudquery core schema: %w", err)
	}
	return nil
}
