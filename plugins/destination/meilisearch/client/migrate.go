package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	c.logger.Info().Msg("Migrate")

	have, err := c.indexes()
	if err != nil {
		return err
	}
	tables := make(schema.Tables, len(messages))
	for i, msg := range messages {
		tables[i] = msg.Table
	}

	want := tablesIndexSchemas(tables)

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
			if !messages[need.Index].MigrateForce {
				return fmt.Errorf("index %s requires force migration. Migrate manually or consider using 'migrate_mode: forced'", uid)
			}
		}
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
