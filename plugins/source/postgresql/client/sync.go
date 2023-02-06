package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/jackc/pgx/v5/pgtype"
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
	if c.pluginSpec.CDC {
		// if we use cdc we need to set the snapshot that was exported when we started the logical
		// replication stream
		if _, err := tx.Exec(ctx, "SET TRANSACTION SNAPSHOT SNAPSHOT '" + c.createReplicationSlotResult.SnapshotName + "'"); err != nil {
			return err
		}
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
			resource, err := c.resourceFromValues(table.Name, values)
			if err != nil {
				return err
			}
			res <- resource
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	if !c.pluginSpec.CDC {
		return nil
	}
	// clientXLogPos := sysident.XLogPos
	// standbyMessageTimeout := time.Second * 10
	// nextStandbyMessageDeadline := time.Now().Add(standbyMessageTimeout)
	relations := map[uint32]*pglogrepl.RelationMessage{}
	typeMap := pgtype.NewMap()

	conn := tx.Conn().PgConn()

	for {
		rawMsg, err := conn.ReceiveMessage(ctx)
		if errMsg, ok := rawMsg.(*pgproto3.ErrorResponse); ok {
			return fmt.Errorf("received Postgres WAL error: %+v", errMsg)
		}
		if err != nil {
			return err
		}
	
		msg, ok := rawMsg.(*pgproto3.CopyData)
		if !ok {
			return fmt.Errorf("received unexpected message: %T", rawMsg)
		}
	
		switch msg.Data[0] {
		case pglogrepl.PrimaryKeepaliveMessageByteID:
			pkm, err := pglogrepl.ParsePrimaryKeepaliveMessage(msg.Data[1:])
			if err != nil {
				return fmt.Errorf("ParsePrimaryKeepaliveMessage failed: %w", err)
			}
			c.logger.Info().Msgf("Primary Keepalive Message =>", "ServerWALEnd:", pkm.ServerWALEnd, "ServerTime:", pkm.ServerTime, "ReplyRequested:", pkm.ReplyRequested)
	
			// if pkm.ReplyRequested {
			// 	nextStandbyMessageDeadline = time.Time{}
			// }
	
		case pglogrepl.XLogDataByteID:
			xld, err := pglogrepl.ParseXLogData(msg.Data[1:])
			if err != nil {
				return fmt.Errorf("ParseXLogData failed: %w", err)
			}
			c.logger.Info().Msgf("XLogData => WALStart %s ServerWALEnd %s ServerTime %s WALData:\n%s\n", xld.WALStart, xld.ServerWALEnd, xld.ServerTime, hex.Dump(xld.WALData))
			logicalMsg, err := pglogrepl.Parse(xld.WALData)
			if err != nil {
				return fmt.Errorf("parse logical replication message: %w", err)
			}
			c.logger.Info().Msgf("Receive a logical replication message: %s", logicalMsg.Type())
			switch logicalMsg := logicalMsg.(type) {
			case *pglogrepl.RelationMessage:
				relations[logicalMsg.RelationID] = logicalMsg
	
			case *pglogrepl.BeginMessage:
				// Indicates the beginning of a group of changes in a transaction. This is only sent for committed transactions. You won't get any events from rolled back transactions.
	
			case *pglogrepl.CommitMessage:
	
			case *pglogrepl.InsertMessage:
				rel, ok := relations[logicalMsg.RelationID]
				if !ok {
					return fmt.Errorf("unknown relation ID %d", logicalMsg.RelationID)
				}
				values := map[string]interface{}{}
				for idx, col := range logicalMsg.Tuple.Columns {
					colName := rel.Columns[idx].Name
					switch col.DataType {
					case 'n': // null
						values[colName] = nil
					case 'u': // unchanged toast
						// This TOAST value was not changed. TOAST values are not stored in the tuple, and logical replication doesn't want to spend a disk read to fetch its value for you.
					case 't': //text
						val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
						if err != nil {
							return fmt.Errorf("error decoding column data: %w", err)
						}
						values[colName] = val
					}
				}
				log.Printf("INSERT INTO %s.%s: %v", rel.Namespace, rel.RelationName, values)
				resource, err := c.resourceFromCDCValues(rel.RelationName, values)
				if err != nil {
					return err
				}
				res <- resource
	
			case *pglogrepl.UpdateMessage:
				rel, ok := relations[logicalMsg.RelationID]
				if !ok {
					log.Fatalf("unknown relation ID %d", logicalMsg.RelationID)
				}
				values := map[string]interface{}{}
				for idx, col := range logicalMsg.NewTuple.Columns {
					colName := rel.Columns[idx].Name
					switch col.DataType {
					case 'n': // null
						values[colName] = nil
					case 'u': // unchanged toast
						// This TOAST value was not changed. TOAST values are not stored in the tuple, and logical replication doesn't want to spend a disk read to fetch its value for you.
					case 't': //text
						val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
						if err != nil {
							log.Fatalln("error decoding column data:", err)
						}
						values[colName] = val
					}
				}
				log.Printf("UPDATE INTO %s.%s: %v", rel.Namespace, rel.RelationName, values)
				resource, err := c.resourceFromCDCValues(rel.RelationName, values)
				if err != nil {
					return err
				}
				res <- resource
			case *pglogrepl.DeleteMessage:
				// ...
			case *pglogrepl.TruncateMessage:
				// ...
	
			case *pglogrepl.TypeMessage:
			case *pglogrepl.OriginMessage:
			default:
				log.Printf("Unknown message type in pgoutput stream: %T", logicalMsg)
			}
		}		
	}
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