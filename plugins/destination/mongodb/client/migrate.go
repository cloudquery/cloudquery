package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

// func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
// 	for _, t := range tables {
// 		if err := c.migrateTable(ctx, t); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func (c *Client) CreateTables(context.Context, []*plugin.MessageCreateTable) error {
	return nil
}

// func (c *Client) migrateTable(ctx context.Context, table *schema.Table) error {
// 	tableName := table.Name
// 	for _, mdl := range c.getIndexTemplates(table) {
// 		res, err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().CreateOne(ctx, mdl)
// 		switch {
// 		case err == nil:
// 			c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("created index")
// 		case isIndexConflictError(err):
// 			c.logger.Debug().Str("index_name", res).Str("table", tableName).Err(err).Msg("create index conflict")
// 			if err := c.migrateTableOnConflict(ctx, table, mdl); err != nil {
// 				return err
// 			}
// 		case isIndexOptionsConflictError(err):
// 			c.logger.Debug().Str("table", tableName).Err(err).Msg("skipped create index")
// 		default:
// 			return fmt.Errorf("create index on %s: %w", tableName, err)
// 		}
// 	}

// 	return nil
// }

// func (c *Client) migrateTableOnConflict(ctx context.Context, table *schema.Table, mdl mongo.IndexModel) error {
// 	tableName := table.Name
// 	if c.spec.MigrateMode != specs.MigrateModeForced {
// 		return fmt.Errorf("collection %s requires forced migration due to changes in unique indexes. use 'migrate_mode: forced'", tableName)
// 	}

// 	if _, err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().DropOne(ctx, *mdl.Options.Name); err != nil {
// 		return fmt.Errorf("drop index on %s: %w", tableName, err)
// 	}
// 	res, err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().CreateOne(ctx, mdl)
// 	if err != nil {
// 		return fmt.Errorf("recreate index on %s: %w", tableName, err)
// 	}
// 	c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("recreated index")
// 	return nil
// }

// func (c *Client) getIndexTemplates(table *schema.Table) []mongo.IndexModel {
// 	var indexes []mongo.IndexModel

// 	pks := table.PrimaryKeys()
// 	if len(pks) > 0 {
// 		indexCols := bson.D{}
// 		for _, name := range pks {
// 			indexCols = append(indexCols, bson.E{Key: name, Value: 1})
// 		}

// 		pkIndexName := "cq_pk"
// 		indexes = append(indexes, mongo.IndexModel{
// 			Keys: indexCols,
// 			Options: &options.IndexOptions{
// 				Unique: &[]bool{true}[0],
// 				Name:   &pkIndexName,
// 			},
// 		})
// 	}

// 	if c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale {
// 		delIndexName := "cq_del"
// 		indexes = append(indexes, mongo.IndexModel{
// 			Keys: bson.D{{Key: schema.CqSourceNameColumn.Name, Value: 1}, {Key: schema.CqSyncTimeColumn.Name, Value: 1}},
// 			Options: &options.IndexOptions{
// 				Name: &delIndexName,
// 			},
// 		})
// 	}

// 	return indexes
// }

// func isIndexConflictError(err error) bool {
// 	cmdErr, ok := err.(mongo.CommandError)
// 	if !ok {
// 		return false
// 	}
// 	return cmdErr.Name == "IndexKeySpecsConflict"
// }

// func isIndexOptionsConflictError(err error) bool {
// 	cmdErr, ok := err.(mongo.CommandError)
// 	if !ok {
// 		return false
// 	}

// 	// This is either "Index already exists with a different name: %s" error or due to uniqueness change
// 	return cmdErr.Name == "IndexOptionsConflict"
// }
