package resource_manager

import (
	"context"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudresourcemanager/v3"
)

func ResourceManagerFolders() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resource_manager_folders",
		Description: "A folder in an organization's resource hierarchy, used to organize that organization's resources",
		Resolver:    fetchResourceManagerFolders,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},

		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "policy",
				Description: "Access control policy for a resource",
				Type:        schema.TypeJSON,
				Resolver:    resolveResourceManagerFolderPolicy,
			},
			{
				Name:        "create_time",
				Description: "Timestamp when the folder was created",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreateTime"),
			},
			{
				Name:        "delete_time",
				Description: "Timestamp when the folder was requested to be deleted",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("DeleteTime"),
			},
			{
				Name:        "display_name",
				Description: "The folder's display name A folder's display name must be unique amongst its siblings For example, no two folders with the same parent can share the same display name The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters This is captured by the regular expression: `[\\p{L}\\p{N}]([\\p{L}\\p{N}_- ]{0,28}[\\p{L}\\p{N}])?`",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "A checksum computed by the server based on the current value of the folder resource This may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The resource name of the folder Its format is `folders/{folder_id}`, for example: \"folders/1234\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "parent",
				Description: "The folder's parent's resource name Updates to the folder's parent must be performed using MoveFolder",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The lifecycle state of the folder Updates to the state must be performed using DeleteFolder and UndeleteFolder  Possible values:   \"STATE_UNSPECIFIED\" - Unspecified state   \"ACTIVE\" - The normal and active state   \"DELETE_REQUESTED\" - The folder has been marked for deletion by the user",
				Type:        schema.TypeString,
			},
			{
				Name:        "update_time",
				Description: "Timestamp when the folder was last modified",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("UpdateTime"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchResourceManagerFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	//Todo service account needs specific permissions to list folders https://cloud.google.com/resource-manager/docs/creating-managing-folders#folder-permissions
	output, err := c.Services.ResourceManager.Folders.List().Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output.Folders
	return nil
}
func resolveResourceManagerFolderPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	p := resource.Item.(*cloudresourcemanager.Folder)
	output, err := cl.Services.ResourceManager.Projects.
		GetIamPolicy("folders/"+p.Name, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	var policy map[string]interface{}
	data, err := json.Marshal(output)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(resource.Set(c.Name, policy))
}
