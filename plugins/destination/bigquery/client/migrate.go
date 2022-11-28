package client

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
)

const concurrentMigrations = 10

// Migrate tables. It is the responsibility of the CLI of the client to lock before running migrations.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	eg := errgroup.Group{}
	eg.SetLimit(concurrentMigrations)
	for _, table := range tables {
		table := table
		eg.Go(func() error {
			c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
			tableExists, err := c.doesTableExist(ctx, table.Name)
			if err != nil {
				return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
			}
			if tableExists {
				c.logger.Debug().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(ctx, table); err != nil {
					return err
				}
			} else {
				c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
				if err := c.createTable(ctx, table); err != nil {
					return err
				}
			}
			if err := c.Migrate(ctx, table.Relations); err != nil {
				return err
			}
			return nil
		})
	}
	eg.Wait()
	return nil
}

func (c *Client) doesTableExist(ctx context.Context, table string) (bool, error) {
	c.logger.Debug().Str("dataset", c.datasetID).Str("table", table).Msg("Checking existence")
	tableRef := c.client.Dataset(c.datasetID).Table(table)
	if md, err := tableRef.Metadata(ctx); err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				return false, nil
			}
		}
		c.logger.Error().Err(err).Msg("Got unexpected error while checking table metadata")
		return false, err
	} else {
		c.logger.Debug().Interface("creation_time", md.CreationTime).Msg("Got table metadata")
	}
	return true, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, table *schema.Table) error {
	bqSchema, err := c.bigQuerySchemaForTable(table)
	if err != nil {
		return fmt.Errorf("failed to create BigQuery schema for table: %w", err)
	}
	tm := bigquery.TableMetadataToUpdate{
		Name:        table.Name,
		Description: table.Description,
		Schema:      bqSchema,
	}
	_, err = c.client.Dataset(c.datasetID).Table(table.Name).Update(ctx, tm, "")
	return err
}

func (c *Client) createTable(ctx context.Context, table *schema.Table) error {
	bqSchema, err := c.bigQuerySchemaForTable(table)
	if err != nil {
		return fmt.Errorf("failed to create BigQuery schema for table: %w", err)
	}
	tm := bigquery.TableMetadata{
		Name:        table.Name,
		Description: table.Description,
		Schema:      bqSchema,
	}
	return c.client.Dataset(c.datasetID).Table(table.Name).Create(ctx, &tm)
}

func (c *Client) bigQuerySchemaForTable(table *schema.Table) (bigquery.Schema, error) {
	s := bigquery.Schema{}
	for _, col := range table.Columns {
		columnType, repeated := c.SchemaTypeToBigQuery(col.Type)
		s = append(s, &bigquery.FieldSchema{
			Name:        col.Name,
			Description: col.Description,
			Repeated:    repeated,
			Type:        columnType,
			Schema:      bigquery.Schema{},
		})
	}
	return s, nil
}
