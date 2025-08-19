package client

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thoas/go-funk"
)

// MigrateTableBatch migrates a table. It forms part of the writer.MixedBatchWriter interface.
func (c *Client) MigrateTableBatch(ctx context.Context, messages message.WriteMigrateTables) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	tables, err := tablesFromMessages(messages)
	if err != nil {
		return err
	}
	pgTables, err := c.listTables(timeoutCtx)
	if err != nil {
		return fmt.Errorf("failed listing postgres tables: %w", err)
	}
	tables = c.normalizeTables(tables)

	safeTables := map[string]bool{}
	for _, msg := range messages {
		// last message takes precedence; we don't actually expect the same table to be
		// in the same batch twice.
		safeTables[msg.Table.Name] = !msg.MigrateForce
	}
	nonAutoMigratableTables := c.nonAutoMigratableTables(tables, pgTables, safeTables)
	if len(nonAutoMigratableTables) > 0 {
		return fmt.Errorf("\nCan't migrate tables automatically, migrate manually or consider using 'migrate_mode: forced'. Non auto migratable tables changes:\n\n%s", schema.GetChangesSummary(nonAutoMigratableTables))
	}

	for _, table := range tables {
		tableName := table.Name
		c.logger.Info().Str("table", tableName).Msg("Migrating table")
		if len(table.Columns) == 0 {
			c.logger.Info().Str("table", tableName).Msg("Table with no columns, skipping")
			continue
		}
		pgTable := pgTables.Get(tableName)
		if pgTable == nil {
			c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
			if err := c.createTableIfNotExist(timeoutCtx, table); err != nil {
				return err
			}
		} else {
			changes := table.GetChanges(pgTable)
			if c.canAutoMigrate(changes) {
				c.logger.Info().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(timeoutCtx, table, changes); err != nil {
					return err
				}
			} else {
				c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, tableName); err != nil {
					return err
				}
				if err := c.createTableIfNotExist(timeoutCtx, table); err != nil {
					return err
				}
			}
		}
		if c.spec.CreatePerformanceIndexes {
			if err := c.createPerformanceIndexes(timeoutCtx, table); err != nil {
				return err
			}
		}
	}
	conn, err := c.conn.Acquire(timeoutCtx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()
	if err := conn.Conn().DeallocateAll(timeoutCtx); err != nil {
		return fmt.Errorf("failed to deallocate all prepared statements: %w", err)
	}
	return nil
}

func (c *Client) normalizeTable(table *schema.Table) *schema.Table {
	normalizedTable := schema.Table{
		Name: table.Name,
	}
	for _, col := range table.Columns {
		if c.pgType == pgTypeCrateDB {
			// CrateDB doesn't allow columns that start with an underscore,
			// so we trim the leading underscore from the column name
			col.Name = strings.TrimLeft(col.Name, "_")
			// CrateDB does not support Unique constraints
			col.Unique = false
		}

		// Postgres doesn't support column names longer than 63 characters
		// and it will automatically truncate them, so we do the same here
		// to make migrations predictable
		if len(col.Name) > 63 {
			col.Name = col.Name[:63]
		}

		if col.PrimaryKey {
			col.NotNull = true
		}
		col.Type = c.PgToSchemaType(c.SchemaTypeToPg(col.Type))

		normalizedTable.Columns = append(normalizedTable.Columns, col)
		// pgTablesToPKConstraints is populated when handling migrate messages
		if entry := c.pgTablesToPKConstraints[table.Name]; entry != nil {
			normalizedTable.PkConstraintName = entry.name
		}
	}

	return &normalizedTable
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) error {
	tableName := table.Name
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeAdd:
			err := c.addColumn(ctx, tableName, change.Current)
			if err != nil {
				return err
			}
		case schema.TableColumnChangeTypeMoveToCQOnly:
			err := c.migrateToCQID(ctx, table, change.Current)
			if err != nil {
				return err
			}
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			err := c.removeUniqueConstraint(ctx, table, change)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (*Client) canAutoMigrate(changes []schema.TableColumnChange) bool {
	// The SDK can detect more granular changes than we can handle
	// We know that when the `TableColumnChangeTypeMoveToCQOnly` is present there will be other changes that were found as well
	// As long as the only change is to remove PK from columns and add it to _cq_id, we can skip handling the changes
	// But we need to make sure there are no other changes
	columnsAddingPK := []string{}
	columnsRemovingPK := []string{}
	cqMigration := false
	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeUpdate:
			if change.Current.Type != change.Previous.Type {
				continue
			}
			if change.Current.PrimaryKey && !change.Previous.PrimaryKey {
				columnsAddingPK = append(columnsAddingPK, change.ColumnName)
			}
			if !change.Current.PrimaryKey && change.Previous.PrimaryKey {
				columnsRemovingPK = append(columnsRemovingPK, change.ColumnName)
			}

		case schema.TableColumnChangeTypeMoveToCQOnly:
			cqMigration = true
		}
	}

	for _, change := range changes {
		switch change.Type {
		case schema.TableColumnChangeTypeRemoveUniqueConstraint:
			continue
		case schema.TableColumnChangeTypeAdd:
			if change.Current.PrimaryKey || change.Current.NotNull {
				return false
			}
		case schema.TableColumnChangeTypeRemove:
			if change.Previous.PrimaryKey || change.Previous.NotNull {
				// nolint:gosimple
				if change.ColumnName == "rowid" {
					// special case for CockroachDB when table has no primary key
					return true
				}
				return false
			}
		case schema.TableColumnChangeTypeMoveToCQOnly:
			continue
		case schema.TableColumnChangeTypeUpdate:
			if cqMigration && ((len(columnsAddingPK) == 1 && columnsAddingPK[0] == schema.CqIDColumn.Name) || funk.Contains(columnsRemovingPK, change.ColumnName)) {
				// We don't need to handle these changes as they are a part of the CQID migration
				continue
			}
			return false
		default:
			return false
		}
	}
	return true
}

// normalize the requested schema to be compatible with what Postgres supports
func (c *Client) normalizeTables(tables schema.Tables) schema.Tables {
	result := make(schema.Tables, len(tables))
	for i, table := range tables {
		result[i] = c.normalizeTable(table)
	}
	return result
}

func (c *Client) nonAutoMigratableTables(tables schema.Tables, pgTables schema.Tables, safeTables map[string]bool) map[string][]schema.TableColumnChange {
	result := make(map[string][]schema.TableColumnChange)
	for _, t := range tables {
		pgTable := pgTables.Get(t.Name)
		if pgTable == nil {
			continue
		}
		changes := t.GetChanges(pgTable)
		if safeTables[t.Name] && !c.canAutoMigrate(changes) {
			result[t.Name] = changes
		}
	}
	return result
}

func (c *Client) dropTable(ctx context.Context, tableName string) error {
	c.logger.Info().Str("table", tableName).Msg("Dropping table")
	sql := "drop table " + tableName
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to drop table %s: %w", tableName, err)
	}
	return nil
}

func (c *Client) migrateToCQID(ctx context.Context, table *schema.Table, _ schema.Column) (err error) {
	// Steps:
	// acquire connection
	var conn *pgxpool.Conn
	conn, err = c.conn.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()
	tableName := table.Name
	sanitizedTableName := pgx.Identifier{tableName}.Sanitize()
	sanitizedPKName, err := c.getPKName(ctx, tableName)
	if err != nil {
		return fmt.Errorf("failed to get primary key name: %w", err)
	}
	// start transaction
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				c.logger.Error().Err(err).Msg("failed to commit transaction")
			}
		}
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				c.logger.Error().Err(rollbackErr).Str("table", tableName).Msg("Failed to rollback transaction")
			}
		}
	}()

	// Drop existing primary key
	_, err = tx.Exec(ctx, "ALTER TABLE "+sanitizedTableName+" DROP CONSTRAINT "+sanitizedPKName)
	if err != nil {
		c.logger.Error().Err(err).Str("table", tableName).Msg("Failed to drop primary key")
		return err
	}
	// Create new Primary Key with CQID
	_, err = tx.Exec(ctx, "ALTER TABLE "+sanitizedTableName+" ADD CONSTRAINT "+sanitizedPKName+" PRIMARY KEY ("+pgx.Identifier{schema.CqIDColumn.Name}.Sanitize()+")")
	if err != nil {
		c.logger.Error().Err(err).Str("table", tableName).Msg("Failed to create new primary key on _cq_id")
		return err
	}

	// CockroachDB doesn't support dropping NOT NULL constraints in the same transaction as the primary key is changed
	// So we have to alter the PK in one transaction and then drop the old NOT NULL constraints in another transaction
	if c.pgType == pgTypeCockroachDB {
		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				c.logger.Error().Err(err).Msg("failed to commit transaction")
			}
		}
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				c.logger.Error().Err(rollbackErr).Str("table", tableName).Msg("Failed to rollback transaction")
			}
		}
		tx, err = conn.BeginTx(ctx, pgx.TxOptions{
			IsoLevel: pgx.Serializable,
		})
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
	}

	c.pgTablesToPKConstraintsMu.Lock()
	entry := c.pgTablesToPKConstraints[tableName]
	if entry == nil {
		entry = new(pkConstraintDetails)
		c.pgTablesToPKConstraints[tableName] = entry
	}
	c.pgTablesToPKConstraintsMu.Unlock()

	for _, colName := range entry.columns {
		_, err = tx.Exec(ctx, "ALTER TABLE "+sanitizedTableName+" ALTER COLUMN "+pgx.Identifier{colName}.Sanitize()+" DROP NOT NULL")
		if err != nil {
			c.logger.Error().Err(err).Str("table", tableName).Str("column", colName).Msg("Failed to drop NOT NULL constraint")
			return err
		}
	}

	return nil
}

func (c *Client) addColumn(ctx context.Context, tableName string, column schema.Column) error {
	c.logger.Info().Str("table", tableName).Str("column", column.Name).Msg("Column doesn't exist, creating")
	columnName := pgx.Identifier{column.Name}.Sanitize()
	columnType := c.SchemaTypeToPg(column.Type)
	sql := "alter table " + tableName + " add column " + columnName + " " + columnType
	if _, err := c.conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("failed to add column %s on table %s: %w", column.Name, tableName, err)
	}
	return nil
}

func (c *Client) createTableIfNotExist(ctx context.Context, table *schema.Table) error {
	var sb strings.Builder
	tableName := table.Name
	sanitizedTableName := pgx.Identifier{tableName}.Sanitize()
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(sanitizedTableName)
	sb.WriteString(" (")
	totalColumns := len(table.Columns)

	primaryKeys := []string{}
	for i, col := range table.Columns {
		pgType := c.SchemaTypeToPg(col.Type)
		columnName := pgx.Identifier{col.Name}.Sanitize()
		fieldDef := columnName + " " + pgType
		if col.Unique {
			fieldDef += " UNIQUE"
		}
		if col.NotNull {
			fieldDef += " NOT NULL"
		}
		sb.WriteString(fieldDef)
		if i != totalColumns-1 {
			sb.WriteString(",")
		}
		if col.PrimaryKey {
			primaryKeys = append(primaryKeys, pgx.Identifier{col.Name}.Sanitize())
		}
	}

	if len(primaryKeys) > 0 {
		// add composite PK constraint on primary key columns
		sb.WriteString(", PRIMARY KEY (")
		sb.WriteString(strings.Join(primaryKeys, ","))
		sb.WriteString(")")
	}
	sb.WriteString(")")
	_, err := c.conn.Exec(ctx, sb.String())
	if err != nil {
		c.logger.Error().Err(err).Str("table", tableName).Str("query", sb.String()).Msg("Failed to create table")
		return fmt.Errorf("failed to create table %s: %w"+sb.String(), tableName, err)
	}
	pkName, err := c.getPKName(ctx, tableName)
	if err != nil {
		return fmt.Errorf("failed to get primary key name: %w", err)
	}
	c.pgTablesToPKConstraints[tableName] = &pkConstraintDetails{
		name:    pkName,
		columns: table.PrimaryKeys(),
	}

	return nil
}

func (c *Client) removeUniqueConstraint(ctx context.Context, table *schema.Table, change schema.TableColumnChange) error {
	// We only support the default unique constraint name
	// If it is using a unique constraint that is not default it means CQ didn't create it so we shouldn't drop it
	indexName := table.Name + "_" + change.ColumnName + "_key"
	sqlStatement := "ALTER TABLE " + pgx.Identifier{table.Name}.Sanitize() + " DROP CONSTRAINT " + pgx.Identifier{indexName}.Sanitize()
	if c.pgType == pgTypeCockroachDB {
		sqlStatement = "DROP INDEX " + pgx.Identifier{indexName}.Sanitize() + " CASCADE"
	}
	_, err := c.conn.Exec(ctx, sqlStatement)
	if err != nil {
		return fmt.Errorf("failed to drop unique constraint on column %s on table %s: %w", change.ColumnName, table.Name, err)
	}
	return nil
}

func (c *Client) getPKName(ctx context.Context, tableName string) (string, error) {
	var pkConstraintName string
	err := c.conn.QueryRow(ctx, `SELECT tco.constraint_name FROM information_schema.table_constraints tco JOIN information_schema.key_column_usage kcu
     ON kcu.constraint_name = tco.constraint_name
     WHERE tco.constraint_type = 'PRIMARY KEY'
    AND kcu.table_name = $1`, tableName).Scan(&pkConstraintName)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return "", err
	}

	return pkConstraintName, nil
}

func (c *Client) createPerformanceIndexes(ctx context.Context, table *schema.Table) error {
	columns := []string{"_cq_source_name", "_cq_sync_time"}

	for _, col := range columns {
		if table.Columns.Get(col) == nil {
			c.logger.Debug().Msgf("no performance indexes created because %s column is not present", col)
			return nil
		}
	}
	// If there are no primary keys, this means that append mode is being used as the write mode as even when table has no PKs the _cq_id is added as a PK
	if len(table.PrimaryKeys()) == 0 {
		c.logger.Debug().Msg("indexes are not needed because no primary keys are present which means that `append` write mode is being used")
		return nil
	}

	indexName := table.Name + "_cqpi"
	if len(table.Name) > 58 && c.pgType != pgTypeCockroachDB {
		indexName = hashTableName(table.Name) + "_cqpi"
	}

	sqlStatement := "CREATE INDEX IF NOT EXISTS " + pgx.Identifier{indexName}.Sanitize() + " ON " + pgx.Identifier{table.Name}.Sanitize() + "(" + pgx.Identifier{columns[0]}.Sanitize() + ", " + pgx.Identifier{columns[1]}.Sanitize() + ")"
	_, err := c.conn.Exec(ctx, sqlStatement)
	if err != nil {
		return fmt.Errorf("failed to create performance index on table %s: %w", table.Name, err)
	}

	return nil
}

func hashTableName(input string) string {
	// Max table name length
	if len(input) > 63 {
		input = input[:63]
	}
	hash := sha256.Sum256([]byte(input))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr[:32]
}
