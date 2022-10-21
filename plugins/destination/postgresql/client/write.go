package client

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func (c *Client) transformValues(values schema.CQTypes) []interface{} {
	pgValues := make([]interface{}, len(values))
	for i, v := range values {
		typ := v.Type()
		switch typ {
		case schema.TypeBool:
			v := v.(*schema.Bool)
			pgValues[i] = pgtype.Bool{
				Bool: v.Bool,
				Status: pgtype.Present,
			}
		case schema.TypeUUID:
			v := v.(*schema.UUID)
			pgValues[i] = pgtype.UUID{
				Bytes: v.Bytes,
				Status: pgtype.Present,
			}
		case schema.TypeString:
			v := v.(*schema.String)
			pgValues[i] = pgtype.Text{
				String: v.String,
				Status: pgtype.Present,
			}
		case schema.TypeInt:
			v := v.(*schema.Int64)
			pgValues[i] = pgtype.Int8{
				Int: v.Int64,
				Status: pgtype.Present,
			}
		default:
			c.logger.Warn().Str("type", typ.String()).Msg("unsupported type")
		}
	}
	return pgValues
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <- chan *schema.DestinationResource) error {
	var sql string
	batch := &pgx.Batch{}

	for r := range res {
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = insert(tables.Get(r.TableName))
		} else {
			sql = upsert(tables.Get(r.TableName))
		}
		values := c.transformValues(r.Data)
		
		batch.Queue(sql, values...)
		if batch.Len() >= c.batchSize {
			fmt.Println(sql)
			br := c.conn.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				if _, ok := err.(*pgconn.PgError); ok {
					atomic.AddUint64(&c.stats.Errors, 1)
					c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
				} else {
					// no recoverable error
					return fmt.Errorf("failed to execute batch: %w", err)
				}
			}
			atomic.AddUint64(&c.stats.Writes, uint64(c.batchSize))
			batch = &pgx.Batch{}
		}
	}

	if batch.Len() > 0 {
		br := c.conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			if _, ok := err.(*pgconn.PgError); ok {
				atomic.AddUint64(&c.stats.Errors, 1)
				c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
			} else {
				// no recoverable error
				return fmt.Errorf("failed to execute batch: %w", err)
			}
		}
		atomic.AddUint64(&c.stats.Writes, uint64(c.batchSize))
	}

	return nil
}

func insert(table *schema.Table) string {
	var sb strings.Builder
<<<<<<< HEAD

	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	// Sort the columns prior to adding data to columns and values arrays
	// Columns need to be in the same order so that the query can be cached during the statement preparation stage
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		columns = append(columns, pgx.Identifier{key}.Sanitize())
		values = append(values, data[key])
	}
=======
>>>>>>> 44257d1fb (feat: Add CQType system)
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
<<<<<<< HEAD
	sort.Strings(columns)
=======
	columns := table.Columns
	columnsLen := len(columns)
>>>>>>> 44257d1fb (feat: Add CQType system)
	for i, c := range columns {
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
		if i < columnsLen - 1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i, _ := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen - 1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func upsert(table *schema.Table) string {
	var sb strings.Builder
	
	sb.WriteString(insert(table))
	columns := table.Columns
	columnsLen := len(columns)

<<<<<<< HEAD
	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))
	// Sort the columns prior to adding data to columns and values arrays
	// Columns need to be in the same order so that the query can be cached during the statement preparation stage
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		columns = append(columns, pgx.Identifier{key}.Sanitize())
		values = append(values, data[key])
	}

	sb.WriteString("insert into ")
	sb.WriteString(table)
	sb.WriteString(" (")
	sort.Strings(columns)
	for i, c := range columns {
		sb.WriteString(c)
		// sb.WriteString("::" + SchemaTypeToPg())
		if i < len(columns)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range values {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < len(values)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	constraintName := fmt.Sprintf("%s_cqpk", table)
=======
	constraintName := fmt.Sprintf("%s_cqpk", table.Name)
>>>>>>> 44257d1fb (feat: Add CQType system)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		if i < columnsLen - 1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String()
}
