package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

type NoOpEmbeddingsClient struct {
}

func (*NoOpEmbeddingsClient) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	return nil
}

func (*NoOpEmbeddingsClient) MigrateTables(ctx context.Context) error {
	return nil
}
