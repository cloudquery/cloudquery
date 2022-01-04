package resources

import (
	"context"
	"time"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_projects",
		Description:  "Project represents a DigitalOcean Project configuration.",
		Resolver:     fetchProjects,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "id",
				Description: "The unique universal identifier of this project.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "owner_uuid",
				Description: "The unique universal identifier of the project owner.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OwnerUUID"),
			},
			{
				Name:        "owner_id",
				Description: "The integer id of the project owner.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OwnerID"),
			},
			{
				Name:        "name",
				Description: "The human-readable name for the project. The maximum length is 175 characters and the name must be unique.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "The description of the project. The maximum length is 255 characters.",
				Type:        schema.TypeString,
			},
			{
				Name:        "purpose",
				Description: "The purpose of the project. The maximum length is 255 characters. It can\nhave one of the following values:\n\n- Just trying out DigitalOcean\n- Class project / Educational purposes\n- Website or blog\n- Web Application\n- Service or API\n- Mobile Application\n- Machine learning / AI / Data processing\n- IoT\n- Operational / Developer tooling\n\nIf another value for purpose is specified, for example, \"your custom purpose\",\nyour purpose will be stored as `Other: your custom purpose`.\n",
				Type:        schema.TypeString,
			},
			{
				Name:        "environment",
				Description: "The environment of the project's resources.",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_default",
				Description: "If true, all resources will be added to this project if no project is specified.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "created_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the project was created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "updated_at",
				Description: "A time value given in ISO8601 combined date and time format that represents when the project was updated.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_project_resources",
				Description: "ProjectResource is the projects API's representation of a resource.",
				Resolver:    fetchProjectResources,
				Columns: []schema.Column{
					{
						Name:        "project_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "urn",
						Description: "The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("URN"),
					},
					{
						Name:        "assigned_at",
						Description: "A time value given in ISO8601 combined date and time format that represents when the project was created.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("AssignedAt", time.RFC3339),
					},
					{
						Name:        "links_self",
						Description: "The links object contains the self object, which contains the resource relationship.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Links.Self"),
					},
					{
						Name:        "status",
						Description: "The status of assigning and fetching the resources.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		projects, resp, err := svc.DoClient.Projects.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- projects
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
func fetchProjectResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)

	project := parent.Item.(godo.Project)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		resources, resp, err := svc.DoClient.Projects.ListResources(ctx, project.ID, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- resources
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}
