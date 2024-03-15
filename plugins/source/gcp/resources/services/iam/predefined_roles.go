package iam

import (
	"context"

	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PredefinedRoles() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_predefined_roles",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/roles#Role`,
		Resolver:    resolvePredefinedProjectIAMRoles,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   client.TransformWithStruct(&iampb.Role{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
	}
}

func resolvePredefinedProjectIAMRoles(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	return fetchIAMRole(ctx, meta, r, res, "")
}
