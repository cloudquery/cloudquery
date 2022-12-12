package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProjectMultiplex(enabledService string) func(schema.ClientMeta) []schema.ClientMeta {
	if _, ok := GcpServices[enabledService]; !ok {
		panic("unknown service: " + enabledService)
	}

	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)
		// preallocate all clients just in case
		l := make([]schema.ClientMeta, 0, len(cl.projects))
		for _, projectId := range cl.projects {
			// This map can only be empty if user has not opted into `EnabledServicesOnly` via the spec
			if len(cl.EnabledServices) == 0 {
				l = append(l, cl.withProject(projectId))
			} else if cl.EnabledServices[projectId] != nil && cl.EnabledServices[projectId][enabledService] {
				l = append(l, cl.withProject(projectId))
			}
		}
		return l
	}
}
