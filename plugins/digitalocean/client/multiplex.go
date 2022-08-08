package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func SpacesRegionMultiplex(client schema.ClientMeta) []schema.ClientMeta {
	doClient := client.(*Client)
	clients := make([]schema.ClientMeta, len(doClient.Regions))
	for i, r := range doClient.Regions {
		clients[i] = doClient.WithSpacesRegion(r)
	}
	return clients
}
