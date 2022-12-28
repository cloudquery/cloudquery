package client

import (
	"context"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

// DeleteStale is not currently implemented for BigQuery, as it only supports "append" write mode.
func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return errors.New("DeleteStale not implemented")
}
