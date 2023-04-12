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
	indexModel := getIndexTemplate(table)
	if indexModel != nil {
		res, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().CreateOne(ctx, *indexModel)
		if err != nil && isIndexConflictError(err) {
			c.logger.Debug().Str("index_name", res).Str("table", table.Name).Err(err).Msg("create index conflict")
			if err := c.migrateTableOnConflict(ctx, table, indexModel); err != nil {
				return err
			}
		} else if err != nil {
			return fmt.Errorf("create index on %s: %w", table.Name, err)
		}

		c.logger.Debug().Str("index_name", res).Str("table", table.Name).Msg("created index")
	}

	for _, subTable := range table.Relations {
		if err := c.migrateTable(ctx, subTable); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) migrateTableOnConflict(ctx context.Context, table *schema.Table, indexModel *mongo.IndexModel) error {
	if c.spec.MigrateMode != specs.MigrateModeForced {
		return fmt.Errorf("collection %s requires forced migration due to changes in unique indexes. use 'migrate_mode: forced'", table.Name)
	}

	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().DropOne(ctx, *indexModel.Options.Name); err != nil {
		return fmt.Errorf("drop index on %s: %w", table.Name, err)
	}
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).Indexes().CreateOne(ctx, *indexModel); err != nil {
		return fmt.Errorf("recreate index on %s: %w", table.Name, err)
	}
	return nil
}

func getIndexTemplate(table *schema.Table) *mongo.IndexModel {
	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		return nil
	}

	indexCols := bson.D{}
	for _, col := range pks {
		indexCols = append(indexCols, bson.E{Key: col, Value: 1})
	}
	indexName := "cq_pk"

	return &mongo.IndexModel{
		Keys: indexCols,
		Options: &options.IndexOptions{
			Unique: &[]bool{true}[0],
			Name:   &indexName,
		},
	}
}

func isIndexConflictError(err error) bool {
	cmdErr, ok := err.(mongo.CommandError)
	if !ok {
		return false
	}
	return cmdErr.Name == "IndexKeySpecsConflict"
}
