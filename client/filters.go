package client

import "github.com/cloudquery/cq-provider-sdk/plugin/schema"

func DeleteAccountFilter(meta schema.ClientMeta) []interface{} {
	client := meta.(*Client)
	return []interface{}{"account_id", client.AccountID}
}

func DeleteAccountRegionFilter(meta schema.ClientMeta) []interface{} {
	client := meta.(*Client)
	return []interface{}{"account_id", client.AccountID, "Region", client.Region}
}
