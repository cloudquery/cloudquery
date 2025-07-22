package client

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/typeconv"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
)

type tableChanges struct {
	alreadyExists         bool
	forcedMigrationNeeded bool
	changes               []schema.TableColumnChange
}

// MigrateTables relies on the CLI/client to lock before running migration.
func (c *Client) MigrateTables(ctx context.Context, messages message.WriteMigrateTables) error {
	have, err := retryGetTableDefinitions(ctx, c.logger, c.database, c.conn, messages)
	if err != nil {
		return err
	}

	want, err := typeconv.CanonizedTables(messages)
	if err != nil {
		return err
	}

	tablesWeCanForceMigrate := map[string]bool{}
	for _, msg := range messages {
		// last message takes precedence; we don't actually expect the same table to be
		// in the same batch twice.
		tablesWeCanForceMigrate[msg.Table.Name] = msg.MigrateForce
	}

	allTablesChanges, err := c.allTablesChanges(ctx, want, have)
	if err != nil {
		return err
	}

	nonAutoMigratableTables := lo.Filter(lo.Keys(allTablesChanges), func(table string, _ int) bool {
		return allTablesChanges[table].forcedMigrationNeeded && !tablesWeCanForceMigrate[table]
	})
	if len(nonAutoMigratableTables) > 0 {
		changes := lo.Map(nonAutoMigratableTables, func(table string, _ int) []schema.TableColumnChange {
			return allTablesChanges[table].changes
		})
		return fmt.Errorf("tables %s with changes %v require migration. Migrate manually or consider using 'migrate_mode: forced'", strings.Join(nonAutoMigratableTables, ","), changes)
	}

	const maxConcurrentMigrate = 10
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(maxConcurrentMigrate)

	for _, want := range want {
		eg.Go(func() (err error) {
			c.logger.Info().Str("table", want.Name).Msg("Migrating table started")
			defer func() {
				c.logger.Err(err).Str("table", want.Name).Msg("Migrating table done")
			}()
			if len(want.Columns) == 0 {
				c.logger.Warn().Str("table", want.Name).Msg("Table with no columns, skipping")
				return nil
			}

			tableName := want.Name
			tableChanges := allTablesChanges[tableName]
			if !tableChanges.alreadyExists {
				c.logger.Info().Str("table", tableName).Msg("Table doesn't exist, creating")
				return c.createTable(ctx, want, c.spec.Partition, c.spec.OrderBy)
			}

			if tableChanges.forcedMigrationNeeded {
				c.logger.Info().Str("table", tableName).Msg("Table exists, force migration required")
				if err := c.dropTable(ctx, want); err != nil {
					return err
				}
				return c.createTable(ctx, want, c.spec.Partition, c.spec.OrderBy)
			}

			return c.autoMigrate(ctx, tableName, tableChanges.changes)
		})
	}

	return eg.Wait()
}

func (c *Client) allTablesChanges(ctx context.Context, want schema.Tables, have schema.Tables) (map[string]tableChanges, error) {
	result := make(map[string]tableChanges)
	for _, t := range want {
		chTable := have.Get(t.Name)
		if chTable == nil {
			result[t.Name] = tableChanges{
				alreadyExists:         false,
				changes:               nil,
				forcedMigrationNeeded: false,
			}
			continue
		}
		changes := t.GetChanges(chTable)
		forcedMigrationNeeded, err := c.forceMigrationNeeded(ctx, t, changes)
		if err != nil {
			return nil, err
		}
		result[t.Name] = tableChanges{
			alreadyExists:         true,
			changes:               changes,
			forcedMigrationNeeded: forcedMigrationNeeded,
		}
	}
	return result, nil
}

func (c *Client) forceMigrationNeeded(ctx context.Context, table *schema.Table, changes []schema.TableColumnChange) (bool, error) {
	if unsafe := unsafeChanges(changes); len(unsafe) > 0 {
		return true, nil
	}

	partitionKeyChange, sortingKeyChange, err := c.checkPartitionOrOrderByChanged(ctx, table, c.spec.Partition, c.spec.OrderBy)
	if err != nil {
		return false, fmt.Errorf("failed to check partition or order by changed: %w", err)
	}
	if partitionKeyChange != "" || sortingKeyChange != "" {
		return true, nil
	}

	return false, nil
}

func unsafeChanges(changes []schema.TableColumnChange) []schema.TableColumnChange {
	unsafe := make([]schema.TableColumnChange, 0, len(changes))
	for _, c := range changes {
		if needsTableDrop(c) {
			unsafe = append(unsafe, c)
		}
	}
	return slices.Clip(unsafe)
}

func (c *Client) createTable(ctx context.Context, table *schema.Table, partition []spec.PartitionStrategy, orderBy []spec.OrderByStrategy) (err error) {
	query, err := queries.CreateTable(table, c.spec.Cluster, c.spec.Engine, partition, orderBy)
	if err != nil {
		return err
	}

	if err := retryExec(ctx, c.logger, c.conn, query); err != nil {
		return fmt.Errorf("failed to create table, query:\n%s\nerror: %w", query, err)
	}
	return nil
}

func (c *Client) dropTable(ctx context.Context, table *schema.Table) error {
	c.logger.Info().Str("table", table.Name).Msg("Dropping table")

	return retryExec(ctx, c.logger, c.conn, queries.DropTable(table, c.spec.Cluster))
}

func needsTableDrop(change schema.TableColumnChange) bool {
	// Support for adding the cq_client_id column without dropping the table
	if change.Type == schema.TableColumnChangeTypeAdd && change.Current.Name == schema.CqClientIDColumn.Name {
		return false
	}

	// We can add new nullable columns or non-nullable columns that are not part of the sort key
	isCompoundType := queries.IsCompoundType(change.Current)
	if change.Type == schema.TableColumnChangeTypeAdd && (isCompoundType || !change.Current.NotNull) {
		return false
	}

	// We can safely ignore removal of nullable columns without dropping the table
	if change.Type == schema.TableColumnChangeTypeRemove && !change.Previous.NotNull {
		return false
	}

	// TODO: add check for update + new type is extending the current type (uint8 -> uint16, float32 -> float64, new struct field, etc).
	return true
}

func (c *Client) autoMigrate(ctx context.Context, tableName string, changes []schema.TableColumnChange) error {
	for _, change := range changes {
		// we only handle new columns
		if change.Type != schema.TableColumnChangeTypeAdd {
			continue
		}

		c.logger.Info().Str("table", tableName).Str("column", change.Current.Name).Msg("Adding new column")

		query, err := queries.AddColumn(tableName, c.spec.Cluster, change.Current)
		if err != nil {
			return err
		}

		err = retryExec(ctx, c.logger, c.conn, query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) checkPartitionOrOrderByChanged(ctx context.Context, table *schema.Table, partition []spec.PartitionStrategy, orderBy []spec.OrderByStrategy) (partitionKeyChange, sortingKeyChange string, err error) {
	resolvedOrderBy, err := queries.ResolveOrderBy(table, orderBy)
	if err != nil {
		return "", "", err
	}

	resolvedPartitionBy, err := queries.ResolvePartitionBy(table, partition)
	if err != nil {
		return "", "", err
	}

	splitPartitionBy := []string{}
	if resolvedPartitionBy != "" {
		splitPartitionBy = strings.Split(resolvedPartitionBy, ",")
	}

	wantPartitionKey := make([]string, 0, len(splitPartitionBy))
	for _, key := range splitPartitionBy {
		wantPartitionKey = append(wantPartitionKey, dequote(key))
	}

	wantSortingKey := make([]string, 0, len(resolvedOrderBy))
	for _, key := range resolvedOrderBy {
		wantSortingKey = append(wantSortingKey, dequote(key))
	}

	havePartitionKey, haveSortingKey, err := c.getPartitionKeyAndSortingKey(ctx, table)
	if err != nil {
		return "", "", err
	}

	partitionKeyChange = ""
	if !slices.Equal(havePartitionKey, wantPartitionKey) {
		partitionKeyChange = fmt.Sprintf("partition key changed (was [%s] and would become [%s])", strings.Join(havePartitionKey, ","), strings.Join(wantPartitionKey, ","))
		c.logger.Info().Str("table", table.Name).Msg(partitionKeyChange)
	}

	sortingKeyChange = ""
	if !slices.Equal(haveSortingKey, wantSortingKey) {
		sortingKeyChange = fmt.Sprintf("sorting key changed (was [%s] and would become [%s])", strings.Join(haveSortingKey, ","), strings.Join(wantSortingKey, ","))
		c.logger.Info().Str("table", table.Name).Msg(sortingKeyChange)
	}

	return partitionKeyChange, sortingKeyChange, nil
}

func dequote(s string) string {
	return strings.TrimSpace(strings.Trim(s, `"'`+"`"))
}
