package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, account := range client.Accounts {
		l = append(l, client.withAccount(account))
	}
	return l
}
