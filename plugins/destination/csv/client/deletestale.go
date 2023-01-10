package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return fmt.Errorf("csv destination doesn't support overwrite-delete-stale mode. please use append mode")
}
