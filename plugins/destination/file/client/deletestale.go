package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (*Client) DeleteStale(ctx context.Context, schemas schema.Schemas, sourceName string, syncTime time.Time) error {
	return fmt.Errorf("file destination doesn't support overwrite-delete-stale mode. please use append mode")
}
