package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) MigrateTables(ctx context.Context, msgs []*message.WriteMigrateTable) error {
	c.logger.Info().Msg("Migrate")

	have, err := c.indexes()
	if err != nil {
		return err
	}
	tables := make(schema.Tables, len(msgs))
	for i, msg := range msgs {
		tables[i] = msg.Table
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
			if !msgs[need.Index].MigrateForce {
				return fmt.Errorf("index %s requires force migration. use 'migrate_mode: forced'", uid)
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
