package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, accountZones := range client.accountsZones {
		l = append(l, client.withAccountID(accountZones.AccountId))
	}
	return l
}

func ZoneMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, accountZones := range client.accountsZones {
		for _, zone := range accountZones.Zones {
			l = append(l, client.withZoneID(accountZones.AccountId, zone))
		}
	}
	return l
}
