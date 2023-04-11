package resourcemanager

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/cloudresourcemanager/v3"
)

func ProjectPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_project_policies",
		Description: `https://cloud.google.com/resource-manager/reference/rest/Shared.Types/Policy`,
		Resolver:    fetchProjectPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Policy{}),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
	}
}
