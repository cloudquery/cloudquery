package client

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

func (c *Client) Sync(ctx context.Context, metrics *source.Metrics, res chan<- *schema.Resource) error {
	c.metrics = metrics
	for _, table := range c.Tables {
		if c.metrics.TableClient[table.Name] == nil {
			c.metrics.TableClient[table.Name] = make(map[string]*source.TableClientMetrics)
			c.metrics.TableClient[table.Name][c.ID()] = &source.TableClientMetrics{}
		}
	}
	return c.syncTables(ctx, res)
}

func (c *Client) syncTable(ctx context.Context, table *schema.Table, res chan<- *schema.Resource) error {
	colMap := make(map[string]struct{}, len(table.Columns))
	for _, col := range table.Columns {
		colMap[col.Name] = struct{}{}
	}
	var err error
	lastDocumentId := ""
	maxBatchSize := c.maxBatchSize
	collection := c.client.Collection(table.Name)
	for {
		orderByField := firestore.DocumentID
		if c.orderByField != "" {
			orderByField = c.orderByField
		}

		dir := firestore.Asc
		if c.orderByDirection == "desc" {
			dir = firestore.Desc
		}

		query := collection.Query.
			OrderBy(orderByField, dir).
			Limit(maxBatchSize)
		if lastDocumentId != "" {
			c.logger.Info().Msgf("Starting after %s", lastDocumentId)
			query = query.StartAfter(lastDocumentId)
		}
		docIter := query.Documents(ctx)
		var documentCount int
		var skippedCount int
		for {
			docSnap, err := docIter.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}
				return err
			}
			documentCount++
			if !docSnap.Exists() {
				skippedCount++
				continue
			}
			lastDocumentId = docSnap.Ref.ID
			item := docSnap.Data()
			resource := schema.NewResourceData(table, nil, item)
			resource.Set("_id", docSnap.Ref.ID)
			resource.Set("_created_at", docSnap.CreateTime)
			resource.Set("_updated_at", docSnap.UpdateTime)
			resource.Set("data", item)
			c.metrics.TableClient[table.Name][c.ID()].Resources++
			res <- resource
		}
		c.logger.Info().Msgf("Synced %d documents from %s", documentCount, table.Name)
		if skippedCount > 0 {
			c.logger.Info().Msgf("Skipped %d documents from %s", skippedCount, table.Name)
		}
		if documentCount < maxBatchSize {
			break
		}
	}
	return err
}

func updateColumn(item map[string]interface{}, columns schema.ColumnList) schema.ColumnList {
	var newCols schema.ColumnList
	colsMap := make(map[string]struct{})
	for _, col := range columns {
		colsMap[col.Name] = struct{}{}
	}
	for key, value := range item {
		if value == nil {
			continue
		}
		_, found := colsMap[key]
		if !found {
			schemaType := determineSchemaType(value)
			if schemaType == schema.TypeInvalid {
				continue
			}
			newCols = append(newCols, schema.Column{
				Name: key,
				Type: schemaType,
			})
		}
	}
	return newCols
}

func determineSchemaType(value interface{}) schema.ValueType {
	switch value.(type) {
	case bool:
		return schema.TypeBool
	case time.Time:
		return schema.TypeTimestamp
	case []byte:
		return schema.TypeByteArray
	case []string:
		return schema.TypeStringArray
	case float64, float32:
		return schema.TypeFloat
	case int8, int16, int32, int64:
		return schema.TypeInt
	case []int8, []int16, []int32, []int64:
		return schema.TypeIntArray
	case map[string]interface{}, []interface{}:
		return schema.TypeJSON
	case string:
		return schema.TypeString
	default:
		return schema.TypeInvalid
	}
}

func (c *Client) syncTables(ctx context.Context, res chan<- *schema.Resource) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.SetLimit(len(c.Tables))

	for _, table := range c.Tables {
		t := table
		eg.Go(func() error {
			if err := c.syncTable(ctx, t, res); err != nil {
				return err
			}
			return nil
		})
	}
	return eg.Wait()
}
