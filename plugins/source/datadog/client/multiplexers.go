package client

import "github.com/cloudquery/plugin-sdk/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	l := make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, account := range client.Accounts {
		l = append(l, client.withAccount(account))
	}
	return l
}
