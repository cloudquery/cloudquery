package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// MigrateTables tries to create constraints & indexes for the tables.
// It will issue `CREATE CONSTRAINT ... IF NOT EXIST`.
// If the error occurs & indicates that the issue is caused by conflicting schema, 2 scenarios can happen:
// 1. Force mode is selected for migration: drop & recreate constraint (without checking for error this time)
// 2. No forced migration is requested - return error.
func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	if len(messages) == 0 {
		return nil
	}

	sess := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer sess.Close(ctx)

	for _, m := range messages {
		if err := c.tryCreateConstraint(ctx, sess, m); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) tryCreateConstraint(ctx context.Context, sess neo4j.SessionWithContext, migrate *message.WriteMigrateTable) error {
	createQuery := createConstraintQuery(migrate.Table)
	if len(createQuery) == 0 {
		c.logger.Debug().Str("table", migrate.Table.Name).Msg("table has no primary keys, skipping")
		// nothing to be done here, table has no primary keys
		return nil
	}

	constraint := constraintName(migrate.Table)

	_, err := sess.Run(ctx, createQuery, map[string]any{})
	if err == nil {
		return nil
	}

	if !migrate.MigrateForce {
		c.logger.Err(err).
			Str("table", migrate.Table.Name).
			Str("constraint", constraint).
			Msg("failed to create constraint")
		return fmt.Errorf("failed to create constraint for %q: %w", migrate.Table.Name, err)
	}

	c.logger.Warn().
		Str("table", migrate.Table.Name).
		Str("constraint", constraint).
		Err(err).
		Msg("failed to create constraint, recreating")

	_, err = sess.Run(ctx, `DROP CONSTRAINT `+constraint+` IF EXISTS;`, map[string]any{})
	if err != nil {
		c.logger.Err(err).
			Str("table", migrate.Table.Name).
			Str("constraint", constraint).
			Msg("failed to drop constraint")
		return fmt.Errorf("failed to drop constraint for %q: %w", migrate.Table.Name, err)
	}

	_, err = sess.Run(ctx, createQuery, map[string]any{})
	return err
}

func createConstraintQuery(table *schema.Table) string {
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// empty query = no need to do anything
		return ""
	}

	var sb strings.Builder
	sb.WriteString(`CREATE CONSTRAINT `)
	sb.WriteString(constraintName(table))
	sb.WriteString(` IF NOT EXISTS FOR (n:`)
	sb.WriteString(table.Name)
	sb.WriteString(`) REQUIRE (`)
	for i, name := range pks {
		if i != 0 {
			sb.WriteString(`, `)
		}
		sb.WriteString(`n.`)
		sb.WriteString(name)
	}
	sb.WriteString(`) IS UNIQUE;`)

	return sb.String()
}

func constraintName(table *schema.Table) string {
	return `_cq_constraint_` + table.Name
}
