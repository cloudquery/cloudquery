package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

const (
	sqlTableInfo      = "PRAGMA table_info('%s');"
	isColumnUniqueSQL = "select count(*) from duckdb_constraints where table_name = $1 and constraint_type = 'UNIQUE' and constraint_column_names=[$2]"
)

type columnInfo struct {
	index        int
	name         string
	typ          string
	notNull      bool
	defaultValue any
	pk           bool
	unique       bool
}

type tableInfo struct {
	columns []columnInfo
}

func (c *Client) duckdbTables(tables schema.Schemas) (schema.Schemas, error) {
	var schemaTables schema.Schemas
	for _, table := range tables {
		tableName := schema.TableName(table)
		info, err := c.getTableInfo(tableName)
		if err != nil {
			return nil, err
		}
		if info == nil {
			continue
		}
		fields := make([]arrow.Field, len(info.columns))
		for i, col := range info.columns {
			md := make(map[string]string)
			if col.pk {
				md[schema.MetadataPrimaryKey] = schema.MetadataTrue
			} else {
				md[schema.MetadataPrimaryKey] = schema.MetadataFalse
			}
			if col.unique {
				md[schema.MetadataUnique] = schema.MetadataTrue
			} else {
				md[schema.MetadataUnique] = schema.MetadataFalse
			}
			fields[i] = arrow.Field{
				Name:     col.name,
				Type:     c.duckdbTypeToSchema(col.typ),
				Nullable: !col.notNull,
				Metadata: arrow.MetadataFrom(md),
			}
		}
		md := arrow.MetadataFrom(map[string]string{
			schema.MetadataTableName: tableName,
		})
		schemaTables = append(schemaTables, arrow.NewSchema(fields, &md))
	}

	return schemaTables, nil
}

func (c *Client) normalizeColumns(tables schema.Schemas) schema.Schemas {
	var normalized schema.Schemas
	for _, table := range tables {
		fields := table.Fields()
		for i := range fields {
			// In DuckDB, a PK column must be NOT NULL, so we need to make sure that the schema we're comparing to has the same
			// constraint.
			metadata := fields[i].Metadata.ToMap()

			if !c.enabledPks() {
				metadata[schema.MetadataPrimaryKey] = schema.MetadataFalse
				metadata[schema.MetadataUnique] = schema.MetadataFalse
			} else {
				if schema.IsPk(fields[i]) {
					fields[i].Nullable = false
				} else {
					metadata[schema.MetadataPrimaryKey] = schema.MetadataFalse
				}
				if !schema.IsUnique(fields[i]) {
					metadata[schema.MetadataUnique] = schema.MetadataFalse
				}
			}

			fields[i].Metadata = arrow.MetadataFrom(metadata)
			// Since multiple schema types can map to the same duckdb type we need to normalize them to avoid false positives when detecting schema changes
			fields[i].Type = c.duckdbTypeToSchema(c.SchemaTypeToDuckDB(fields[i].Type))
		}
		md := table.Metadata()

		normalized = append(normalized, arrow.NewSchema(fields, &md))
	}

	return normalized
}

func (c *Client) nonAutoMigrableTables(tables schema.Schemas, duckdbTables schema.Schemas) ([]string, [][]schema.FieldChange) {
	var result []string
	var tableChanges [][]schema.FieldChange
	for _, t := range tables {
		tName := schema.TableName(t)
		duckdbTable := duckdbTables.SchemaByName(tName)
		if duckdbTable == nil {
			continue
		}
		changes := schema.GetSchemaChanges(t, duckdbTable)
		if !c.canAutoMigrate(changes) {
			result = append(result, tName)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func (c *Client) autoMigrateTable(table *arrow.Schema, changes []schema.FieldChange) error {
	tableName := schema.TableName(table)
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(tableName, change.Current.Name, c.SchemaTypeToDuckDB(change.Current.Type)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (*Client) canAutoMigrate(changes []schema.FieldChange) bool {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd && (schema.IsPk(change.Current) || !change.Current.Nullable) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeRemove && (schema.IsPk(change.Previous) || !change.Previous.Nullable) {
			return false
		}

		if change.Type == schema.TableColumnChangeTypeUpdate {
			return false
		}
	}
	return true
}

// Migrate migrates to the latest schema
func (c *Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	duckdbTables, err := c.duckdbTables(tables)
	if err != nil {
		return err
	}

	normalizedTables := c.normalizeColumns(tables)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(normalizedTables, duckdbTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %s require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range normalizedTables {
		tableName := schema.TableName(table)
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Fields()) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}
		duckdb := duckdbTables.SchemaByName(tableName)
		if duckdb == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(tableName, table); err != nil {
				return err
			}
			continue
		}

		changes := schema.GetSchemaChanges(table, duckdb)
		if c.canAutoMigrate(changes) {
			c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
			if err := c.autoMigrateTable(table, changes); err != nil {
				return err
			}
		} else {
			c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
			if err := c.recreateTable(table); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) recreateTable(table *arrow.Schema) error {
	tableName := schema.TableName(table)
	sql := "drop table if exists \"" + tableName + "\""
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
	}
	return c.createTableIfNotExist(tableName, table)
}

func (c *Client) addColumn(tableName string, columnName string, columnType string) error {
	sql := "alter table \"" + tableName + "\" add column \"" + columnName + "\" \"" + columnType + `"`
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", columnName, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(tableName string, table *arrow.Schema) error {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" (")
	totalColumns := len(table.Fields())

	var pks []string
	for i, col := range table.Fields() {
		sqlType := c.SchemaTypeToDuckDB(col.Type)
		// TODO: sanitize column name
		fieldDef := `"` + col.Name + `" ` + sqlType
		if schema.IsPk(col) {
			pks = append(pks, col.Name)
		}
		if schema.IsUnique(col) && c.enabledPks() {
			fieldDef += " UNIQUE"
		}
		if !col.Nullable {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
	}
	if len(pks) > 0 && c.enabledPks() {
		sb.WriteString(", PRIMARY KEY (")
		for i, pk := range pks {
			sb.WriteString(`"` + pk + `"`)
			if i != len(pks)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.db.Exec(sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table with '%s': %w", sb.String(), err)
	}
	return nil
}

func (c *Client) isColumnUnique(tableName string, columName string) (bool, error) {
	rows, err := c.db.Query(isColumnUniqueSQL, tableName, columName)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			return false, err
		}
		if n == 1 {
			return true, nil
		}
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func (c *Client) getTableInfo(tableName string) (*tableInfo, error) {
	info := tableInfo{}
	rows, err := c.db.Query(fmt.Sprintf(sqlTableInfo, tableName))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("Table with name %s does not exist!", tableName)) {
			// Table doesn't exist
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		colInfo := columnInfo{}
		if err := rows.Scan(
			&colInfo.index,
			&colInfo.name,
			&colInfo.typ,
			&colInfo.notNull,
			&colInfo.defaultValue,
			&colInfo.pk); err != nil {
			return nil, err
		}
		colInfo.typ = strings.ToLower(colInfo.typ)
		colInfo.unique, err = c.isColumnUnique(tableName, colInfo.name)
		if err != nil {
			return nil, err
		}
		info.columns = append(info.columns, colInfo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(info.columns) == 0 {
		// Table doesn't exist
		return nil, nil
	}
	return &info, nil
}

func (c *Client) enabledPks() bool {
	return c.spec.WriteMode == specs.WriteModeOverwrite || c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale
}
