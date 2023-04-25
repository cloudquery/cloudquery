package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
)

const (
	concurrentMigrations = 10
	checkTableFrequency  = 6 * time.Second
	maxTableChecks       = 20
)

// Migrate tables. It is the responsibility of the CLI of the client to lock before running migrations.
func (c *Client) Migrate(ctx context.Context, schemas schema.Schemas) error {
	eg, gctx := errgroup.WithContext(ctx)
	eg.SetLimit(concurrentMigrations)
	for _, sc := range schemas {
		sc := sc
		eg.Go(func() error {
			tableName := schema.TableName(sc)
			c.logger.Debug().Str("table", tableName).Msg("Migrating table")
			tableExists, err := c.doesTableExist(gctx, c.client, tableName)
			if err != nil {
				return fmt.Errorf("failed to check if table %s exists: %w", tableName, err)
			}
			if tableExists {
				c.logger.Debug().Str("table", tableName).Msg("Table exists, auto-migrating")
				if err := c.autoMigrateTable(gctx, c.client, sc); err != nil {
					return err
				}
				err = c.waitForSchemaToMatch(gctx, c.client, sc)
				if err != nil {
					return err
				}
			} else {
				c.logger.Debug().Str("table", tableName).Msg("Table doesn't exist, creating")
				if err := c.createTable(gctx, c.client, sc); err != nil {
					return err
				}
				err = c.waitForTableToExist(gctx, c.client, sc)
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
	c.logger.Debug().Str("dataset", c.pluginSpec.DatasetID).Str("table", table).Msg("Checking existence")
	tableRef := client.Dataset(c.pluginSpec.DatasetID).Table(table)
	md, err := tableRef.Metadata(ctx)
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == http.StatusNotFound {
				return false, nil
			}
		}
		c.logger.Error().Str("dataset", c.pluginSpec.DatasetID).Str("table", table).Err(err).Msg("Got unexpected error while checking table metadata")
		return false, err
	}

	c.logger.Debug().Interface("creation_time", md.CreationTime).Msg("Got table metadata")
	return true, nil
}

// wait until we can confirm that table now exists to avoid issues if writes are done
// immediately after the migration
func (c *Client) waitForTableToExist(ctx context.Context, client *bigquery.Client, sc *arrow.Schema) error {
	tableName := schema.TableName(sc)
	c.logger.Debug().Str("table", tableName).Msg("Waiting for table to be created")
	for i := 0; i < maxTableChecks; i++ {
		tableExists, err := c.doesTableExist(ctx, client, tableName)
		if err != nil {
			return err
		}
		if tableExists {
			c.logger.Debug().Str("table", tableName).Msg("Table created")
			return nil
		}
		c.logger.Debug().Str("table", tableName).Int("i", i).Msg("Waiting for table to be created")
		time.Sleep(checkTableFrequency)
	}
	return fmt.Errorf("failed to confirm table creation for %v within timeout period", tableName)
}

// wait until we can confirm that schema now matches, to avoid issues if writes are done
// immediately after the migration
func (c *Client) waitForSchemaToMatch(ctx context.Context, client *bigquery.Client, sc *arrow.Schema) error {
	tableName := schema.TableName(sc)
	c.logger.Debug().Str("table", tableName).Msg("Waiting for schemas to match")
	wantSchema := c.bigQuerySchemaForTable(sc)
	for i := 0; i < maxTableChecks; i++ {
		// require this check to pass 3 times in a row to mitigate getting different responses from different BQ servers
		tries := 3
		for j := 0; j < tries; j++ {
			md, err := client.Dataset(c.pluginSpec.DatasetID).Table(tableName).Metadata(ctx)
			if err != nil {
				return err
			}
			haveSchema := md.Schema
			if !schemasMatch(haveSchema, wantSchema) {
				continue
			}
			if j == tries-1 {
				c.logger.Debug().Str("table", tableName).Msg("Schemas match")
				return nil
			}
		}
		c.logger.Debug().Str("table", tableName).Int("i", i).Msg("Waiting for schemas to match")
		time.Sleep(checkTableFrequency)
	}
	return fmt.Errorf("failed to confirm schema update for %v within timeout period", tableName)
}

func (c *Client) autoMigrateTable(ctx context.Context, client *bigquery.Client, sc *arrow.Schema) error {
	tableName := schema.TableName(sc)
	tableDescription := schema.TableDescription(sc)
	bqTable := client.Dataset(c.pluginSpec.DatasetID).Table(tableName)
	md, err := bqTable.Metadata(ctx)
	if err != nil {
		return fmt.Errorf("failed to get metadata for table %q with error: %w", tableName, err)
	}
	haveSchema := md.Schema
	wantSchema := c.bigQuerySchemaForTable(sc)
	wantSchema, err = mergeSchemas(haveSchema, wantSchema)
	if err != nil {
		return fmt.Errorf("failed to migrate schema for table %q with error: %w", tableName, err)
	}
	tm := bigquery.TableMetadataToUpdate{
		Name:        tableName,
		Description: tableDescription,
		Schema:      wantSchema,
	}
	_, err = bqTable.Update(ctx, tm, "")
	if err != nil {
		return fmt.Errorf("failed to update schema for table %q with error: %w", tableName, err)
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

func (c *Client) createTable(ctx context.Context, client *bigquery.Client, sc *arrow.Schema) error {
	bqSchema := c.bigQuerySchemaForTable(sc)
	tableName := schema.TableName(sc)
	tableDescription := schema.TableDescription(sc)
	tm := bigquery.TableMetadata{
		Name:             tableName,
		Location:         "",
		Description:      tableDescription,
		Schema:           bqSchema,
		TimePartitioning: c.timePartitioning(),
	}
	return client.Dataset(c.pluginSpec.DatasetID).Table(tableName).Create(ctx, &tm)
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

func (c *Client) bigQuerySchemaForTable(sc *arrow.Schema) bigquery.Schema {
	s := bigquery.Schema{}
	for _, field := range sc.Fields() {
		columnType, repeated := c.SchemaTypeToBigQueryType(field.Type)
		s = append(s, &bigquery.FieldSchema{
			Name:     field.Name,
			Repeated: repeated,
			Type:     columnType,
		})
	}
	return s
}
