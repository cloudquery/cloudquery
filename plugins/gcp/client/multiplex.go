package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for _, projectId := range client.projects {
		l = append(l, client.withProject(projectId))
	}
	return l
}
