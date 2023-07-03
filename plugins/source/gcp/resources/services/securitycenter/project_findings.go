package securitycenter

import (
	"context"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ProjectFindings() *schema.Table {
	return &schema.Table{
		Name:          "gcp_securitycenter_project_findings",
		Description:   `https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult`,
		Resolver:      fetchProjectFindings,
		IsIncremental: true,
		Multiplex:     client.ProjectMultiplexEnabledServices("securitycenter.googleapis.com"),
		Transform:     client.TransformWithStruct(&pb.ListFindingsResponse_ListFindingsResult{}),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Finding.Name"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchProjectFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := "projects/" + c.ProjectId + "/sources/-"
	return fetchFindings("gcp_securitycenter_project_findings", p)(ctx, meta, parent, res)
}
