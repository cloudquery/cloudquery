package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IamServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_accounts",
		Description: "An IAM service account A service account is an account for an application or a virtual machine (VM) instance, not a person You can use a service account to call Google APIs To learn more, read the overview of service accounts (https://cloudgooglecom/iam/help/service-accounts/overview) When you create a service account, you specify the project ID that owns the service account, as well as a name that must be unique within the project IAM uses these values to create an email address that identifies the service account",
		Resolver:    fetchIamServiceAccounts,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:        "description",
				Description: "A user-specified, human-readable description of the service account The maximum length is 256 UTF-8 bytes",
				Type:        schema.TypeString,
			},
			{
				Name:        "disabled",
				Description: "Whether the service account is disabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "display_name",
				Description: "A user-specified, human-readable name for the service account The maximum length is 100 UTF-8 bytes",
				Type:        schema.TypeString,
			},
			{
				Name:        "email",
				Description: "The email address of the service account",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "Deprecated Do not use",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The resource name of the service account Use one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` * `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` As an alternative, you can use the `-` wildcard character instead of the project ID: * `projects/-/serviceAccounts/{EMAIL_ADDRESS}` * `projects/-/serviceAccounts/{UNIQUE_ID}` When possible, avoid using the `-` wildcard character, because it can cause response messages to contain misleading error codes For example, if you try to get the service account `projects/-/serviceAccounts/fake@examplecom`, which does not exist, the response contains an HTTP `403 Forbidden` error instead of a `404 Not Found` error",
				Type:        schema.TypeString,
			},
			{
				Name:        "oauth2_client_id",
				Description: "The OAuth 20 client ID for the service account",
				Type:        schema.TypeString,
			},
			{
				Name:        "project_id",
				Description: "The ID of the project that owns the service account",
				Type:        schema.TypeString,
			},
			{
				Name:        "unique_id",
				Description: "The unique, stable numeric ID for the service account Each service account retains its unique ID even if you delete the service account For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamServiceAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Iam.Projects.ServiceAccounts.List("projects/" + c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		res <- output.Accounts
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
