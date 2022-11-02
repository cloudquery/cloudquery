package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/cqtypes"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func cqToSqlTextArray(c *cqtypes.TextArray) pgtype.TextArray {
	r := pgtype.TextArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.Str, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgInt8Array(c *cqtypes.Int8Array) pgtype.Int8Array {
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

func cqToPgUUIDArray(c *cqtypes.UUIDArray) pgtype.UUIDArray {
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

func cqToCCCIDRArray(c *cqtypes.CIDRArray) pgtype.InetArray {
	r := pgtype.InetArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgCIDRArray(c *cqtypes.CIDRArray) pgtype.CIDRArray {
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

func cqToCCMacaddrArray(c *cqtypes.MacaddrArray) pgtype.TextArray {
	r := pgtype.TextArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.String(), Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func cqToPgMacaddrArray(c *cqtypes.MacaddrArray) pgtype.MacaddrArray {
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

func cqToPgInetArray(c *cqtypes.InetArray) pgtype.InetArray {
	r := pgtype.InetArray{}
	for _, v := range c.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(c.Status)
	for _, d := range c.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return r
}

func (c *Client) transformValues(table *schema.Table, values cqtypes.CQTypes) []interface{} {
	pgValues := make([]interface{}, len(values))
	for i, v := range values {
		switch t := v.(type) {
		case *cqtypes.Bool:
			pgValues[i] = t.Bool
		case *cqtypes.Int8:
			pgValues[i] = t.Int
		case *cqtypes.Float8:
			pgValues[i] = t.Float
		case *cqtypes.UUID:
			pgValues[i] = t.String()
		case *cqtypes.Text:
			pgValues[i] = t.Str
		case *cqtypes.Bytea:
			pgValues[i] = t.Bytes
		case *cqtypes.TextArray:
			pgValues[i] = t.String()
		case *cqtypes.Int8Array:
			pgValues[i] = t.String()
		case *cqtypes.Timestamptz:
			pgValues[i] = t.String()
		case *cqtypes.JSON:
			pgValues[i] = t.String()
		case *cqtypes.UUIDArray:
			pgValues[i] = t.String()
		case *cqtypes.Inet:
			pgValues[i] = t.String()
		case *cqtypes.CIDR:
			pgValues[i] = t.String()
		case *cqtypes.CIDRArray:
			pgValues[i] = t.String()
		case *cqtypes.Macaddr:
			pgValues[i] = t.String()
		case *cqtypes.MacaddrArray:
			pgValues[i] = t.String()
		case *cqtypes.InetArray:
			pgValues[i] = t.String()
		default:
			c.logger.Warn().Str("table", table.Name).Str("column", table.Columns[i].Name).Msgf("unknown type %T", v)
			c.metrics.Errors++
		}
	}
	return pgValues
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *schema.DestinationResource) error {
	var sql string
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
		if _, err := c.db.Exec(sql, values...); err != nil {
			return fmt.Errorf("failed to write to table %s: %w", table.Name, err)
		}
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
