package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (*Client) DeleteStale(context.Context, schema.Schemas, string, time.Time) error {
	return fmt.Errorf("not implemented. plugin supports only 'append' mode")
}
