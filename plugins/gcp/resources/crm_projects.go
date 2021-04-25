package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CrmProjects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_crm_projects",
		Resolver:    fetchCrmProjects,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name: "create_time",
				Type: schema.TypeString,
			},
			{
				Name: "delete_time",
				Type: schema.TypeString,
			},
			{
				Name: "display_name",
				Type: schema.TypeString,
			},
			{
				Name: "etag",
				Type: schema.TypeString,
			},
			{
				Name: "labels",
				Type: schema.TypeJSON,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "parent",
				Type: schema.TypeString,
			},
			{
				Name: "project_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "update_time",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCrmProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	call := c.Services.Crm.Projects.List().Context(ctx)
	for {
		call.PageToken(nextPageToken)
		resp, err := call.Do()
		if err != nil {
			return err
		}
		res <- resp.Projects

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return nil
}
