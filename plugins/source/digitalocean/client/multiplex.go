package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func SpacesRegionMultiplex(client schema.ClientMeta) []schema.ClientMeta {
	doClient := client.(*Client)
	if len(doClient.Regions) == 0 {
		return []schema.ClientMeta{doClient}
	}
	clients := make([]schema.ClientMeta, len(doClient.Regions))
	for i, r := range doClient.Regions {
		clients[i] = doClient.WithSpacesRegion(r)
	}
	return clients
}
