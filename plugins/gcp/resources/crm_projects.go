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
		Description:  "A project is a high-level Google Cloud entity",
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Resolver:     fetchCrmProjects,
		Columns: []schema.Column{
			{
				Name:        "create_time",
				Description: "Creation time",
				Type:        schema.TypeString,
			},
			{
				Name:        "delete_time",
				Description: "The time at which this resource was requested for deletion",
				Type:        schema.TypeString,
			},
			{
				Name:        "display_name",
				Description: "A user-assigned display name of the project When present it must be between 4 to 30 characters Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "A checksum computed by the server based on the current value of the Project resource This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "The labels associated with this project",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "name",
				Description: "The unique resource name of the project It is an int64 generated number prefixed by \"projects/\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "parent",
				Description: "A reference to a parent Resource eg, `organizations/123` or `folders/876`",
				Type:        schema.TypeString,
			},
			{
				Name:        "project_id",
				Description: "Immutable The unique, user-assigned id of the project It must be 6 to 30 lowercase ASCII letters, digits, or hyphens It must start with a letter Trailing hyphens are prohibited",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The project lifecycle state  Possible values:   \"STATE_UNSPECIFIED\" - Unspecified state This is only used/useful for distinguishing unset values   \"ACTIVE\" - The normal and active state   \"DELETE_REQUESTED\" - The project has been marked for deletion by the user (by invoking DeleteProject) or by the system (Google Cloud Platform) This can generally be reversed by invoking UndeleteProject",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "The most recent time this resource was modified",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCrmProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	call := c.Services.Crm.Projects.Get("projects/" + c.ProjectId).Context(ctx)
	project, err := call.Do()
	if err != nil {
		return err
	}
	res <- []*cloudresourcemanager.Project{project}
	return nil
}
