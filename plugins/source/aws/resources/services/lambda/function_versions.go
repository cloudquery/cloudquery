// Code generated by codegen; DO NOT EDIT.

package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FunctionVersions() *schema.Table {
	return &schema.Table{
		Name:      "aws_lambda_function_versions",
		Resolver:  fetchLambdaFunctionVersions,
		Multiplex: client.ServiceAccountRegionMultiplexer("lambda"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "architectures",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Architectures"),
			},
			{
				Name:     "code_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CodeSha256"),
			},
			{
				Name:     "code_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CodeSize"),
			},
			{
				Name:     "dead_letter_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeadLetterConfig"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "environment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Environment"),
			},
			{
				Name:     "ephemeral_storage",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EphemeralStorage"),
			},
			{
				Name:     "file_system_configs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FileSystemConfigs"),
			},
			{
				Name:     "function_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FunctionName"),
			},
			{
				Name:     "handler",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Handler"),
			},
			{
				Name:     "image_config_response",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageConfigResponse"),
			},
			{
				Name:     "kms_key_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KMSKeyArn"),
			},
			{
				Name:     "last_modified",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastModified"),
			},
			{
				Name:     "last_update_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdateStatus"),
			},
			{
				Name:     "last_update_status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdateStatusReason"),
			},
			{
				Name:     "last_update_status_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdateStatusReasonCode"),
			},
			{
				Name:     "layers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Layers"),
			},
			{
				Name:     "master_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterArn"),
			},
			{
				Name:     "memory_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MemorySize"),
			},
			{
				Name:     "package_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PackageType"),
			},
			{
				Name:     "revision_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevisionId"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "runtime",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Runtime"),
			},
			{
				Name:     "signing_job_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SigningJobArn"),
			},
			{
				Name:     "signing_profile_version_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SigningProfileVersionArn"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason"),
			},
			{
				Name:     "state_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReasonCode"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Timeout"),
			},
			{
				Name:     "tracing_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TracingConfig"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
		},
	}
}
