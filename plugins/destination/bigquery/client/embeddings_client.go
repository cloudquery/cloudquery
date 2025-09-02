package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

type EmbeddingsClient interface {
	MigrateTables(ctx context.Context) error
	WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error
}

func NewEmbeddingsClient(cl *Client, spec *Spec) EmbeddingsClient {
	if spec.TextEmbeddings == nil {
		return &NoOpEmbeddingsClient{}
	}
	c := &ConcreteEmbeddingsClient{
		client:    cl,
		spec:      spec.TextEmbeddings,
		ProjectID: spec.ProjectID,
		DatasetID: spec.DatasetID,
	}
	return c
}

type ConcreteEmbeddingsClient struct {
	client *Client
	spec   *TextEmbeddingsSpec

	hasMigrated bool

	// The id of the project where the destination BigQuery database resides.
	ProjectID string `json:"project_id" jsonschema:"required,minLength=1"`

	//  The name of the BigQuery dataset within the project, e.g. `my_dataset`.
	//  This dataset needs to be created before running a sync or migration.
	DatasetID string `json:"dataset_id" jsonschema:"required,minLength=1"`
}

// findTableConfig finds the table configuration by source table name
func (c *ConcreteEmbeddingsClient) findTableConfig(name string) (*TableConfig, error) {
	for _, table := range c.spec.Tables {
		if table.SourceTableName == name {
			return &table, nil
		}
	}
	return nil, fmt.Errorf("table configuration not found for source table: %s", name)
}
