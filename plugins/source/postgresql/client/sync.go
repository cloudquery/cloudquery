package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c *Client) Sync(ctx context.Context, metrics *source.Metrics, res chan<- *schema.Resource) error {
	// var conn *pgconn.PgConn
	c.metrics = metrics
	for _, table := range c.Tables {
		if c.metrics.TableClient[table.Name] == nil {
			c.metrics.TableClient[table.Name] = make(map[string]*source.TableClientMetrics)
			c.metrics.TableClient[table.Name][c.ID()] = &source.TableClientMetrics{}
		}
	}

	if c.pluginSpec.CDC {
		if err := c.startCDC(ctx); err != nil {
			return err
		}
	}

	// tx, err := c.Conn.BeginTx(ctx, pgx.TxOptions{
	// 	// this transaction is needed for us to take a snapshot and we need to close it only at the end of the initial sync
	// 	IsoLevel: pgx.RepeatableRead,
	// 	AccessMode: pgx.ReadOnly,
	// })
	// if err != nil {
	// 	return err
	// }
	// defer tx.Commit(ctx)
	// if err != nil {
	// 	return err
	// }
	// if c.pluginSpec.CDC {
	// 	// if we use cdc we need to set the snapshot that was exported when we started the logical
	// 	// replication stream
	// 	if _, err := tx.Exec(ctx, "SET TRANSACTION SNAPSHOT '" + c.createReplicationSlotResult.SnapshotName + "'"); err != nil {
	// 		return fmt.Errorf("failed to 'SET TRANSACTION SNAPSHOT %s': %w", c.createReplicationSlotResult.SnapshotName, err)
	// 	}
	// }
	// for _, table := range c.Tables {
	// 	colNames := make([]string, len(table.Columns))
	// 	for i, col := range table.Columns {
	// 		colNames[i] = pgx.Identifier{col.Name}.Sanitize()
	// 	}
	// 	query := "SELECT " + strings.Join(colNames, ",") + " FROM " + pgx.Identifier{table.Name}.Sanitize()
	// 	rows, err := c.Conn.Query(ctx, query)
	// 	if err != nil {
	// 		c.metrics.TableClient[table.Name][c.ID()].Errors++
	// 		return err
	// 	}
	// 	defer rows.Close()
	// 	for rows.Next() {
	// 		values, err := rows.Values()
	// 		if err != nil {
	// 			c.metrics.TableClient[table.Name][c.ID()].Errors++
	// 			return err
	// 		}
	// 		resource, err := c.resourceFromValues(table.Name, values)
	// 		if err != nil {
	// 			c.metrics.TableClient[table.Name][c.ID()].Errors++
	// 			return err
	// 		}
	// 		c.metrics.TableClient[table.Name][c.ID()].Resources++
	// 		res <- resource
	// 	}
	// }
	// if err := tx.Commit(ctx); err != nil {
	// 	return err
	// }
	if !c.pluginSpec.CDC {
		return nil
	}
	
	if err := c.listenCDC(ctx, res); err != nil {
		return err
	}
	return nil
}

func (c *Client) resourceFromValues(tableName string, values []interface{}) (*schema.Resource, error) {
	table := c.Tables.Get(tableName)
	resource := schema.NewResourceData(table, nil, values)
	for i, col := range table.Columns {
		if err := resource.Set(col.Name, values[i]); err != nil {
			return nil, err
		}
	}
	return resource, nil
}

func (c *Client) resourceFromCDCValues(tableName string, values map[string]interface{}) (*schema.Resource, error) {
	table := c.Tables.Get(tableName)
	resource := schema.NewResourceData(table, nil, values)
	for _, col := range table.Columns {
		if err := resource.Set(col.Name, values[col.Name]); err != nil {
			return nil, err
		}
	}
	return resource, nil
}


func decodeTextColumnData(mi *pgtype.Map, data []byte, dataType uint32) (interface{}, error) {
	if dt, ok := mi.TypeForOID(dataType); ok {
		return dt.Codec.DecodeValue(mi, dataType, pgtype.TextFormatCode, data)
	}
	return string(data), nil
}