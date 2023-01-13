package resourcemanager

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/cloudresourcemanager/v3"
)

func ProjectPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_project_policies",
		Description: `https://cloud.google.com/resource-manager/reference/rest/Shared.Types/Policy`,
		Resolver:    fetchProjectPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Policy{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
	}
}
