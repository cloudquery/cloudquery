package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ServiceRegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	clients := make([]schema.ClientMeta, 0, len(client.services)*len(client.regions))
	for _, serviceID := range client.services {
		for _, region := range client.regions {
			clients = append(clients, client.withServiceAndRegion(serviceID, region))
		}
	}
	return clients
}
