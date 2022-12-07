package client

import "github.com/cloudquery/plugin-sdk/schema"

func TeamMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, teamID := range client.AllTeamIDs {
		l = append(l, client.withTeamID(teamID))
	}
	return l
}
