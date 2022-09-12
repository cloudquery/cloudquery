package client

import "github.com/cloudquery/plugin-sdk/schema"

func SpacesRegionMultiplex(client schema.ClientMeta) []schema.ClientMeta {
	doClient := client.(*Client)
	clients := make([]schema.ClientMeta, len(doClient.Regions))
	for i, r := range doClient.Regions {
		clients[i] = doClient.WithSpacesRegion(r)
	}
	return clients
}
