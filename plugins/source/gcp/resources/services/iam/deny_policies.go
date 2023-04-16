package iam

import (
	pb "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func DenyPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_deny_policies",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy`,
		Resolver:    fetchDenyPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Policy{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
