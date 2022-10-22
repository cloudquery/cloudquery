package client

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/cqtypes"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)


func CQTextArray(c *cqtypes.TextArray) pgtype.TextArray {
	r := pgtype.TextArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.String, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func CQInt8Array(c *cqtypes.Int8Array) pgtype.Int8Array {
	r := pgtype.Int8Array{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Int8{Int: v.Int, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func CQUUIDArray(c *cqtypes.UUIDArray) pgtype.UUIDArray {
	r := pgtype.UUIDArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.UUID{Bytes: v.Bytes, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func CQCIDRArray(c *cqtypes.CIDRArray) pgtype.CIDRArray {
	r := pgtype.CIDRArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.CIDR{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func CQMacaddrArray(c *cqtypes.MacaddrArray) pgtype.MacaddrArray {
	r := pgtype.MacaddrArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Macaddr{Addr: v.Addr, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func (c *Client) transformValues(table *schema.Table, values schema.CQTypes) []interface{} {
	pgValues := make([]interface{}, len(values))
	for i, v := range values {
		switch t := v.(type) {
		case *cqtypes.Bool:
			pgValues[i] = pgtype.Bool{
				Bool: t.Bool,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.Int8:
			pgValues[i] = pgtype.Int8{
				Int: t.Int,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.Float8:
			pgValues[i] = pgtype.Float8{
				Float: t.Float,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.UUID:
			pgValues[i] = pgtype.UUID{
				Bytes: t.Bytes,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.Text:
			pgValues[i] = pgtype.Text{
				String: t.String,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.Bytea:
			pgValues[i] = pgtype.Bytea{
				Bytes: t.Bytes,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.TextArray:
			pgValues[i] = CQTextArray(t)
		case *cqtypes.Int8Array:
			pgValues[i] = CQInt8Array(t)
		case *cqtypes.Timestamptz:
			pgValues[i] = pgtype.Timestamptz{
				Time: t.Time,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.JSON:
			pgValues[i] = pgtype.JSON{
				Bytes: t.Bytes,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.UUIDArray:
			pgValues[i] = CQUUIDArray(t)
		case *cqtypes.Inet:
			pgValues[i] = pgtype.Inet{
				IPNet: t.IPNet,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.CIDR:
			pgValues[i] = pgtype.CIDR{
				IPNet: t.IPNet,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.CIDRArray:
			pgValues[i] = CQCIDRArray(t)
		case *cqtypes.Macaddr:
			pgValues[i] = pgtype.Macaddr{
				Addr: t.Addr,
				Status: pgtype.Status(t.Status),
			}
		case *cqtypes.MacaddrArray:
			pgValues[i] = CQMacaddrArray(t)
		default:
			pgValues[i] = pgtype.Text{
				Status: pgtype.Null,
			}
			c.stats.Errors++
			panic(fmt.Errorf("unsupported type %T at index %d column_name %s table_name %s", v, i, table.Columns[i].Name, table.Name))
			// c.logger.Error().Msgf()
		}
	}
	return pgValues
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <- chan *schema.DestinationResource) error {
	var sql string
	batch := &pgx.Batch{}

	for r := range res {
		table := tables.Get(r.TableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", r.TableName))
		}
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = insert(table)
		} else {
			sql = upsert(table)
		}
		values := c.transformValues(table, r.Data)
		batch.Queue(sql, values...)
		if batch.Len() >= c.batchSize {
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
