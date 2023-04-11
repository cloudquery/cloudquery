package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/meilisearch/meilisearch-go"
	"golang.org/x/exp/slices"
)

type indexSchema struct {
	UID        string
	PrimaryKey string
	Attributes []string
}

func (i *indexSchema) init(index *meilisearch.Index) (*indexSchema, error) {
	i.UID = index.UID
	i.PrimaryKey = index.PrimaryKey

	attrs, err := index.GetFilterableAttributes()
	if err != nil {
		return nil, err
	}
	if attrs != nil {
		i.Attributes = *attrs
	}

	return i, nil
}

func (i *indexSchema) canMigrate(o *indexSchema) bool {
	return i.UID == o.UID && i.PrimaryKey == o.PrimaryKey
}

func (c *Client) tableIndexSchema(table *schema.Table) *indexSchema {
	return &indexSchema{
		UID:        table.Name,
		PrimaryKey: c.pkColumn,
		Attributes: table.Columns.Names(),
	}
}

func (c *Client) tablesIndexSchemas(tables schema.Tables) map[string]*indexSchema {
	res := make(map[string]*indexSchema)
	for _, table := range tables.FlattenTables() {
		s := c.tableIndexSchema(table)
		res[s.UID] = s
	}
	return res
}

func (c *Client) indexes() (map[string]*indexSchema, error) {
	req := &meilisearch.IndexesQuery{Limit: 100, Offset: 0}

	result := make(map[string]*indexSchema)

	for {
		resp, err := c.Meilisearch.GetIndexes(req)
		if err != nil {
			return nil, err
		}

		for _, index := range resp.Results {
			entry, err := new(indexSchema).init(&index)
			if err != nil {
				return nil, err
			}
			result[entry.UID] = entry
		}

		if resp.Offset+resp.Limit >= resp.Total {
			break
		}
		req.Offset += resp.Limit
	}

	return result, nil
}

func (c *Client) configureIndex(ctx context.Context, s *indexSchema) error {
	c.logger.Debug().Str("index", s.UID).Msg("configuring index")

	index, err := c.Meilisearch.GetIndex(s.UID)
	if err != nil {
		return err
	}

	attributes := slices.Clone(s.Attributes)

	attrs, err := index.GetFilterableAttributes()
	if err != nil {
		return err
	}

	var current []string
	if attrs != nil {
		current = *attrs
	}

	if len(current) > 0 {
		attributes = append(attributes, current...)
		slices.Sort(attributes)
		attributes = slices.Compact(attributes)
		if len(attributes) == len(s.Attributes) {
			// already the same filtered attributes, skip
			c.logger.Info().Str("index", index.UID).Msg("index is already properly configured, skip")
			return nil
		}
	}

	taskInfo, err := index.UpdateFilterableAttributes(&attributes)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to update filterable attributes for index %q: %w", index.UID, err)
	}

	taskInfo, err = index.UpdateSortableAttributes(&attributes)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to update sortable attributes for index %q: %w", index.UID, err)
	}

	return nil
}

func (c *Client) createIndex(ctx context.Context, s *indexSchema) error {
	c.logger.Debug().Str("index", s.UID).Msg("creating index")

	taskInfo, err := c.Meilisearch.CreateIndex(&meilisearch.IndexConfig{
		Uid:        s.UID,
		PrimaryKey: s.PrimaryKey,
	})
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to create index %q: %w", s.UID, err)
	}

	return c.configureIndex(ctx, s)
}

func (c *Client) recreateIndex(ctx context.Context, s *indexSchema) error {
	if err := c.deleteIndex(ctx, s); err != nil {
		return err
	}
	return c.createIndex(ctx, s)
}

func (c *Client) deleteIndex(ctx context.Context, s *indexSchema) error {
	c.logger.Debug().Str("index", s.UID).Msg("deleting index")

	taskInfo, err := c.Meilisearch.DeleteIndex(s.UID)
	if err != nil {
		return err
	}

	if err := c.waitTask(ctx, taskInfo); err != nil {
		return fmt.Errorf("failed to delete index %q: %w", s.UID, err)
	}

	return nil
}
