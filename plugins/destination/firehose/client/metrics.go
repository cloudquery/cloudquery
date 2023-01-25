package client

import (
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func (c *Client) Metrics() destination.Metrics {
	return c.metrics
}
