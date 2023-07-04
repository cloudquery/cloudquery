package securitycenter

import (
	"context"

	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func FolderFindings() *schema.Table {
	return &schema.Table{
		Name:          "gcp_securitycenter_folder_findings",
		Description:   `https://cloud.google.com/security-command-center/docs/reference/rest/v1/ListFindingsResponse#ListFindingsResult`,
		Resolver:      fetchFolderFindings,
		Multiplex:     client.FolderMultiplex,
		IsIncremental: true,
		Transform:     client.TransformWithStruct(&pb.ListFindingsResponse_ListFindingsResult{}),
		Columns: []schema.Column{
			{
				Name:       "folder_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveFolder,
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

func fetchFolderFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	// FolderId is already in the format "folders/{id}"
	p := c.FolderId + "/sources/-"
	return fetchFindings("gcp_securitycenter_folder_findings", p)(ctx, meta, parent, res)
}
