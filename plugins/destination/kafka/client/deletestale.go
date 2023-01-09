package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return fmt.Errorf("not implemented. plugin supports only 'append' mode")
}
