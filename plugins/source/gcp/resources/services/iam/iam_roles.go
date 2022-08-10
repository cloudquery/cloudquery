package iam

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/iam/v1"
)

func IamRoles() *schema.Table {
	return &schema.Table{
		Name:         "gcp_iam_roles",
		Description:  "A role in the Identity and Access Management API",
		Resolver:     fetchIamRoles,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "deleted",
				Description: "The current deleted state of the role This field is read only It will be ignored in calls to CreateRole and UpdateRole",
				Type:        schema.TypeBool,
			},
			{
				Name:        "description",
				Description: "A human-readable description for the role",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "Used to perform a consistent read-modify-write",
				Type:        schema.TypeString,
			},
			{
				Name:          "included_permissions",
				Description:   "The names of the permissions this role grants when bound in an IAM policy",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "name",
				Description: "The name of the role",
				Type:        schema.TypeString,
			},
			{
				Name:        "stage",
				Description: "The current launch stage of the role If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role  Possible values:   \"ALPHA\" - The user has indicated this role is currently in an Alpha phase If this launch stage is selected, the `stage` field will not be included when requesting the definition for a given role   \"BETA\" - The user has indicated this role is currently in a Beta phase   \"GA\" - The user has indicated this role is generally available   \"DEPRECATED\" - The user has indicated this role is being deprecated   \"DISABLED\" - This role is disabled and will not contribute permissions to any members it is granted to in policies   \"EAP\" - The user has indicated this role is currently in an EAP phase",
				Type:        schema.TypeString,
			},
			{
				Name:        "title",
				Description: "A human-readable title for the role Typically this is limited to 100 UTF-8 bytes",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Iam.Projects.Roles.List("projects/" + c.ProjectId).PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return diag.WrapError(err)
		}
		output := list.(*iam.ListRolesResponse)

		res <- output.Roles
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
