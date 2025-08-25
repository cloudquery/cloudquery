package client

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	sqlTableInfoStart = "SELECT table_name, column_name, data_type, is_nullable FROM information_schema.columns WHERE table_schema=CURRENT_SCHEMA() AND UPPER(table_name) = ANY (SELECT COLUMN1 FROM VALUES "
	sqlTableInfoEnd   = ") ORDER BY table_name, ordinal_position"
)

type snowflakeYesNo bool

func (s *snowflakeYesNo) Scan(value any) error {
	if value == nil {
		*s = false
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into snowflakeYesNo", value)
	}

	switch str {
	case "YES":
		*s = true
	case "NO":
		*s = false
	default:
		return fmt.Errorf("failed to scan yes/no string: %s", str)
	}
	return nil
}

func (c *Client) getTableInfo(ctx context.Context, tableNames []string) (schema.Tables, error) {
	sort.Strings(tableNames)

	const limit = 200

	if len(tableNames) <= limit {
		return c.getTableInfoBatch(ctx, tableNames)
	}

	var allTables schema.Tables
	for i := 0; i < len(tableNames); i += limit {
		end := i + limit
		if end > len(tableNames) {
			end = len(tableNames)
		}
		batch := tableNames[i:end]
		tbls, err := c.getTableInfoBatch(ctx, batch)
		if err != nil {
			return nil, err
		}
		allTables = append(allTables, tbls...)
	}
	return allTables, nil
}

func (c *Client) getTableInfoBatch(ctx context.Context, tableNames []string) (schema.Tables, error) {
	infos := make(map[string]*schema.Table, len(tableNames))

	tnAny := make([]any, len(tableNames))
	for i := range tableNames {
		tnAny[i] = strings.ToUpper(tableNames[i])
	}
	completeSQL := sqlTableInfoStart + "(" + strings.Repeat("?,", len(tableNames)-1) + "?)" + sqlTableInfoEnd

	rows, err := c.db.QueryContext(ctx, completeSQL, tnAny...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			tableName string
			colName   string
			colType   string
			nullable  snowflakeYesNo
		)

		if err := rows.Scan(
			&tableName,
			&colName,
			&colType,
			&nullable); err != nil {
			return nil, err
		}

		colType = strings.ToLower(colType)
		info := infos[strings.ToUpper(tableName)]
		if info == nil {
			info = &schema.Table{
				Name:    tableName,
				Columns: make([]schema.Column, 0),
			}
		}
		c := schema.Column{
			Name:    colName,
			Type:    SnowflakeToSchemaType(colType),
			NotNull: !bool(nullable),
		}
		info.Columns = append(info.Columns, c)
		infos[strings.ToUpper(tableName)] = info
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	tbls := make(schema.Tables, 0, len(infos))
	for _, t := range infos {
		tbls = append(tbls, t)
	}
	return tbls, nil
}
