package security

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	secretmanager "google.golang.org/api/secretmanager/v1"
)

//go:generate cq-gen --resource secrets --config ./gen.hcl --output .
func Secrets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_security_secrets",
		Description: "Secret: A Secret is a logical secret whose value and versions can be accessed",
		Resolver:    fetchSecuritySecrets,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"resource_name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "id",
				Description: "The id of the secret",
				Type:        schema.TypeString,
				Resolver:    ResolveSecuritySecretID,
			},
			{
				Name:        "resource_name",
				Description: "The resource name of the Secret in the format `projects/*/secrets/*`",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "topics",
				Description: "A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions. In the format `projects/*/topics/*`",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveSecuritySecretTopics,
			},
			{
				Name:        "is_automatically_replicated",
				Description: "If true, the secret is automatically replicated by GCP. Otherwise, replications are user-managed.",
				Type:        schema.TypeBool,
				Resolver:    ResolveSecuritySecretIsAutomaticallyReplicated,
			},
			{
				Name:        "create_time",
				Description: "The time at which the Secret was created",
				Type:        schema.TypeString,
			},
			{
				Name:        "etag",
				Description: "Etag of the currently stored Secret",
				Type:        schema.TypeString,
			},
			{
				Name:          "expire_time",
				Description:   "Timestamp in UTC when the Secret is scheduled to expire.",
				Type:          schema.TypeString,
				IgnoreInTests: true, // We don't want to set our terraform resources to expire
			},
			{
				Name:        "labels",
				Description: "The labels assigned to this Secret",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "automatic_replication_customer_managed_encryption_kms_key_name",
				Description:   "If the secret is automatically replicated, contains the customer-managed-encryption kms-key-name. Only valid if 'is_automatically_replicated' is true. If null, then the secret is encrypted with a google-managed key.",
				Type:          schema.TypeString,
				Resolver:      resolveSecretsAutomaticReplicationCustomerManagedEncryptionKmsKeyName,
				IgnoreInTests: true, // Terraform doesn't support customer-managed-encryption for automatically-replicated secrets
			},
			{
				Name:        "next_rotation_time",
				Description: "Timestamp in UTC at which the Secret is scheduled to rotate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Rotation.NextRotationTime"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_security_secret_user_managed_replicas",
				Description: "Describes user-managed replicas of this secret. Empty for automatically replicated secrets",
				Resolver:    fetchSecuritySecretUserManagedReplicas,
				Columns: []schema.Column{
					{
						Name:        "secret_cq_id",
						Description: "Unique CloudQuery ID of gcp_security_secrets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "customer_managed_encryption_kms_key_name",
						Description: "If the replica is encrypted with customer-managed encryption, contains the kms key name. If the column is NULL, the replica is encrypted with a google-managed key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CustomerManagedEncryption.KmsKeyName"),
					},
					{
						Name:        "location",
						Description: "The canonical IDs of the location to replicate data. For example: \"us-east1\"",
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

func fetchSecuritySecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	gcpClient := meta.(*client.Client)

	nextPageToken := ""

	for {
		output, err := gcpClient.Services.SecretManager.Projects.Secrets.List("projects/" + gcpClient.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Secrets
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}

	return nil
}

// extracts the 'id' from the resource-name ("projects/<project_id>/secrets/<id>")
func ResolveSecuritySecretID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	secret := resource.Item.(*secretmanager.Secret)

	split_resource_name := strings.Split(secret.Name, "/")

	return errors.WithStack(resource.Set(c.Name, split_resource_name[len(split_resource_name)-1]))
}
func ResolveSecuritySecretTopics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	secret := resource.Item.(*secretmanager.Secret)

	topicNames := make([]string, 0, len(secret.Topics))

	for _, topic := range secret.Topics {
		if topic == nil {
			continue
		}

		topicNames = append(topicNames, topic.Name)
	}

	return errors.WithStack(resource.Set(c.Name, topicNames))
}
func ResolveSecuritySecretIsAutomaticallyReplicated(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	secret := resource.Item.(*secretmanager.Secret)

	if secret.Replication == nil {
		return nil
	}

	if secret.Replication.Automatic != nil {
		return errors.WithStack(resource.Set(c.Name, true))
	}

	return errors.WithStack(resource.Set(c.Name, false))
}
func resolveSecretsAutomaticReplicationCustomerManagedEncryptionKmsKeyName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	secret := resource.Item.(*secretmanager.Secret)

	if secret.Replication == nil || secret.Replication.Automatic == nil ||
		secret.Replication.Automatic.CustomerManagedEncryption == nil {
		return nil
	}

	return errors.WithStack(resource.Set(c.Name, secret.Replication.Automatic.CustomerManagedEncryption.KmsKeyName))
}
func fetchSecuritySecretUserManagedReplicas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	secret := parent.Item.(*secretmanager.Secret)

	if secret.Replication == nil || secret.Replication.UserManaged == nil {
		return nil
	}

	res <- secret.Replication.UserManaged.Replicas

	return nil
}
