package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

// DeleteStale is not currently implemented for BigQuery, as it only supports "append" write mode.
func (*Client) DeleteStale(context.Context, message.WriteDeleteStales) error {
	return errors.New("DeleteStale not implemented")
}
