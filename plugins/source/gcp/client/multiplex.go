package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.projects))
	for i, projectId := range client.projects {
		l[i] = client.withProject(projectId)
	}
	return l
}

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.orgs))
	for i, orgId := range client.orgs {
		l[i] = client.withOrg(orgId)
	}
	return l
}
