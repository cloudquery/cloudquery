package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	for _, msg := range msgs {
		if err := c.migrateTable(ctx, msg.MigrateForce, msg.Table); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) migrateTable(ctx context.Context, force bool, table *schema.Table) error {
	tableName := table.Name
	for _, tmpl := range c.getIndexTemplates(table) {
		res, err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().CreateOne(ctx, tmpl.model)
		switch {
		case err == nil:
			c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("created index")
		case isIndexConflictError(err):
			c.logger.Debug().Str("index_name", res).Str("table", tableName).Err(err).Msg("create index conflict")
			if err := c.migrateTableOnConflict(ctx, force, table, tmpl.model, tmpl.name); err != nil {
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

func (c *Client) migrateTableOnConflict(ctx context.Context, force bool, table *schema.Table, mdl mongo.IndexModel, indexName string) error {
	tableName := table.Name
	if !force {
		return fmt.Errorf("collection %s requires forced migration due to changes in unique indexes. Migrate manually or consider using 'migrate_mode: forced'", tableName)
	}

	if err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().DropOne(ctx, indexName); err != nil {
		return fmt.Errorf("drop index on %s: %w", tableName, err)
	}
	res, err := c.client.Database(c.spec.Database).Collection(tableName).Indexes().CreateOne(ctx, mdl)
	if err != nil {
		return fmt.Errorf("recreate index on %s: %w", tableName, err)
	}
	c.logger.Debug().Str("index_name", res).Str("table", tableName).Msg("recreated index")
	return nil
}

type indexTemplate struct {
	model mongo.IndexModel
	name  string
}

func (*Client) getIndexTemplates(table *schema.Table) []indexTemplate {
	var indexes []indexTemplate

	pks := table.PrimaryKeys()
	if len(pks) > 0 {
		indexCols := bson.D{}
		for _, name := range pks {
			indexCols = append(indexCols, bson.E{Key: name, Value: 1})
		}

		pkIndexName := "cq_pk"
		indexes = append(indexes, indexTemplate{
			name: pkIndexName,
			model: mongo.IndexModel{
				Keys:    indexCols,
				Options: options.Index().SetUnique(true).SetName(pkIndexName),
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
