package client

import (
	"context"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

// DeleteStale is not currently implemented for BigQuery, as it only supports "append" write mode.
func (*Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return errors.New("DeleteStale not implemented")
}
