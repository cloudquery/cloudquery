package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func DeleteAccountFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{"account_id", client.AccountId}
}

func DeleteFilter(_ schema.ClientMeta, _ *schema.Resource) []interface{} {
	return []interface{}{}
}
