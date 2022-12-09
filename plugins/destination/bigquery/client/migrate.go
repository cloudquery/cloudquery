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
	client, err := c.bqClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	eg, gctx := errgroup.WithContext(ctx)
	eg.SetLimit(concurrentMigrations)
	for _, table := range tables.FlattenTables() {
		table := table
		eg.Go(func() error {
			c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
			tableExists, err := c.doesTableExist(gctx, client, table.Name)
			if err != nil {
				return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
			}
			if tableExists {
				c.logger.Debug().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(gctx, client, table); err != nil {
					return err
				}
			} else {
				c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
				if err := c.createTable(gctx, client, table); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return eg.Wait()
}

func (c *Client) doesTableExist(ctx context.Context, client *bigquery.Client, table string) (bool, error) {
	c.logger.Debug().Str("dataset", c.pluginSpec.DatasetID).Str("table", table).Msg("Checking existence")
	tableRef := client.Dataset(c.pluginSpec.DatasetID).Table(table)
	md, err := tableRef.Metadata(ctx)
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				return false, nil
			}
		}
		c.logger.Error().Err(err).Msg("Got unexpected error while checking table metadata")
		return false, err
	}

	c.logger.Debug().Interface("creation_time", md.CreationTime).Msg("Got table metadata")
	return true, nil
}

func (c *Client) autoMigrateTable(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	bqSchema := c.bigQuerySchemaForTable(table)
	tm := bigquery.TableMetadataToUpdate{
		Name:        table.Name,
		Description: table.Description,
		Schema:      bqSchema,
	}
	_, err := client.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Update(ctx, tm, "")
	return err
}

func (c *Client) createTable(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	bqSchema := c.bigQuerySchemaForTable(table)
	tm := bigquery.TableMetadata{
		Name:             table.Name,
		Location:         "",
		Description:      table.Description,
		Schema:           bqSchema,
		TimePartitioning: c.timePartitioning(),
	}
	return client.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Create(ctx, &tm)
}

func (c *Client) timePartitioning() *bigquery.TimePartitioning {
	switch c.pluginSpec.TimePartitioning {
	case TimePartitioningOptionHour:
		return &bigquery.TimePartitioning{
			Type:  "HOUR",
			Field: "_cq_sync_time",
		}
	case TimePartitioningOptionDay:
		return &bigquery.TimePartitioning{
			Type:  "DAY",
			Field: "_cq_sync_time",
		}
	default:
		return nil
	}
}

func (c *Client) bigQuerySchemaForTable(table *schema.Table) bigquery.Schema {
	s := bigquery.Schema{}
	for _, col := range table.Columns {
		columnType, repeated := c.SchemaTypeToBigQuery(col.Type)
		s = append(s, &bigquery.FieldSchema{
			Name:        col.Name,
			Description: col.Description,
			Repeated:    repeated,
			Type:        columnType,
		})
	}
	return s
}
