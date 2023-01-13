package client

import "github.com/cloudquery/plugin-sdk/schema"

func AccountRegionMultiplexer(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.Accounts {
		for region := range client.Regions {
			l = append(l, client.withAccountIDAndRegion(accountID, region))
		}
	}
	return l
}
