package client

import "github.com/cloudquery/plugin-sdk/schema"

func TeamMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	clients := make([]schema.ClientMeta, len(client.Teams))
	for i, team := range client.Teams {
		clients[i] = client.withTeamID(team.ID)
	}
	return clients
}
