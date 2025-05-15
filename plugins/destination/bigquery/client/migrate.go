package client

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/sync/errgroup"
)

const (
	concurrentMigrations = 10
	checkTableFrequency  = 6 * time.Second
	maxTableChecks       = 20
	maxDescriptionLength = 16384
)

func trimDescription(description string) string {
	if len(description) > maxDescriptionLength {
		return description[:maxDescriptionLength]
	}
	return description
}

func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	eg, gctx := errgroup.WithContext(ctx)
	eg.SetLimit(concurrentMigrations)
	for _, msg := range msgs {
		table := msg.Table
		eg.Go(func() error {
			c.logger.Debug().Str("table", table.Name).Msg("Migrating table")
			tableExists, err := c.doesTableExist(gctx, c.client, table.Name)
			if err != nil {
				return fmt.Errorf("failed to check if table %s exists: %w", table.Name, err)
			}
			if tableExists {
				c.logger.Debug().Str("table", table.Name).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(gctx, c.client, table); err != nil {
					return err
				}
				err = c.waitForSchemaToMatch(gctx, c.client, table)
				if err != nil {
					return err
				}
			} else {
				c.logger.Debug().Str("table", table.Name).Msg("Table doesn't exist, creating")
				if err := c.createTable(gctx, c.client, table); err != nil {
					return err
				}
				err = c.waitForTableToExist(gctx, c.client, table)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	return eg.Wait()
}

func (c *Client) doesTableExist(ctx context.Context, client *bigquery.Client, table string) (bool, error) {
	c.logger.Debug().Str("dataset", c.spec.DatasetID).Str("table", table).Msg("Checking existence")
	tableRef := client.DatasetInProject(c.spec.ProjectID, c.spec.DatasetID).Table(table)
	md, err := tableRef.Metadata(ctx)
	if err != nil {
		if isAPINotFoundError(err) {
			return false, nil
		}
		c.logger.Error().Str("dataset", c.spec.DatasetID).Str("table", table).Err(err).Msg("Got unexpected error while checking table metadata")
		return false, err
	}

	c.logger.Debug().Interface("creation_time", md.CreationTime).Msg("Got table metadata")
	return true, nil
}

// wait until we can confirm that table now exists to avoid issues if writes are done
// immediately after the migration
func (c *Client) waitForTableToExist(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	c.logger.Debug().Str("table", table.Name).Msg("Waiting for table to be created")
	for i := 0; i < maxTableChecks; i++ {
		tableExists, err := c.doesTableExist(ctx, client, table.Name)
		if err != nil {
			return err
		}
		if tableExists {
			c.logger.Debug().Str("table", table.Name).Msg("Table created")
			return nil
		}
		c.logger.Debug().Str("table", table.Name).Int("i", i).Msg("Waiting for table to be created")
		time.Sleep(checkTableFrequency)
	}
	return fmt.Errorf("failed to confirm table creation for %v within timeout period", table.Name)
}

// wait until we can confirm that schema now matches, to avoid issues if writes are done
// immediately after the migration
func (c *Client) waitForSchemaToMatch(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	c.logger.Debug().Str("table", table.Name).Msg("Waiting for schemas to match")
	wantSchema := c.bigQuerySchemaForTable(table)
	for i := 0; i < maxTableChecks; i++ {
		// require this check to pass 3 times in a row to mitigate getting different responses from different BQ servers
		tries := 3
		for j := 0; j < tries; j++ {
			md, err := client.DatasetInProject(c.spec.ProjectID, c.spec.DatasetID).Table(table.Name).Metadata(ctx)
			if err != nil {
				return err
			}
			haveSchema := md.Schema
			if !schemasMatch(haveSchema, wantSchema) {
				continue
			}
			if j == tries-1 {
				c.logger.Debug().Str("table", table.Name).Msg("Schemas match")
				return nil
			}
		}
		c.logger.Debug().Str("table", table.Name).Int("i", i).Msg("Waiting for schemas to match")
		time.Sleep(checkTableFrequency)
	}
	return fmt.Errorf("failed to confirm schema update for %v within timeout period", table.Name)
}

func (c *Client) autoMigrateTable(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	bqTable := client.DatasetInProject(c.spec.ProjectID, c.spec.DatasetID).Table(table.Name)
	md, err := bqTable.Metadata(ctx)
	if err != nil {
		return fmt.Errorf("failed to get metadata for table %q with error: %w", table.Name, err)
	}
	haveSchema := md.Schema
	wantSchema := c.bigQuerySchemaForTable(table)
	wantSchema, err = mergeSchemas(haveSchema, wantSchema)
	if err != nil {
		return fmt.Errorf("failed to migrate schema for table %q with error: %w", table.Name, err)
	}
	tm := bigquery.TableMetadataToUpdate{
		Name:        table.Name,
		Description: trimDescription(table.Description),
		Schema:      wantSchema,
	}
	_, err = bqTable.Update(ctx, tm, "")
	if err != nil {
		return fmt.Errorf("failed to update schema for table %q with error: %w", table.Name, err)
	}
	return nil
}

func schemasMatch(haveSchema, wantSchema bigquery.Schema) bool {
	// Schemas are considered a match if everything in the want schema is in the have schema,
	// and they have the same types.
	// We don't mind if there are extra fields in the have schema.
	haveMap := make(map[string]*bigquery.FieldSchema)
	for _, f := range haveSchema {
		haveMap[f.Name] = f
	}
	for _, wf := range wantSchema {
		if hf, ok := haveMap[wf.Name]; !ok {
			return false
		} else if hf.Type != wf.Type {
			return false
		}
	}
	return true
}

// mergeSchemas merges the schema we want with the schema we have, to avoid
// losing any existing data
func mergeSchemas(haveSchema, wantSchema bigquery.Schema) (bigquery.Schema, error) {
	haveMap := make(map[string]*bigquery.FieldSchema)
	for _, f := range haveSchema {
		haveMap[f.Name] = f
	}
	wantMap := make(map[string]*bigquery.FieldSchema)
	for _, f := range wantSchema {
		wantMap[f.Name] = f
	}
	merged := make(bigquery.Schema, 0, len(wantSchema))
	// keep everything in the schema we have, as long as the types didn't change
	// or an unknown column isn't required
	for _, f := range haveSchema {
		if want, ok := wantMap[f.Name]; ok {
			if want.Type != f.Type {
				return nil, fmt.Errorf("column %v changed type from %v to %v. Try dropping the column and re-running", f.Name, f.Type, want.Type)
			}
		} else if f.Required {
			return nil, fmt.Errorf("column %v is required but not in new schema", f.Name)
		}
		merged = append(merged, f)
	}
	// add anything new
	for _, f := range wantSchema {
		if _, ok := haveMap[f.Name]; !ok {
			merged = append(merged, f)
		}
	}
	return merged, nil
}

func (c *Client) createTable(ctx context.Context, client *bigquery.Client, table *schema.Table) error {
	bqSchema := c.bigQuerySchemaForTable(table)
	tm := bigquery.TableMetadata{
		Name:             table.Name,
		Location:         "",
		Description:      trimDescription(table.Description),
		Schema:           bqSchema,
		TimePartitioning: c.timePartitioning(),
	}
	return client.DatasetInProject(c.spec.ProjectID, c.spec.DatasetID).Table(table.Name).Create(ctx, &tm)
}

func (c *Client) timePartitioning() *bigquery.TimePartitioning {
	switch c.spec.TimePartitioning {
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
		s = append(s, c.ColumnToBigQuerySchema(col))
	}
	return s
}
