package secretsmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecretsmanagerSecrets() *schema.Table {
	return &schema.Table{
		Name:          "aws_secretsmanager_secrets",
		Description:   "A structure that contains the details about a secret",
		Resolver:      fetchSecretsmanagerSecrets,
		Multiplex:     client.ServiceAccountRegionMultiplexer("secretsmanager"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "policy",
				Description: "A JSON-formatted string that describes the permissions that are associated with the attached secret.",
				Type:        schema.TypeJSON,
				Resolver:    fetchSecretsmanagerSecretPolicy,
			},
			{
				Name:        "replication_status",
				Description: "A replication object consisting of a RegionReplicationStatus object and includes a Region, KMSKeyId, status, and status message.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSecretsmanagerSecretReplicationStatus,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the secret",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "created_date",
				Description: "The date and time when a secret was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "deleted_date",
				Description: "The date and time the deletion of the secret occurred",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The user-provided description of the secret.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kms_key_id",
				Description: "The ARN or alias of the Amazon Web Services KMS customer master key (CMK) used to encrypt the SecretString and SecretBinary fields in each version of the secret",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_accessed_date",
				Description: "The last date that this secret was accessed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_changed_date",
				Description: "The last date and time that this secret was modified in any way.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_rotated_date",
				Description: "The most recent date and time that the Secrets Manager rotation process was successfully completed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The friendly name of the secret",
				Type:        schema.TypeString,
			},
			{
				Name:        "owning_service",
				Description: "Returns the name of the service that created the secret.",
				Type:        schema.TypeString,
			},
			{
				Name:        "primary_region",
				Description: "The Region where Secrets Manager originated the secret.",
				Type:        schema.TypeString,
			},
			{
				Name:        "rotation_enabled",
				Description: "Indicates whether automatic, scheduled rotation is enabled for this secret.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "rotation_lambda_arn",
				Description: "The ARN of an Amazon Web Services Lambda function invoked by Secrets Manager to rotate and expire the secret either automatically per the schedule or manually by a call to RotateSecret.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RotationLambdaARN"),
			},
			{
				Name:        "rotation_rules_automatically_after_days",
				Description: "Specifies the number of days between automatic scheduled rotations of the secret",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("RotationRules.AutomaticallyAfterDays"),
			},
			{
				Name:        "secret_versions_to_stages",
				Description: "A list of all of the currently assigned SecretVersionStage staging labels and the SecretVersionId attached to each one",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tags",
				Description: "The list of user-defined tags associated with the secret",
				Type:        schema.TypeJSON,
				Resolver:    resolveSecretsmanagerSecretsTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSecretsmanagerSecrets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SecretsManager
	cfg := secretsmanager.ListSecretsInput{}
	for {
		response, err := svc.ListSecrets(ctx, &cfg, func(options *secretsmanager.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		var secrets []WrappedSecret

		// get more details about the secret
		for _, n := range response.SecretList {
			cfg := secretsmanager.DescribeSecretInput{
				SecretId: n.ARN,
			}
			response, err := svc.DescribeSecret(ctx, &cfg, func(options *secretsmanager.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}

			secrets = append(secrets, WrappedSecret{
				SecretListEntry:   n,
				ReplicationStatus: response.ReplicationStatus,
				RotationRules:     response.RotationRules,
			})
		}

		res <- secrets

		if aws.ToString(response.NextToken) == "" {
			break
		}
		cfg.NextToken = response.NextToken
	}
	return nil
}

func fetchSecretsmanagerSecretPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(WrappedSecret)

	if !ok {
		return fmt.Errorf("expected WrappedSecret but got %T", r)
	}

	cl := meta.(*client.Client)
	svc := cl.Services().SecretsManager
	cfg := secretsmanager.GetResourcePolicyInput{
		SecretId: r.ARN,
	}
	response, err := svc.GetResourcePolicy(ctx, &cfg, func(options *secretsmanager.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	b, err := json.Marshal(response.ResourcePolicy)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, b)
}

func resolveSecretsmanagerSecretReplicationStatus(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(WrappedSecret)
	if !ok {
		return fmt.Errorf("expected WrappedSecret but got %T", r)
	}
	var replicationStatus = make([]map[string]interface{}, len(r.ReplicationStatus))

	for i, replication := range r.ReplicationStatus {
		replicationStatus[i] = map[string]interface{}{
			"kms_key_id":         aws.ToString(replication.KmsKeyId),
			"last_accessed_date": aws.ToTime(replication.LastAccessedDate),
			"region":             aws.ToString(replication.Region),
			"status":             replication.Status,
			"status_massage":     aws.ToString(replication.StatusMessage),
		}
	}
	b, err := json.Marshal(replicationStatus)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, b)
}

func resolveSecretsmanagerSecretsTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(WrappedSecret)
	if !ok {
		return fmt.Errorf("expected SecretListEntry but got %T", r)
	}
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}

type WrappedSecret struct {
	types.SecretListEntry
	RotationRules     *types.RotationRulesType
	ReplicationStatus []types.ReplicationStatusType
}
