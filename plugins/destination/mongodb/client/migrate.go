package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, t := range tables {
		if err := c.migrateTable(ctx, t); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) migrateTable(ctx context.Context, table *schema.Table) error {
	indexModels := c.getIndexTemplates(table)
	for _, mdl := range indexModels {
		res, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().CreateOne(ctx, mdl)
		if err == nil {
			c.logger.Debug().Str("index_name", res).Str("table", table.Name).Msg("created index")
		} else if isIndexConflictError(err) {
			c.logger.Debug().Str("index_name", res).Str("table", table.Name).Err(err).Msg("create index conflict")
			if err := c.migrateTableOnConflict(ctx, table, mdl); err != nil {
				return err
			}
		} else if isIndexAlreadyExistsWithADifferentNameError(err) {
			c.logger.Debug().Str("table", table.Name).Err(err).Msg("skipped create index")
		} else {
			return fmt.Errorf("create index on %s: %w", table.Name, err)
		}
	}

	for _, subTable := range table.Relations {
		if err := c.migrateTable(ctx, subTable); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) migrateTableOnConflict(ctx context.Context, table *schema.Table, mdl mongo.IndexModel) error {
	if c.spec.MigrateMode != specs.MigrateModeForced {
		return fmt.Errorf("collection %s requires forced migration due to changes in unique indexes. use 'migrate_mode: forced'", table.Name)
	}

	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().DropOne(ctx, *mdl.Options.Name); err != nil {
		return fmt.Errorf("drop index on %s: %w", table.Name, err)
	}
	res, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().CreateOne(ctx, mdl)
	if err != nil {
		return fmt.Errorf("recreate index on %s: %w", table.Name, err)
	}
	c.logger.Debug().Str("index_name", res).Str("table", table.Name).Msg("recreated index")
	return nil
}

func (c *Client) getIndexTemplates(table *schema.Table) []mongo.IndexModel {
	var indexes []mongo.IndexModel

	pks := table.PrimaryKeys()
	if len(pks) > 0 {
		indexCols := bson.D{}
		for _, col := range pks {
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

func isIndexAlreadyExistsWithADifferentNameError(err error) bool {
	cmdErr, ok := err.(mongo.CommandError)
	if !ok {
		return false
	}
	return cmdErr.Name == "IndexOptionsConflict" // Index already exists with a different name: %s
}
