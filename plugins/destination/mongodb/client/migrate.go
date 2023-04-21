package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	for _, t := range tables {
		if err := c.migrateTable(ctx, t); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) migrateTable(ctx context.Context, table *arrow.Schema) error {
	tableName := schema.TableName(table)
	for _, mdl := range c.getIndexTemplates(table) {
		res, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).Indexes().CreateOne(ctx, mdl)
		switch {
		case err == nil:
			c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("created index")
		case isIndexConflictError(err):
			c.logger.Debug().Str("index_name", res).Str("table", tableName).Err(err).Msg("create index conflict")
			if err := c.migrateTableOnConflict(ctx, table, mdl); err != nil {
				return err
			}
		case isIndexOptionsConflictError(err):
			c.logger.Debug().Str("table", tableName).Err(err).Msg("skipped create index")
		default:
			return fmt.Errorf("create index on %s: %w", tableName, err)
		}
	}

	return nil
}

func (c *Client) migrateTableOnConflict(ctx context.Context, table *arrow.Schema, mdl mongo.IndexModel) error {
	tableName := schema.TableName(table)
	if c.spec.MigrateMode != specs.MigrateModeForced {
		return fmt.Errorf("collection %s requires forced migration due to changes in unique indexes. use 'migrate_mode: forced'", tableName)
	}

	if _, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).Indexes().DropOne(ctx, *mdl.Options.Name); err != nil {
		return fmt.Errorf("drop index on %s: %w", tableName, err)
	}
	res, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).Indexes().CreateOne(ctx, mdl)
	if err != nil {
		return fmt.Errorf("recreate index on %s: %w", tableName, err)
	}
	c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("recreated index")
	return nil
}

func (c *Client) getIndexTemplates(table *arrow.Schema) []mongo.IndexModel {
	var indexes []mongo.IndexModel

	pks := schema.PrimaryKeyIndices(table)
	if len(pks) > 0 {
		indexCols := bson.D{}
		for _, colIndex := range pks {
			col := table.Field(colIndex).Name
			indexCols = append(indexCols, bson.E{Key: col, Value: 1})
		}

		pkIndexName := "cq_pk"
		indexes = append(indexes, mongo.IndexModel{
			Keys: indexCols,
			Options: &options.IndexOptions{
				Unique: &[]bool{true}[0],
				Name:   &pkIndexName,
			},
		})
	}

	if c.spec.WriteMode == specs.WriteModeOverwriteDeleteStale {
		delIndexName := "cq_del"
		indexes = append(indexes, mongo.IndexModel{
			Keys: bson.D{{Key: schema.CqSourceNameColumn.Name, Value: 1}, {Key: schema.CqSyncTimeColumn.Name, Value: 1}},
			Options: &options.IndexOptions{
				Name: &delIndexName,
			},
		})
	}

	return indexes
}

func isIndexConflictError(err error) bool {
	cmdErr, ok := err.(mongo.CommandError)
	if !ok {
		return false
	}
	return cmdErr.Name == "IndexKeySpecsConflict"
}

func isIndexOptionsConflictError(err error) bool {
	cmdErr, ok := err.(mongo.CommandError)
	if !ok {
		return false
	}

	// This is either "Index already exists with a different name: %s" error or due to uniqueness change
	return cmdErr.Name == "IndexOptionsConflict"
}
