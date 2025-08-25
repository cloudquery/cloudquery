package client

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

const (
	sqlTableInfoStart = "SELECT table_name, column_name, data_type, is_nullable FROM information_schema.columns WHERE table_schema=CURRENT_SCHEMA() AND UPPER(table_name) = ANY (SELECT COLUMN1 FROM VALUES "
	sqlTableInfoEnd   = ") ORDER BY table_name, ordinal_position"

	sqlShowPrimaryKeys = `SHOW PRIMARY KEYS IN SCHEMA ->> SELECT "table_name", "column_name", "constraint_name", "key_sequence" FROM $1`
	sqlShowUniques     = `SHOW UNIQUE KEYS IN SCHEMA ->> SELECT "table_name", "column_name", "constraint_name", "key_sequence" FROM $1`
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
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(c.spec.MigrateConcurrency)

	var (
		existingTables schema.Tables
		pks            map[string][]constInfo
		uniques        map[string][]constInfo
	)

	g.Go(func() error {
		const limit = 200

		sort.Strings(tableNames)

		var err error
		if len(tableNames) <= limit {
			existingTables, err = c.getTableInfoBatch(ctx, tableNames)
			return err
		}

		for i := 0; i < len(tableNames); i += limit {
			end := i + limit
			if end > len(tableNames) {
				end = len(tableNames)
			}
			batch := tableNames[i:end]
			tbls, err := c.getTableInfoBatch(ctx, batch)
			if err != nil {
				return err
			}
			existingTables = append(existingTables, tbls...)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		pks, err = c.getConstraints(ctx, sqlShowPrimaryKeys)
		if err != nil {
			return fmt.Errorf("failed to get list of primary keys: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		uniques, err = c.getConstraints(ctx, sqlShowUniques)
		if err != nil {
			return fmt.Errorf("failed to get list of unique constraints: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// meld primary keys into existing tables
	{
		for tableName, info := range pks {
			table := existingTables.Get(tableName)
			if table == nil {
				continue
			}
			pkCols := make(map[string]bool)
			for _, pk := range info {
				pkCols[strings.ToUpper(pk.colName)] = true
				table.PkConstraintName = pk.constName
			}
			for i, col := range table.Columns {
				if pkCols[strings.ToUpper(col.Name)] {
					table.Columns[i].PrimaryKey = true
				}
			}
		}
	}

	// meld unique constraints into existing tables
	{
		for tableName, info := range uniques {
			table := existingTables.Get(tableName)
			if table == nil {
				continue
			}
			uniCols := make(map[string]bool)
			for _, u := range info {
				uniCols[strings.ToUpper(u.colName)] = true
			}
			for i, col := range table.Columns {
				if uniCols[strings.ToUpper(col.Name)] {
					table.Columns[i].Unique = true
				}
			}
		}
	}

	return existingTables, nil
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

type constInfo struct {
	tableName string
	colName   string
	constName string
	keySeq    int
}

func (c *Client) getConstraints(ctx context.Context, query string) (map[string][]constInfo, error) {
	list := make(map[string][]constInfo)

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row constInfo

		if err := rows.Scan(
			&row.tableName,
			&row.colName,
			&row.constName,
			&row.keySeq,
		); err != nil {
			return nil, err
		}

		list[strings.ToUpper(row.tableName)] = append(list[strings.ToUpper(row.tableName)], row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
