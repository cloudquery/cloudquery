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


func (c *Client) startCDC(ctx context.Context) error {
	connPool, err := c.Conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	// this must be closed only at the end of the sync process
	defer connPool.Release()
	conn := connPool.Conn().PgConn()
	tables := strings.Join(c.Tables.TableNames(), ",")
	sql := fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s", pgx.Identifier{c.spec.Name}.Sanitize(), tables)
	reader := conn.Exec(ctx, sql)
	_, err = reader.ReadAll()
	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return fmt.Errorf("failed to create publication: %w", err)
		}
		if pgErr.Code != "42710" {
			// not recoverable error
			return fmt.Errorf("failed to create publication with pgerror %s: %w", pgErrToStr(pgErr), err)
		}
	}

	clientXLogPos, err := c.getLastXlogPos(ctx)
	if err != nil {
		return err
	}
	if clientXLogPos != 0 {
		return nil
	}

	sysident, err := pglogrepl.IdentifySystem(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to identify system: %w", err)
	}

	sql = fmt.Sprintf("CREATE_REPLICATION_SLOT %s LOGICAL pgoutput EXPORT_SNAPSHOT", c.spec.Name)
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
	}

	if err := pglogrepl.StartReplication(ctx, conn, c.spec.Name, sysident.XLogPos,
		pglogrepl.StartReplicationOptions{
			PluginArgs: []string{"proto_version '1'", "publication_names '" + c.spec.Name + "'"},
		}); err != nil {
		return fmt.Errorf("failed to start replication: %w", err)
	}

	if err := c.setLastXlogPos(ctx, sysident.XLogPos); err != nil {
		return fmt.Errorf("failed to set last xlog pos: %w", err)
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
			if err := c.setLastXlogPos(ctx, clientXLogPos); err != nil {
				return fmt.Errorf("failed to set last xlog pos: %w", err)
			}
			// log.Println("Sent Standby status message")
			nextStandbyMessageDeadline = time.Now().Add(standbyMessageTimeout)
		}

		ctx, cancel := context.WithDeadline(context.Background(), nextStandbyMessageDeadline)
		rawMsg, err := conn.ReceiveMessage(ctx)
		cancel()
		if err != nil {
			if pgconn.Timeout(err) {
				continue
			}
			log.Fatalln("ReceiveMessage failed:", err)
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
			// c.logger.Info().Msgf("Primary Keepalive Message =>", "ServerWALEnd:", pkm.ServerWALEnd, "ServerTime:", pkm.ServerTime, "ReplyRequested:", pkm.ReplyRequested)	
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
				// logicalMsg.
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
	var xLogPos uint64
	if err := c.Conn.QueryRow(ctx, "SELECT x_log_pos FROM cq_source_pg_cdc.state WHERE slot_name = $1", c.spec.Name).Scan(&xLogPos); err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to get last xlog pos: %w", err)
	}
	return pglogrepl.LSN(xLogPos), nil
}

func (c *Client) setLastXlogPos(ctx context.Context, xLogPos pglogrepl.LSN) error {
	if _, err := c.Conn.Exec(ctx, "INSERT INTO cq_source_pg_cdc.state (slot_name, x_log_pos) VALUES ($1, $2) ON CONFLICT (slot_name) DO UPDATE SET x_log_pos = $2", c.spec.Name, uint64(xLogPos)); err != nil {
		return fmt.Errorf("failed to set last xlog pos: %w", err)
	}
	return nil
}