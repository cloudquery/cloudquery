package client

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/jackc/pgx/v5/pgtype"
)

func tablesWithPks(filteredTables schema.Tables) []string {
	var tables []string
	for _, table := range filteredTables {
		if len(table.PrimaryKeys()) > 0 {
			tables = append(tables, pgx.Identifier{table.Name}.Sanitize())
		}
	}
	return tables
}

func (c *Client) createPublicationForTables(ctx context.Context, filteredTables schema.Tables, conn *pgconn.PgConn) error {
	tables := tablesWithPks(filteredTables)
	if len(tables) == 0 {
		return fmt.Errorf("cdc is enabled but no tables with primary keys were found")
	}
	sql := fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s", pgx.Identifier{c.cdcId}.Sanitize(), strings.Join(tables, ","))
	reader := conn.Exec(ctx, sql)
	if _, err := reader.ReadAll(); err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return fmt.Errorf("failed to create publication: %w", err)
		}
		// 42710 means publication already exists
		if pgErr.Code != "42710" {
			// not recoverable error
			return fmt.Errorf("failed to create publication with pgerror %s: %w", pgErrToStr(pgErr), err)
		}
		sql = fmt.Sprintf("ALTER PUBLICATION %s SET TABLE %s", pgx.Identifier{c.cdcId}.Sanitize(), strings.Join(tables, ","))
		reader := conn.Exec(ctx, sql)
		if _, err := reader.ReadAll(); err != nil {
			return fmt.Errorf("failed to alter publication: %w", err)
		}
	}
	return nil
}

func (c *Client) startCDC(ctx context.Context, filteredTables schema.Tables, conn *pgconn.PgConn) (string, error) {
	if err := c.createPublicationForTables(ctx, filteredTables, conn); err != nil {
		return "", err
	}
	replicationName := pgx.Identifier{getReplicationName(c.cdcId)}.Sanitize()
	sql := fmt.Sprintf("CREATE_REPLICATION_SLOT %s LOGICAL pgoutput EXPORT_SNAPSHOT", replicationName)
	createReplicationSlotResult, err := pglogrepl.ParseCreateReplicationSlot(conn.Exec(ctx, sql))
	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) {
			// not recoverable error
			return "", fmt.Errorf("failed to create replication slot: %w", err)
		}
		if pgErr.Code != "42710" {
			// not recoverable error
			return "", fmt.Errorf("failed to create replication slot %s with pgerror %s: %w", replicationName, pgErrToStr(pgErr), err)
		}
		// replication slot already exists
		return "", nil
	}
	return createReplicationSlotResult.SnapshotName, nil
}

func getReplicationName(specName string) string {
	return strings.ReplaceAll(specName, "-", "_")
}

func (c *Client) listenCDC(ctx context.Context, res chan<- message.SyncMessage) error {
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

	if err := pglogrepl.StartReplication(ctx, conn, getReplicationName(c.cdcId), clientXLogPos,
		pglogrepl.StartReplicationOptions{
			PluginArgs: []string{"proto_version '1'", "publication_names '" + c.cdcId + "'"},
		}); err != nil {
		return fmt.Errorf("failed to start replication: %w", err)
	}

	standbyMessageTimeout := time.Second * 10
	nextStandbyMessageDeadline := time.Now().Add(standbyMessageTimeout)
	relations := map[uint32]*pglogrepl.RelationMessage{}
	typeMap := pgtype.NewMap()

	for {
		if time.Now().After(nextStandbyMessageDeadline) {
			err := pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clientXLogPos})
			if err != nil {
				return fmt.Errorf("failed to send standby status update: %w", err)
			}
			nextStandbyMessageDeadline = time.Now().Add(standbyMessageTimeout)
		}

		rawMsg, err := conn.ReceiveMessage(ctx)
		if err != nil {
			// send update with latest xlog pos
			pglogreplErr := pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clientXLogPos})
			if pglogreplErr != nil {
				c.logger.Error().Err(pglogreplErr).Msg("failed to send standby status update")
			}
			return err
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
				err := pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{WALWritePosition: clientXLogPos})
				if err != nil {
					return fmt.Errorf("failed to send standby status update: %w", err)
				}
				nextStandbyMessageDeadline = time.Now().Add(standbyMessageTimeout)
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
				values := map[string]any{}
				for idx, col := range logicalMsg.Tuple.Columns {
					colName := rel.Columns[idx].Name
					switch col.DataType {
					case 'n': // null
						values[colName] = nil
					case 'u': // unchanged toast
						// This TOAST value was not changed. TOAST values are not stored in the tuple, and logical replication doesn't want to spend a disk read to fetch its value for you.
					case 't': // text
						val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
						if err != nil {
							return fmt.Errorf("error decoding column data: %w", err)
						}
						values[colName] = val
					}
				}
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
				values := map[string]any{}
				for idx, col := range logicalMsg.NewTuple.Columns {
					colName := rel.Columns[idx].Name
					switch col.DataType {
					case 'n': // null
						values[colName] = nil
					case 'u': // unchanged toast
						// This TOAST value was not changed. TOAST values are not stored in the tuple, and logical replication doesn't want to spend a disk read to fetch its value for you.
					case 't': // text
						val, err := decodeTextColumnData(typeMap, col.Data, rel.Columns[idx].DataType)
						if err != nil {
							return fmt.Errorf("error decoding column data: %w", err)
						}
						values[colName] = val
					}
				}
				resource, err := c.resourceFromCDCValues(rel.RelationName, values)
				if err != nil {
					return err
				}
				res <- resource
			case *pglogrepl.DeleteMessage:
			case *pglogrepl.TruncateMessage:
			case *pglogrepl.TypeMessage:
			case *pglogrepl.OriginMessage:
			default:
				c.logger.Error().Msgf("Unknown message type in pgoutput stream: %T", logicalMsg)
			}
			clientXLogPos = xld.WALStart + pglogrepl.LSN(len(xld.WALData))
		}
	}
}

func (c *Client) getLastXlogPos(ctx context.Context) (pglogrepl.LSN, error) {
	var xLogPosStr string
	replicationName := getReplicationName(c.cdcId)
	if err := c.Conn.QueryRow(ctx, "SELECT confirmed_flush_lsn FROM pg_replication_slots WHERE slot_name = $1", replicationName).Scan(&xLogPosStr); err != nil {
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

func decodeTextColumnData(mi *pgtype.Map, data []byte, dataType uint32) (any, error) {
	if dt, ok := mi.TypeForOID(dataType); ok {
		return dt.Codec.DecodeValue(mi, dataType, pgtype.TextFormatCode, data)
	}
	return string(data), nil
}

func (c *Client) resourceFromCDCValues(tableName string, values map[string]any) (message.SyncMessage, error) {
	table := c.tables.Get(tableName)
	arrowSchema := table.ToArrowSchema()
	builder := array.NewRecordBuilder(memory.DefaultAllocator, arrowSchema)
	transformers := transformersForSchema(arrowSchema)

	for i, col := range table.Columns {
		val, err := transformers[i](values[col.Name])
		if err != nil {
			return nil, err
		}

		s := scalar.NewScalar(arrowSchema.Field(i).Type)
		if err := s.Set(val); err != nil {
			return nil, fmt.Errorf("error setting value for column %s: %w", col.Name, err)
		}

		scalar.AppendToBuilder(builder.Field(i), s)
	}
	return &message.SyncInsert{Record: builder.NewRecord()}, nil
}
