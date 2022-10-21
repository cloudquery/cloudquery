package client

import (
	"github.com/cloudquery/plugin-sdk/plugins"
)

func (c *Client) Stats() plugins.DestinationStats {
		return c.stats
}