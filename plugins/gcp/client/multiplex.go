package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	if len(client.projects) == 0 {
		return []schema.ClientMeta{meta}
	}

	l := make([]schema.ClientMeta, len(client.projects))
	for i, projectId := range client.projects {
		l[i] = client.withProject(projectId)
	}
	return l
}
