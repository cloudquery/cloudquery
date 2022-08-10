package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, o := range client.Orgs {
		l = append(l, client.WithOrg(o))
	}
	return l
}
