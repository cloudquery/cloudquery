package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectMultiplex(enabledService GcpService) func(schema.ClientMeta) []schema.ClientMeta {
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
