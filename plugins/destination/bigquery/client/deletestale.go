package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

// DeleteStale is not currently implemented for BigQuery, as it only supports "append" write mode.
func (*Client) DeleteStale(context.Context, message.WriteDeleteStales) error {
	return plugin.ErrNotImplemented
}
