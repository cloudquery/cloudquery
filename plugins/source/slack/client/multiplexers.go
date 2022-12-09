package client

import "github.com/cloudquery/plugin-sdk/schema"

func TeamMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, team := range client.Teams {
		l = append(l, client.withTeamID(team.ID))
	}
	return l
}
