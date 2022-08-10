package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, accountZones := range client.accountsZones {
		l = append(l, client.withAccountId(accountZones.AccountId))
	}
	return l
}

func ZoneMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, accountZones := range client.accountsZones {
		for _, zone := range accountZones.Zones {
			l = append(l, client.withZoneId(accountZones.AccountId, zone))
		}
	}
	return l
}
