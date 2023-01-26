package securitycenter

import (
	"context"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func OrganizationFindings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_securitycenter_organization_findings",
		Description: `https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult`,
		Resolver:    fetchOrganizationFindings,
		Multiplex:   client.OrgMultiplex,
		Transform:   transformers.TransformWithStruct(&pb.ListFindingsResponse_ListFindingsResult{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganization,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Finding.Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchOrganizationFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := "organizations/" + c.OrgId + "/sources/-"
	return fetchFindings(p)(ctx, meta, parent, res)
}
