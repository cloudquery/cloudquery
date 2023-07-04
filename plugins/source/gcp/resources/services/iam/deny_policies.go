package iam

import (
	pb "cloud.google.com/go/iam/apiv2/iampb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
