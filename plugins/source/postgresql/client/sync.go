package client

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5"
)

func (c *Client) Sync(ctx context.Context, res chan<- *schema.Resource) error {
	if c.pluginSpec.CDC {
		connPool, err := c.Conn.Acquire(ctx)
		if err != nil {
			return fmt.Errorf("failed to acquire connection: %w", err)
		}
		defer connPool.Release()
		conn := connPool.Conn().PgConn()
	
		tables := strings.Join(c.Tables.TableNames(), ",")
		reader := conn.Exec(ctx, fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s;", pgx.Identifier{c.spec.Name}.Sanitize(), tables))
		_, err = reader.ReadAll()
		if err != nil {
			return fmt.Errorf("failed to create publication: %w", err)
		}
	
		sysident, err := pglogrepl.IdentifySystem(ctx, conn)
		if err != nil {
			return fmt.Errorf("failed to identify system: %w", err)
		}
	
		sql := fmt.Sprintf("CREATE_REPLICATION_SLOT %s LOGICAL pgoutput EXPORT_SNAPSHOT", c.spec.Name)
		c.createReplicationSlotResult, err = pglogrepl.ParseCreateReplicationSlot(conn.Exec(ctx, sql))
		if err != nil {
			return fmt.Errorf("failed to create replication slot: %w", err)
		}
	
		if err := pglogrepl.StartReplication(ctx, conn, c.createReplicationSlotResult.SlotName, sysident.XLogPos,
			pglogrepl.StartReplicationOptions{
				PluginArgs: []string{"proto_version '1'", "publication_names '" + c.spec.Name + "'"},
			}); err != nil {
			return fmt.Errorf("failed to start replication: %w", err)
		}
	}
	

	tx, err := c.Conn.BeginTx(ctx, pgx.TxOptions{
		// this transaction is needed for us to take a snapshot and we need to close it only at the end of the initial sync
		IsoLevel: pgx.RepeatableRead,
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		return err
	}
	defer tx.Commit(ctx)
	if err != nil {
		return err
	}

	for _, table := range c.Tables {
		colNames := make([]string, len(table.Columns))
		for i, col := range table.Columns {
			colNames[i] = pgx.Identifier{col.Name}.Sanitize()
		}
		query := "SELECT " + strings.Join(colNames, ",") + " FROM " + pgx.Identifier{table.Name}.Sanitize()
		rows, err := c.Conn.Query(ctx, query)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}
			resource, err := resourceFromValues(table, values)
			if err != nil {
				return err
			}
			res <- resource
		}
	}

	conn := tx.Conn().PgConn()


	log.Println("create publication pglogrepl_demo")
	return nil
}

func resourceFromValues(table *schema.Table, values []interface{}) (*schema.Resource, error) {
	resource := schema.NewResourceData(table, nil, values)
	for i, col := range table.Columns {
		if err := resource.Set(col.Name, values[i]); err != nil {
			return nil, err
		}
	}
	return resource, nil
}