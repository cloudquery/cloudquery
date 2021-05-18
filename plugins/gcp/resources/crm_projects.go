package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/cloudresourcemanager/v3"
)

func CrmProjects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_crm_projects",
		Resolver:     fetchCrmProjects,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
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
	call := c.Services.Crm.Projects.Get("projects/" + c.ProjectId).Context(ctx)
	project, err := call.Do()
	if err != nil {
		return err
	}
	res <- []*cloudresourcemanager.Project{project}
	return nil
}
