package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

// MigrateTables tries to create indexes for the tables.
// It will issue `CREATE RANGE INDEX ... IF NOT EXIST`.
// If the error occurs & indicates that the issue is caused by conflicting schema, 2 scenarios can happen:
// 1. Force mode is selected for migration: drop & recreate index (without checking for error this time)
// 2. No forced migration is requested - return error.
func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	if len(messages) == 0 {
		return nil
	}

	createQueries := make([]string, len(messages))
	for i, m := range messages {
		createQuery, err := createIndexQuery(m.Table)
		if err != nil {
			return err
		}
		createQueries[i] = createQuery
	}

	sess := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer sess.Close(ctx)

	for i, m := range messages {
		if err := c.tryCreateIndex(ctx, sess, m, createQueries[i]); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) tryCreateIndex(ctx context.Context, sess neo4j.Session, migrate *message.WriteMigrateTable, createQuery string) error {
	if len(createQuery) == 0 {
		c.logger.Debug().Str("table", migrate.Table.Name).Msg("table has no primary keys, skipping")
		// nothing to be done here, table has no primary keys
		return nil
	}

	index := indexName(migrate.Table)

	_, err := sess.Run(ctx, createQuery, map[string]any{})
	if err == nil {
		return nil
	}

	if !migrate.MigrateForce {
		c.logger.Err(err).
			Str("table", migrate.Table.Name).
			Str("index", index).
			Msg("failed to create index")
		return fmt.Errorf("failed to create index for %q: %w", migrate.Table.Name, err)
	}

	c.logger.Warn().
		Str("table", migrate.Table.Name).
		Str("index", index).
		Err(err).
		Msg("failed to create index, recreating")

	dropQuery, err := dropIndexQuery(migrate.Table)
	if err != nil {
		return err
	}
	_, err = sess.Run(ctx, dropQuery, map[string]any{})
	if err != nil {
		c.logger.Err(err).
			Str("table", migrate.Table.Name).
			Str("index", index).
			Msg("failed to drop index")
		return fmt.Errorf("failed to drop index for %q: %w", migrate.Table.Name, err)
	}

	_, err = sess.Run(ctx, createQuery, map[string]any{})
	return err
}

func createIndexQuery(table *schema.Table) (string, error) {
	label, err := quoteIdentifier(table.Name)
	if err != nil {
		return "", err
	}
	index, err := quoteIdentifier(indexName(table))
	if err != nil {
		return "", err
	}
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// empty query = no need to do anything
		return "", nil
	}

	var sb strings.Builder
	sb.WriteString(`CREATE INDEX `)
	sb.WriteString(index)
	sb.WriteString(` IF NOT EXISTS FOR (n:`)
	sb.WriteString(label)
	sb.WriteString(`) ON (`)
	for i, name := range pks {
		key, err := quoteIdentifier(name)
		if err != nil {
			return "", err
		}
		if i != 0 {
			sb.WriteString(`, `)
		}
		sb.WriteString(`n.`)
		sb.WriteString(key)
	}
	sb.WriteString(`);`)

	return sb.String(), nil
}

func dropIndexQuery(table *schema.Table) (string, error) {
	index, err := quoteIdentifier(indexName(table))
	if err != nil {
		return "", err
	}
	return `DROP INDEX ` + index + ` IF EXISTS;`, nil
}

func indexName(table *schema.Table) string {
	return `_cq_index_` + table.Name
}
