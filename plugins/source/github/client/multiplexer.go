package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	for o, repos := range client.orgRepositories {
		for _, repo := range repos {
			l = append(l, client.WithOrg(o).WithRepository(repo))
		}
	}
	return l
}
