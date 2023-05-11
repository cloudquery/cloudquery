package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

const (
	sqlTableInfo = "PRAGMA table_info('%s');"
)

type columnInfo struct {
	index        int
	name         string
	typ          string
	notNull      bool
	defaultValue any
	pk           int
}

type tableInfo struct {
	columns []columnInfo
}

func (c *Client) sqliteTables(schemas schema.Schemas) (schema.Schemas, error) {
	var schemaTables schema.Schemas
	for _, sc := range schemas {
		var fields []arrow.Field
		tableName := schema.TableName(sc)
		if tableName == "" {
			return nil, fmt.Errorf("schema %s has no table name", sc.String())
		}
		info, err := c.getTableInfo(tableName)
		if info == nil {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, col := range info.columns {
			var fieldMetadata schema.MetadataFieldOptions
			if col.pk != 0 {
				fieldMetadata.PrimaryKey = true
			}
			fields = append(fields, arrow.Field{
				Name:     col.name,
				Type:     c.sqliteTypeToArrowType(col.typ),
				Nullable: !col.notNull,
				Metadata: schema.NewFieldMetadataFromOptions(fieldMetadata),
			})
		}
		var tableMetadata schema.MetadataSchemaOptions
		tableMetadata.TableName = tableName
		m := schema.NewSchemaMetadataFromOptions(tableMetadata)
		schemaTables = append(schemaTables, arrow.NewSchema(fields, &m))
	}
	return schemaTables, nil
}

func (c *Client) normalizeSchemas(scs schema.Schemas) schema.Schemas {
	var normalized schema.Schemas
	for _, sc := range scs {
		fields := make([]arrow.Field, 0)
		for _, f := range sc.Fields() {
			keys := make([]string, 0)
			values := make([]string, 0)
			origKeys := f.Metadata.Keys()
			origValues := f.Metadata.Values()
			for k, v := range origKeys {
				if v != schema.MetadataUnique {
					keys = append(keys, v)
					values = append(values, origValues[k])
				}
			}
			fields = append(fields, arrow.Field{
				Name:     f.Name,
				Type:     c.arrowTypeToSqlite(f.Type),
				Nullable: f.Nullable,
				Metadata: arrow.NewMetadata(keys, values),
			})
		}

		md := sc.Metadata()
		normalized = append(normalized, arrow.NewSchema(fields, &md))
	}

	return normalized
}

func (c *Client) nonAutoMigrableTables(tables schema.Schemas, sqliteTables schema.Schemas) ([]string, [][]schema.FieldChange) {
	var result []string
	var tableChanges [][]schema.FieldChange
	for _, t := range tables {
		tableName := schema.TableName(t)
		sqliteTable := sqliteTables.SchemaByName(tableName)
		if sqliteTable == nil {
			continue
		}
		changes := schema.GetSchemaChanges(t, sqliteTable)
		if !c.canAutoMigrate(changes) {
			result = append(result, tableName)
			tableChanges = append(tableChanges, changes)
		}
	}
	return result, tableChanges
}

func (c *Client) autoMigrateTable(table *arrow.Schema, changes []schema.FieldChange) error {
	for _, change := range changes {
		if change.Type == schema.TableColumnChangeTypeAdd {
			if err := c.addColumn(schema.TableName(table), change.Current.Name, c.arrowTypeToSqliteStr(change.Current.Type)); err != nil {
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

// This is the responsibility of the CLI of the client to lock before running migration
func (c *Client) Migrate(ctx context.Context, schemas schema.Schemas) error {
	schemas = c.normalizeSchemas(schemas)
	sqliteTables, err := c.sqliteTables(schemas)
	if err != nil {
		return err
	}

	if c.spec.MigrateMode != specs.MigrateModeForced {
		nonAutoMigrableTables, changes := c.nonAutoMigrableTables(schemas, sqliteTables)
		if len(nonAutoMigrableTables) > 0 {
			return fmt.Errorf("tables %s with changes %v require force migration. use 'migrate_mode: forced'", strings.Join(nonAutoMigrableTables, ","), changes)
		}
	}

	for _, table := range schemas {
		tableName := schema.TableName(table)
		if tableName == "" {
			return fmt.Errorf("schema %s has no table name", table.String())
		}
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Fields()) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}

		sqlite := sqliteTables.SchemaByName(tableName)
		if sqlite == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(table); err != nil {
				return err
			}
		} else {
			changes := schema.GetSchemaChanges(table, sqlite)
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
	}

	return nil
}

func (c *Client) recreateTable(table *arrow.Schema) error {
	tableName, ok := table.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return fmt.Errorf("schema %s has no table name", table.String())
	}
	sql := "drop table if exists \"" + tableName + "\""
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
	}
	return c.createTableIfNotExist(table)
}

func (c *Client) addColumn(tableName string, columnName string, columnType string) error {
	sql := "alter table \"" + tableName + "\" add column \"" + columnName + "\" \"" + columnType + `"`
	if _, err := c.db.Exec(sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", columnName, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(sc *arrow.Schema) error {
	var sb strings.Builder
	tableName, ok := sc.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return fmt.Errorf("schema %s has no table name", sc.String())
	}
	// TODO sanitize tablename
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(`"` + tableName + `"`)
	sb.WriteString(" (")
	totalColumns := len(sc.Fields())

	primaryKeys := []string{}
	for i, col := range sc.Fields() {
		sqlType := c.arrowTypeToSqliteStr(col.Type)
		if sqlType == "" {
			c.logger.Warn().Str("table", tableName).Str("column", col.Name).Msg("Column type is not supported, skipping")
			continue
		}
		// TODO: sanitize column name
		fieldDef := `"` + col.Name + `" ` + sqlType
		if !col.Nullable {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}

		if c.enabledPks() && schema.IsPk(col) {
			primaryKeys = append(primaryKeys, `"`+col.Name+`"`)
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", CONSTRAINT ")
		sb.WriteString(tableName)
		sb.WriteString("_cqpk PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.db.Exec(sb.String())
	if err != nil {
		return fmt.Errorf("failed to create table with '%s': %w", sb.String(), err)
	}
	return nil
}

func (c *Client) getTableInfo(tableName string) (*tableInfo, error) {
	info := tableInfo{}
	rows, err := c.db.Query(fmt.Sprintf(sqlTableInfo, tableName))
	if err != nil {
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
