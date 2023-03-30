package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
)

func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	c.logger.Info().Msg("Migrate")

	have, err := c.indexes()
	if err != nil {
		return err
	}

	want := c.tablesIndexSchemas(tables)

	var recreate, create, update []*indexSchema
	for uid, need := range want {
		got := have[uid]
		switch {
		case got == nil:
			create = append(create, need)
		case got.canMigrate(need):
			update = append(update, need)
		default:
			recreate = append(recreate, need)
		}
	}

	if len(recreate) > 0 && c.dstSpec.MigrateMode != specs.MigrateModeForced {
		names := make([]string, len(recreate))
		for i, index := range recreate {
			names[i] = index.UID
		}
		return fmt.Errorf("indexes %v require force migration. use 'migrate_mode: forced'", names)
	}

	for _, index := range create {
		if err := c.createIndex(ctx, index); err != nil {
			return err
		}
	}

	for _, index := range update {
		if err := c.configureIndex(ctx, index); err != nil {
			return err
		}
	}

	for _, index := range recreate {
		if err := c.recreateIndex(ctx, index); err != nil {
			return err
		}
	}

	return nil
}
