package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func TeamMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	c := meta.(*Client)

	l := make([]schema.ClientMeta, 0, len(c.TeamIDs))
	for _, o := range c.TeamIDs {
		l = append(l, c.WithTeamID(o))
	}

	return l
}
