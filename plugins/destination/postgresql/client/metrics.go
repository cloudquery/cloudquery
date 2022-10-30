package client

import (
	"github.com/cloudquery/plugin-sdk/plugins"
)

func (c *Client) Metrics() plugins.DestinationMetrics {
	return c.metrics
}
