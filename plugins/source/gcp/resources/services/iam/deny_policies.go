package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/iam/v2beta"
)

func DenyPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_deny_policies",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy`,
		Resolver:    fetchDenyPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.GoogleIamV2betaPolicy{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
	}
}
