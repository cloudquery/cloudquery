package client

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/jackc/pgx/v5/pgtype"
)

func (c *Client) createPublicationForTables(ctx context.Context, conn *pgconn.PgConn) error {
	sql := fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s", pgx.Identifier{c.spec.Name}.Sanitize(), strings.Join(c.Tables.TableNames(), ","))
	reader := conn.Exec(ctx, sql)
	if _, err := reader.ReadAll(); err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return fmt.Errorf("failed to create publication: %w", err)
		}
		if pgErr.Code != "42710" {
			// not recoverable error
			return fmt.Errorf("failed to create publication with pgerror %s: %w", pgErrToStr(pgErr), err)
		}
		sql = fmt.Sprintf("ALTER PUBLICATION %s SET TABLE %s", pgx.Identifier{c.spec.Name}.Sanitize(), strings.Join(c.Tables.TableNames(), ","))
		reader := conn.Exec(ctx, sql)
		if _, err := reader.ReadAll(); err != nil {
			return fmt.Errorf("failed to alter publication: %w", err)
		}
	}
	return nil
}

func (c *Client) startCDC(ctx context.Context) error {
	connPool, err := c.Conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	// this must be closed only at the end of the sync process
	defer connPool.Release()
	conn := connPool.Conn().PgConn()
	if err := c.createPublicationForTables(ctx, conn); err != nil {
		return err
	}

	sql := fmt.Sprintf("CREATE_REPLICATION_SLOT %s LOGICAL pgoutput EXPORT_SNAPSHOT", c.spec.Name)
	c.createReplicationSlotResult, err = pglogrepl.ParseCreateReplicationSlot(conn.Exec(ctx, sql))
	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return fmt.Errorf("failed to create publication: %w", err)
		}
		if pgErr.Code != "42710" {
			// not recoverable error
			return fmt.Errorf("failed to create replication slot %s with pgerror %s: %w", c.spec.Name, pgErrToStr(pgErr), err)
		}
		// replication slot already exists
		return nil
	}
	
	return nil
}

func (c *Client) listenCDC(ctx context.Context, res chan<- *schema.Resource) error {
	connPool, err := c.Conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer connPool.Release()
	conn := connPool.Conn().PgConn()

	clientXLogPos, err := c.getLastXlogPos(ctx)
	if err != nil {
		return err
	}
	if clientXLogPos == 0 {
		return fmt.Errorf("didn't find last xlog pos")
	}

	if err := pglogrepl.StartReplication(ctx, conn, c.spec.Name, clientXLogPos,
		pglogrepl.StartReplicationOptions{
			PluginArgs: []string{"proto_version '1'", "publication_names '" + c.spec.Name + "'"},
		}); err != nil {
		return fmt.Errorf("failed to start replication: %w", err)
	}

	standbyMessageTimeout := time.Second * 10
	nextStandbyMessageDeadline := time.Now().Add(standbyMessageTimeout)
	relations := map[uint32]*pglogrepl.RelationMessage{}
	typeMap := pgtype.NewMap()
	
	for {
		if time.Now().After(nextStandbyMessageDeadline) {
			err := pglogrepl.SendStandbyStatusUpdate(context.Background(), conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clientXLogPos})
			if err != nil {
				return fmt.Errorf("failed to send standby status update: %w", err)
			}
			nextStandbyMessageDeadline = time.Now().Add(standbyMessageTimeout)
		}

		ctx, cancel := context.WithDeadline(ctx, nextStandbyMessageDeadline)
		rawMsg, err := conn.ReceiveMessage(ctx)
		cancel()
		if err != nil {
			if pgconn.Timeout(err) {
				continue
			}
			return fmt.Errorf("ReceiveMessage failed: %w", err)
		}

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
			if pkm.ReplyRequested {
				nextStandbyMessageDeadline = time.Time{}
			}
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
					return fmt.Errorf("unknown relation ID %d", logicalMsg.RelationID)
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
							return fmt.Errorf("error decoding column data: %w", err)
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
				rel, ok := relations[logicalMsg.RelationID]
				if !ok {
					return fmt.Errorf("unknown relation ID %d", logicalMsg.RelationID)
				}
				values := map[string]interface{}{}
				for idx, col := range logicalMsg.OldTuple.Columns {
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
				log.Printf("DELETE FROM %s.%s: %v", rel.Namespace, rel.RelationName, values)
				resource, err := c.resourceFromCDCValues(rel.RelationName, values)
				if err != nil {
					return err
				}
				res <- resource
			case *pglogrepl.TruncateMessage:	
			case *pglogrepl.TypeMessage:
			case *pglogrepl.OriginMessage:
			default:
				log.Printf("Unknown message type in pgoutput stream: %T", logicalMsg)
			}
			clientXLogPos = xld.WALStart + pglogrepl.LSN(len(xld.WALData))
		}		
	}
}

func (c *Client) createCDCStateTable(ctx context.Context) error {
	if _, err := c.Conn.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS cq_source_pg_cdc"); err != nil {
		return fmt.Errorf("failed to create cq_source_pg_cdc schema: %w", err)
	}
	if _, err := c.Conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS cq_source_pg_cdc.state (slot_name text PRIMARY KEY, x_log_pos bigint)"); err != nil {
		return fmt.Errorf("failed to create cq_source_pg_cdc.state table: %w", err)
	}
	return nil
}

func (c *Client) getLastXlogPos(ctx context.Context) (pglogrepl.LSN, error) {
	var xLogPosStr string
	if err := c.Conn.QueryRow(ctx, "SELECT confirmed_flush_lsn FROM pg_replication_slots WHERE slot_name = $1", c.spec.Name).Scan(&xLogPosStr); err != nil {
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("slot not found: %w", err)
		}
		return 0, fmt.Errorf("failed to get last xlog pos: %w", err)
	}
	xLogPos, err := pglogrepl.ParseLSN(xLogPosStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse last xlog pos: %w", err)
	}
	return xLogPos, nil
}