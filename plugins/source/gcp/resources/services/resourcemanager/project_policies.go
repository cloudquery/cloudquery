package resourcemanager

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}
