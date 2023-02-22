package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, o := range client.orgs {
		l = append(l, client.WithOrg(o))
	}
	return l
}

func OrgRepositoryMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, o := range client.orgs {
		for i := range client.orgRepositories[o] {
			l = append(l, client.WithOrg(o).WithRepository(client.orgRepositories[o][i]))
		}
	}
	return l
}
