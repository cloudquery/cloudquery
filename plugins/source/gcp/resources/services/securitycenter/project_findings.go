package securitycenter

import (
	"context"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ProjectFindings() *schema.Table {
	return &schema.Table{
		Name:          "gcp_securitycenter_project_findings",
		Description:   `https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult`,
		Resolver:      fetchProjectFindings,
		IsIncremental: true,
		Multiplex:     client.ProjectMultiplexEnabledServices("securitycenter.googleapis.com"),
		Transform:     transformers.TransformWithStruct(&pb.ListFindingsResponse_ListFindingsResult{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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

func fetchProjectFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := "projects/" + c.ProjectId + "/sources/-"
	return fetchFindings("gcp_securitycenter_project_findings", p)(ctx, meta, parent, res)
}
