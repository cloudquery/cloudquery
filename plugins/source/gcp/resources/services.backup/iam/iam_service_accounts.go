package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iam/v1"
)

func IamServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_accounts",
		Description: "An IAM service account A service account is an account for an application or a virtual machine (VM) instance, not a person You can use a service account to call Google APIs To learn more, read the overview of service accounts (https://cloudgooglecom/iam/help/service-accounts/overview) When you create a service account, you specify the project ID that owns the service account, as well as a name that must be unique within the project IAM uses these values to create an email address that identifies the service account",
		Resolver:    fetchIamServiceAccounts,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
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
				Name:        "name",
				Description: "The resource name of the service account In one of the following formats: * `projects/{PROJECT_ID}/serviceAccounts/{EMAIL_ADDRESS}` OR `projects/{PROJECT_ID}/serviceAccounts/{UNIQUE_ID}` OR `projects/-/serviceAccounts/{UNIQUE_ID}",
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
				Name:        "id",
				Description: "The unique, stable numeric ID for the service account Each service account retains its unique ID even if you delete the service account For example, if you delete a service account, then create a new service account with the same name, the new service account has a different unique ID than the deleted service account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UniqueId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_iam_service_account_keys",
				Description: "Represents a service account key A service account has two sets of key-pairs: user-managed, and system-managed User-managed key-pairs can be created and deleted by users Users are responsible for rotating these keys periodically to ensure security of their service accounts Users retain the private key of these key-pairs, and Google retains ONLY the public key System-managed keys are automatically rotated by Google, and are used for signing for a maximum of two weeks The rotation process is probabilistic, and usage of the new key will gradually ramp up and down over the key's lifetime If you cache the public key set for a service account, we recommend that you update the cache every 15 minutes User-managed keys can be added and removed at any time, so it is important to update the cache frequently For Google-managed keys, Google will publish a key at least 6 hours before it is first used for signing and will keep publishing it for at least 6 hours after it was last used for signing Public keys for all service accounts are also published at the OAuth2 Service Account API",
				Resolver:    fetchIamServiceAccountKeys,
				Columns: []schema.Column{
					{
						Name:        "service_account_cq_id",
						Description: "Unique CloudQuery ID of gcp_iam_service_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key_algorithm",
						Description: "Specifies the algorithm (and possibly key size) for the key  Possible values:   \"KEY_ALG_UNSPECIFIED\" - An unspecified key algorithm   \"KEY_ALG_RSA_1024\" - 1k RSA Key   \"KEY_ALG_RSA_2048\" - 2k RSA Key",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_origin",
						Description: "The key origin  Possible values:   \"ORIGIN_UNSPECIFIED\" - Unspecified key origin   \"USER_PROVIDED\" - Key is provided by user   \"GOOGLE_PROVIDED\" - Key is provided by Google",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_type",
						Description: "The key type  Possible values:   \"KEY_TYPE_UNSPECIFIED\" - Unspecified key type The presence of this in the message will immediately result in an error   \"USER_MANAGED\" - User-managed keys (managed and rotated by the user)   \"SYSTEM_MANAGED\" - System-managed keys (managed and rotated by Google)",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The resource name of the service account key in the following format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}`",
						Type:        schema.TypeString,
					},
					{
						Name:        "valid_after_time",
						Description: "The key can be used after this timestamp",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("ValidAfterTime"),
					},
					{
						Name:        "valid_before_time",
						Description: "The key can be used before this timestamp For system-managed key pairs, this timestamp is the end time for the private key signing operation The public key could still be used for verification for a few hours after this time",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("ValidBeforeTime"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchIamServiceAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Iam.Projects.ServiceAccounts.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Accounts
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchIamServiceAccountKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	p := parent.Item.(*iam.ServiceAccount)
	output, err := c.Services.Iam.Projects.ServiceAccounts.Keys.List(p.Name).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	res <- output.Keys
	return nil
}
