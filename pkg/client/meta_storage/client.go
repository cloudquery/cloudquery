// Package meta_storage interacts with core database schema and stores cloudquery metadata such as fetch summaries
package meta_storage

import (
	"context"
	"embed"
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cq-provider-sdk/database/dsn"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
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
	db     execution.QueryExecer
	Logger hclog.Logger
}

func NewClient(db execution.QueryExecer, logger hclog.Logger) *Client {
	return &Client{
		db:     db,
		Logger: logger,
	}
}

func (c *Client) MigrateCore(ctx context.Context, de database.DialectExecutor) error {
	err := c.db.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS cloudquery")
	if err != nil {
		return err
	}

	newDSN, err := de.Setup(ctx)
	if err != nil {
		return err
	}

	migrations, err := migrator.ReadMigrationFiles(c.Logger, coreMigrations)
	if err != nil {
		return err
	}
	newDSN, err = dsn.SetDSNElement(newDSN, map[string]string{"search_path": "cloudquery"})
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
