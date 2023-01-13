package client

import "github.com/cloudquery/plugin-sdk/schema"

func AccountRegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, accountID := range client.Accounts {
		for _, region := range client.Regions {
			l = append(l, client.withAccountIDAndRegion(accountID, region))
		}
	}
	return l
}
