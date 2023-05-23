package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return fmt.Errorf("delete-stale is not implemented")
}
