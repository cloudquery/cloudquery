package client

import (
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
)

func (c *Client) Metrics() destination.Metrics {
	return c.metrics
}
