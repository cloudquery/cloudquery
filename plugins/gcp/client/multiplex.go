package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.projects))
	for i, projectId := range client.projects {
		l[i] = client.withProject(projectId)
	}
	return l
}

// ProjectMultiplexEnabledAPIs returns a project multiplexer but filters those project who have disabled apis
func ProjectMultiplexEnabledAPIs(enabledService GcpService) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)

		// preallocate all clients just in case
		l := make([]schema.ClientMeta, 0, len(cl.projects))
		for _, projectId := range cl.projects {
			if cl.EnabledServices[projectId] != nil && cl.EnabledServices[projectId][enabledService] {
				l = append(l, cl.withProject(projectId))
			}
		}
		return l
	}
}
