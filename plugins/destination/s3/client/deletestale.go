package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (*Client) DeleteStale(ctx context.Context, schemas schema.Tables, sourceName string, syncTime time.Time) error {
	return fmt.Errorf("destination plugin doesn't support overwrite-delete-stale mode. please use append mode")
}
