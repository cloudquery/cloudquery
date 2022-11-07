package client

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

var cqStatusToPgStatus = map[schema.Status]pgtype.Status{
	schema.Null:      pgtype.Null,
	schema.Undefined: pgtype.Null,
	schema.Present:   pgtype.Present,
}

func cqToPgTextArray(c *schema.TextArray) pgtype.TextArray {
	r := pgtype.TextArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.Str, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgInt8Array(c *schema.Int8Array) pgtype.Int8Array {
	r := pgtype.Int8Array{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Int8{Int: v.Int, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgUUIDArray(c *schema.UUIDArray) pgtype.UUIDArray {
	r := pgtype.UUIDArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.UUID{Bytes: v.Bytes, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToCCCIDRArray(c *schema.CIDRArray) pgtype.InetArray {
	r := pgtype.InetArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgCIDRArray(c *schema.CIDRArray) pgtype.CIDRArray {
	r := pgtype.CIDRArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.CIDR{IPNet: v.IPNet, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToCCMacaddrArray(c *schema.MacaddrArray) pgtype.TextArray {
	r := pgtype.TextArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.String(), Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgMacaddrArray(c *schema.MacaddrArray) pgtype.MacaddrArray {
	r := pgtype.MacaddrArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Macaddr{Addr: v.Addr, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgInetArray(c *schema.InetArray) pgtype.InetArray {
	r := pgtype.InetArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: cqStatusToPgStatus[v.Status]})
	}
	r.Status = cqStatusToPgStatus[c.Status]
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func (c *Client) transformValues(table *schema.Table, values schema.CQTypes) []interface{} {
	pgValues := make([]interface{}, len(values))
	for i, v := range values {
		switch t := v.(type) {
		case *schema.Bool:
			pgValues[i] = pgtype.Bool{
				Bool:   t.Bool,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.Int8:
			pgValues[i] = pgtype.Int8{
				Int:    t.Int,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.Float8:
			pgValues[i] = pgtype.Float8{
				Float:  t.Float,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.UUID:
			pgValues[i] = pgtype.UUID{
				Bytes:  t.Bytes,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.Text:
			pgValues[i] = pgtype.Text{
				String: t.Str,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.Bytea:
			pgValues[i] = pgtype.Bytea{
				Bytes:  t.Bytes,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.TextArray:
			pgValues[i] = cqToPgTextArray(t)
		case *schema.Int8Array:
			pgValues[i] = cqToPgInt8Array(t)
		case *schema.Timestamptz:
			pgValues[i] = pgtype.Timestamptz{
				Time:   t.Time,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.JSON:
			pgValues[i] = pgtype.JSON{
				Bytes:  t.Bytes,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.UUIDArray:
			pgValues[i] = cqToPgUUIDArray(t)
		case *schema.Inet:
			pgValues[i] = pgtype.Inet{
				IPNet:  t.IPNet,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.CIDR:
			pgValues[i] = pgtype.CIDR{
				IPNet:  t.IPNet,
				Status: cqStatusToPgStatus[t.Status],
			}
		case *schema.CIDRArray:
			if c.pgType == pgTypeCockroachDB {
				pgValues[i] = cqToCCCIDRArray(t)
			} else {
				pgValues[i] = cqToPgCIDRArray(t)
			}
		case *schema.Macaddr:
			if c.pgType == pgTypeCockroachDB {
				pgValues[i] = pgtype.Text{
					String: t.String(),
					Status: cqStatusToPgStatus[t.Status],
				}
			} else {
				pgValues[i] = pgtype.Macaddr{
					Addr:   t.Addr,
					Status: cqStatusToPgStatus[t.Status],
				}
			}
		case *schema.MacaddrArray:
			if c.pgType == pgTypeCockroachDB {
				pgValues[i] = cqToCCMacaddrArray(t)
			} else {
				pgValues[i] = cqToPgMacaddrArray(t)
			}
		case *schema.InetArray:
			pgValues[i] = cqToPgInetArray(t)
		default:
			c.logger.Warn().Str("table", table.Name).Str("column", table.Columns[i].Name).Msgf("unknown type %T", v)
			c.metrics.Errors++
		}
	}
	return pgValues
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *schema.DestinationResource) error {
	var sql string
	batch := &pgx.Batch{}

	for r := range res {
		table := tables.Get(r.TableName)
		if table == nil {
			panic(fmt.Errorf("table %s not found", r.TableName))
		}
		if c.spec.WriteMode == specs.WriteModeAppend {
			sql = c.insert(table)
		} else {
			sql = c.upsert(table)
		}
		values := c.transformValues(table, r.Data)

		batch.Queue(sql, values...)
		if batch.Len() >= c.batchSize {
			br := c.conn.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				if _, ok := err.(*pgconn.PgError); !ok {
					// not recoverable error
					return fmt.Errorf("failed to execute batch: %w", err)
				}
				atomic.AddUint64(&c.metrics.Errors, 1)
				c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
			}
			atomic.AddUint64(&c.metrics.Writes, uint64(c.batchSize))
			batch = &pgx.Batch{}
		}
	}

	if batch.Len() > 0 {
		br := c.conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			if _, ok := err.(*pgconn.PgError); !ok {
				// no recoverable error
				return fmt.Errorf("failed to execute batch: %w", err)
			}
			c.logger.Error().Err(err).Msgf("failed to execute batch with pgerror")
		}
		atomic.AddUint64(&c.metrics.Writes, uint64(c.batchSize))
	}

	return nil
}

func (c *Client) filterSupportedColumns(columns schema.ColumnList) schema.ColumnList {
	filtered := make(schema.ColumnList, 0, len(columns))
	for _, col := range columns {
		if c.SchemaTypeToPg(col.Type) != "" {
			filtered = append(filtered, col)
		}
	}
	return filtered
}

func (c *Client) insert(table *schema.Table) string {
	var sb strings.Builder
	sb.WriteString("insert into ")
	sb.WriteString(pgx.Identifier{table.Name}.Sanitize())
	sb.WriteString(" (")
	columns := c.filterSupportedColumns(table.Columns)
	columnsLen := len(columns)
	for i, c := range columns {
		sb.WriteString(pgx.Identifier{c.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(") values (")
		}
	}
	for i := range columns {
		sb.WriteString(fmt.Sprintf("$%d", i+1))
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(")")
		}
	}
	return sb.String()
}

func (c *Client) upsert(table *schema.Table) string {
	var sb strings.Builder

	sb.WriteString(c.insert(table))
	columns := c.filterSupportedColumns(table.Columns)
	columnsLen := len(columns)

	constraintName := fmt.Sprintf("%s_cqpk", table.Name)
	sb.WriteString(" on conflict on constraint ")
	sb.WriteString(constraintName)
	sb.WriteString(" do update set ")
	for i, column := range columns {
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		sb.WriteString("=excluded.") // excluded references the new values
		sb.WriteString(pgx.Identifier{column.Name}.Sanitize())
		if i < columnsLen-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString("")
		}
	}

	return sb.String()
}
