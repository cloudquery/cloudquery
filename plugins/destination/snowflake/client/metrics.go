package client

import (
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
)

func (c *Client) Metrics() destination.Metrics {
	return c.metrics
}
