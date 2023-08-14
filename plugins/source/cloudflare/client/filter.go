package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func DeleteAccountFilter(meta schema.ClientMeta, _ *schema.Resource) []any {
	client := meta.(*Client)
	return []any{"account_id", client.AccountId}
}

func DeleteAccountZoneFilter(meta schema.ClientMeta, _ *schema.Resource) []any {
	client := meta.(*Client)
	return []any{"account_id", client.AccountId, "zone_id", client.ZoneId}
}

func DeleteFilter(_ schema.ClientMeta, _ *schema.Resource) []any {
	return []any{}
}
