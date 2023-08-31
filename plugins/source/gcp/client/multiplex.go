package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.projects))
	for i, projectId := range client.projects {
		l[i] = client.withProject(projectId)
	}
	return l
}

func ProjectMultiplexEnabledServices(enabledService string) func(schema.ClientMeta) []schema.ClientMeta {
	if _, ok := GcpServices[enabledService]; !ok {
		panic("unknown service: " + enabledService)
	}

	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)
		// preallocate all clients just in case
		l := make([]schema.ClientMeta, 0, len(cl.projects))
		for _, projectId := range cl.projects {
			// When EnabledServicesOnly is empty, we can assume that all services should be synced
			if len(cl.EnabledServices) == 0 {
				l = append(l, cl.withProject(projectId))
				continue
			}
			// When EnabledServices[projectId] is nil then we can assume that all services should be synced as the Listing of Enabled Services must have failed
			if cl.EnabledServices[projectId] == nil {
				l = append(l, cl.withProject(projectId))
				continue
			}
			// When `projectId` key is set then we can assume listing was completed successfully and only if the enabledService is set should we sync
			if cl.EnabledServices[projectId] != nil && cl.EnabledServices[projectId][enabledService] != nil {
				l = append(l, cl.withProject(projectId))
			}
		}
		return l
	}
}

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.orgs))
	for i := range client.orgs {
		l[i] = client.withOrg(client.orgs[i])
	}
	return l
}

func FolderMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)

	l := make([]schema.ClientMeta, len(client.folderIds))
	for i, folderId := range client.folderIds {
		l[i] = client.withFolder(folderId)
	}
	return l
}

func ProjectLocationMultiplexEnabledServices(service string, locations []string) func(schema.ClientMeta) []schema.ClientMeta {
	if _, ok := GcpServices[service]; !ok {
		panic("unknown service: " + service)
	}

	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)

		l := make([]schema.ClientMeta, 0, len(cl.projects))
		for _, projectId := range cl.projects {
			// This map can only be empty if user has not opted into `EnabledServicesOnly` via the spec
			if len(cl.EnabledServices) == 0 {
				for _, location := range locations {
					l = append(l, cl.withProject(projectId).withLocation(location))
				}
			} else if cl.EnabledServices[projectId] != nil && cl.EnabledServices[projectId][service] != nil {
				for _, location := range locations {
					l = append(l, cl.withProject(projectId).withLocation(location))
				}
			}
		}

		return l
	}
}
