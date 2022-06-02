package lambda

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type AliasWrapper struct {
	*types.AliasConfiguration
	UrlConfig *lambda.GetFunctionUrlConfigOutput
}

//go:generate cq-gen --resource ledgers --config gen.hcl --output .
func Functions() *schema.Table {
	return &schema.Table{
		Name:                 "aws_lambda_functions",
		Description:          "AWS Lambda is a serverless compute service that lets you run code without provisioning or managing servers, creating workload-aware cluster scaling logic, maintaining event integrations, or managing runtimes",
		Resolver:             fetchLambdaFunctions,
		Multiplex:            client.ServiceAccountRegionMultiplexer("lambda"),
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolvePolicyCodeSigningConfig,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "policy_document",
				Description: "The resource-based policy.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "policy_revision_id",
				Description: "A unique identifier for the current revision of the policy.",
				Type:        schema.TypeString,
			},
			{
				Name:          "code_signing_allowed_publishers_version_arns",
				Description:   "The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user who can sign a code package.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "code_signing_config_arn",
				Description:   "The Amazon Resource Name (ARN) of the Code signing configuration.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "code_signing_config_id",
				Description:   "Unique identifier for the Code signing configuration.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "code_signing_policies_untrusted_artifact_on_deployment",
				Description:   "Code signing configuration policy for deployment validation failure.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "code_signing_description",
				Description:   "Code signing configuration description.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "code_signing_last_modified",
				Description:   "The date and time that the Code signing configuration was last modified, in ISO-8601 format (YYYY-MM-DDThh:mm:ss.sTZD).",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "code_image_uri",
				Description:   "URI of a container image in the Amazon ECR registry.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Code.ImageUri"),
				IgnoreInTests: true,
			},
			{
				Name:        "code_location",
				Description: "A presigned URL that you can use to download the deployment package.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Code.Location"),
			},
			{
				Name:        "code_repository_type",
				Description: "The service that's hosting the file.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name:          "code_resolved_image_uri",
				Description:   "The resolved URI for the image.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Code.ResolvedImageUri"),
				IgnoreInTests: true,
			},
			{
				Name:          "concurrency_reserved_concurrent_executions",
				Description:   "The number of concurrent executions that are reserved for this function",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("Concurrency.ReservedConcurrentExecutions"),
				IgnoreInTests: true,
			},
			{
				Name:        "architectures",
				Description: "The instruction set architecture that the function supports",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Configuration.Architectures"),
			},
			{
				Name:        "code_sha256",
				Description: "The SHA256 hash of the function's deployment package.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.CodeSha256"),
			},
			{
				Name:        "code_size",
				Description: "The size of the function's deployment package, in bytes.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Configuration.CodeSize"),
			},
			{
				Name:        "dead_letter_config_target_arn",
				Description: "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.DeadLetterConfig.TargetArn"),
			},
			{
				Name:        "description",
				Description: "The function's description.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Description"),
			},
			{
				Name:          "environment_error_code",
				Description:   "The error code.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.Environment.Error.ErrorCode"),
				IgnoreInTests: true,
			},
			{
				Name:          "environment_error_message",
				Description:   "The error message.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.Environment.Error.Message"),
				IgnoreInTests: true,
			},
			{
				Name:        "environment_variables",
				Description: "Environment variable key-value pairs.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Configuration.Environment.Variables"),
			},
			{
				Name:        "ephemeral_storage_size",
				Description: "The size of the function’s /tmp directory.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Configuration.EphemeralStorage.Size"),
			},
			{
				Name:        "arn",
				Description: "The function's Amazon Resource Name (ARN).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.FunctionArn"),
			},
			{
				Name:        "name",
				Description: "The name of the function.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.FunctionName"),
			},
			{
				Name:        "handler",
				Description: "The function that Lambda calls to begin executing your function.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Handler"),
			},
			{
				Name:          "error_code",
				Description:   "Error code.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ImageConfigResponse.Error.ErrorCode"),
				IgnoreInTests: true,
			},
			{
				Name:          "error_message",
				Description:   "Error message.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ImageConfigResponse.Error.Message"),
				IgnoreInTests: true,
			},
			{
				Name:          "image_config_command",
				Description:   "Specifies parameters that you want to pass in with ENTRYPOINT.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.Command"),
				IgnoreInTests: true,
			},
			{
				Name:          "image_config_entry_point",
				Description:   "Specifies the entry point to their application, which is typically the location of the runtime executable.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.EntryPoint"),
				IgnoreInTests: true,
			},
			{
				Name:          "image_config_working_directory",
				Description:   "Specifies the working directory.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.WorkingDirectory"),
				IgnoreInTests: true,
			},
			{
				Name:          "kms_key_arn",
				Description:   "The KMS key that's used to encrypt the function's environment variables",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.KMSKeyArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "last_modified",
				Description: "The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).",
				Type:        schema.TypeTimestamp,
				Resolver:    client.ResolveTimestampField("Configuration.LastModified", TimestampLayoutNano),
			},
			{
				Name:        "last_update_status",
				Description: "The status of the last update that was performed on the function",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.LastUpdateStatus"),
			},
			{
				Name:          "last_update_status_reason",
				Description:   "The reason for the last update that was performed on the function.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.LastUpdateStatusReason"),
				IgnoreInTests: true,
			},
			{
				Name:        "last_update_status_reason_code",
				Description: "The reason code for the last update that was performed on the function.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.LastUpdateStatusReasonCode"),
			},
			{
				Name:          "master_arn",
				Description:   "For Lambda@Edge functions, the ARN of the main function.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.MasterArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "memory_size",
				Description: "The amount of memory available to the function at runtime.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Configuration.MemorySize"),
			},
			{
				Name:        "package_type",
				Description: "The type of deployment package",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.PackageType"),
			},
			{
				Name:        "revision_id",
				Description: "The latest updated revision of the function or alias.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.RevisionId"),
			},
			{
				Name:        "role",
				Description: "The function's execution role.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Role"),
			},
			{
				Name:        "runtime",
				Description: "The runtime environment for the Lambda function.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Runtime"),
			},
			{
				Name:          "signing_job_arn",
				Description:   "The ARN of the signing job.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.SigningJobArn"),
				IgnoreInTests: true,
			},
			{
				Name:          "signing_profile_version_arn",
				Description:   "The ARN of the signing profile version.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.SigningProfileVersionArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "The current state of the function",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.State"),
			},
			{
				Name:          "state_reason",
				Description:   "The reason for the function's current state.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.StateReason"),
				IgnoreInTests: true,
			},
			{
				Name:        "state_reason_code",
				Description: "The reason code for the function's current state",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.StateReasonCode"),
			},
			{
				Name:        "timeout",
				Description: "The amount of time in seconds that Lambda allows a function to run before stopping it.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Configuration.Timeout"),
			},
			{
				Name:        "tracing_config_mode",
				Description: "The tracing mode.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.TracingConfig.Mode"),
			},
			{
				Name:        "version",
				Description: "The version of the Lambda function.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Version"),
			},
			{
				Name:          "vpc_config_security_group_ids",
				Description:   "A list of VPC security groups IDs.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Configuration.VpcConfig.SecurityGroupIds"),
				IgnoreInTests: true,
			},
			{
				Name:          "vpc_config_subnet_ids",
				Description:   "A list of VPC subnet IDs.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("Configuration.VpcConfig.SubnetIds"),
				IgnoreInTests: true,
			},
			{
				Name:          "vpc_config_vpc_id",
				Description:   "The ID of the VPC.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Configuration.VpcConfig.VpcId"),
				IgnoreInTests: true,
			},
			{
				Name:          "tags",
				Description:   "The function's tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html).",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Tags"),
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lambda_function_file_system_configs",
				Description:   "Details about the connection between a Lambda function and an Amazon EFS file system. ",
				Resolver:      fetchLambdaFunctionFileSystemConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "function_arn",
						Description: "The Amazon Resource Name (ARN) of the lambda function",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("arn"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
						Type:        schema.TypeString,
					},
					{
						Name:        "local_mount_path",
						Description: "The path where the function can access the file system, starting with /mnt/.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_lambda_function_layers",
				Description: "An Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).",
				Resolver:    fetchLambdaFunctionLayers,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "function_arn",
						Description: "The Amazon Resource Name (ARN) of the lambda function",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("arn"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the function layer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "code_size",
						Description: "The size of the layer archive in bytes.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:          "signing_job_arn",
						Description:   "The Amazon Resource Name (ARN) of a signing job.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "signing_profile_version_arn",
						Description:   "The Amazon Resource Name (ARN) for a signing profile version.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "aws_lambda_function_event_invoke_configs",
				Description:   "A configuration object that specifies the destination of an event after Lambda processes it. ",
				Resolver:      fetchLambdaFunctionEventInvokeConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "on_failure_destination",
						Description: "The Amazon Resource Name (ARN) of the destination resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationConfig.OnFailure.Destination"),
					},
					{
						Name:        "on_success_destination",
						Description: "The Amazon Resource Name (ARN) of the destination resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationConfig.OnSuccess.Destination"),
					},
					{
						Name:        "function_arn",
						Description: "The Amazon Resource Name (ARN) of the function.",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_modified",
						Description: "The date and time that the configuration was last updated.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "maximum_event_age_in_seconds",
						Description: "The maximum age of a request that Lambda sends to a function for processing.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "maximum_retry_attempts",
						Description: "The maximum number of times to retry when the function returns an error.",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:          "aws_lambda_function_aliases",
				Description:   "Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).",
				Resolver:      fetchLambdaFunctionAliases,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "function_arn",
						Description: "The Amazon Resource Name (ARN) of the lambda function",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("arn"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the alias.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AliasConfiguration.AliasArn"),
					},
					{
						Name:        "description",
						Description: "A description of the alias.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AliasConfiguration.Description"),
					},
					{
						Name:        "function_version",
						Description: "The function version that the alias invokes.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AliasConfiguration.FunctionVersion"),
					},
					{
						Name:        "name",
						Description: "The name of the alias.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AliasConfiguration.Name"),
					},
					{
						Name:        "revision_id",
						Description: "A unique identifier that changes when you update the alias.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AliasConfiguration.RevisionId"),
					},
					{
						Name:        "routing_config_additional_version_weights",
						Description: "The second version, and the percentage of traffic that's routed to it.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("AliasConfiguration.RoutingConfig.AdditionalVersionWeights"),
					},
					{
						Name:        "url_config_auth_type",
						Description: "The type of authentication that your function URL uses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UrlConfig.AuthType"),
					},
					{
						Name:        "url_config_creation_time",
						Description: "When the function URL was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("UrlConfig.CreationTime"),
					},
					{
						Name:        "url_config_function_arn",
						Description: "The Amazon Resource Name (ARN) of your function.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UrlConfig.FunctionArn"),
					},
					{
						Name:        "url_config_function_url",
						Description: "The HTTP URL endpoint for your function.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UrlConfig.FunctionUrl"),
					},
					{
						Name:        "url_config_last_modified_time",
						Description: "When the function URL configuration was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("UrlConfig.LastModifiedTime"),
					},
					{
						Name:        "url_config_cors",
						Description: "The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL.",
						Type:        schema.TypeJSON,
						Resolver:    resolveFunctionAliasesUrlConfigCors,
					},
				},
			},
			{
				Name:        "aws_lambda_function_versions",
				Description: "Details about a function's configuration.",
				Resolver:    fetchLambdaFunctionVersions,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "architectures",
						Description: "The instruction set architecture that the function supports",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "code_sha256",
						Description: "The SHA256 hash of the function's deployment package.",
						Type:        schema.TypeString,
					},
					{
						Name:        "code_size",
						Description: "The size of the function's deployment package, in bytes.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "dead_letter_config_target_arn",
						Description: "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DeadLetterConfig.TargetArn"),
					},
					{
						Name:        "description",
						Description: "The function's description.",
						Type:        schema.TypeString,
					},
					{
						Name:          "environment_error_error_code",
						Description:   "The error code.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Environment.Error.ErrorCode"),
						IgnoreInTests: true,
					},
					{
						Name:          "environment_error_message",
						Description:   "The error message.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Environment.Error.Message"),
						IgnoreInTests: true,
					},
					{
						Name:        "environment_variables",
						Description: "Environment variable key-value pairs.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("Environment.Variables"),
					},
					{
						Name:        "ephemeral_storage_size",
						Description: "The size of the function’s /tmp directory.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("EphemeralStorage.Size"),
					},
					{
						Name:        "function_arn",
						Description: "The function's Amazon Resource Name (ARN).",
						Type:        schema.TypeString,
					},
					{
						Name:        "function_name",
						Description: "The name of the function.",
						Type:        schema.TypeString,
					},
					{
						Name:        "handler",
						Description: "The function that Lambda calls to begin executing your function.",
						Type:        schema.TypeString,
					},
					{
						Name:          "error_code",
						Description:   "Error code.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ImageConfigResponse.Error.ErrorCode"),
						IgnoreInTests: true,
					},
					{
						Name:          "error_message",
						Description:   "Error message.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ImageConfigResponse.Error.Message"),
						IgnoreInTests: true,
					},
					{
						Name:          "image_config_command",
						Description:   "Specifies parameters that you want to pass in with ENTRYPOINT.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ImageConfigResponse.ImageConfig.Command"),
						IgnoreInTests: true,
					},
					{
						Name:          "image_config_entry_point",
						Description:   "Specifies the entry point to their application, which is typically the location of the runtime executable.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ImageConfigResponse.ImageConfig.EntryPoint"),
						IgnoreInTests: true,
					},
					{
						Name:          "image_config_working_directory",
						Description:   "Specifies the working directory.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ImageConfigResponse.ImageConfig.WorkingDirectory"),
						IgnoreInTests: true,
					},
					{
						Name:          "kms_key_arn",
						Description:   "The KMS key that's used to encrypt the function's environment variables",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KMSKeyArn"),
						IgnoreInTests: true,
					},
					{
						Name:        "last_modified",
						Description: "The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ResolveTimestampField("LastModified", TimestampLayout),
					},
					{
						Name:        "last_update_status",
						Description: "The status of the last update that was performed on the function",
						Type:        schema.TypeString,
					},
					{
						Name:          "last_update_status_reason",
						Description:   "The reason for the last update that was performed on the function.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "last_update_status_reason_code",
						Description: "The reason code for the last update that was performed on the function.",
						Type:        schema.TypeString,
					},
					{
						Name:          "master_arn",
						Description:   "For Lambda@Edge functions, the ARN of the main function.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "memory_size",
						Description: "The amount of memory available to the function at runtime.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "package_type",
						Description: "The type of deployment package",
						Type:        schema.TypeString,
					},
					{
						Name:        "revision_id",
						Description: "The latest updated revision of the function or alias.",
						Type:        schema.TypeString,
					},
					{
						Name:        "role",
						Description: "The function's execution role.",
						Type:        schema.TypeString,
					},
					{
						Name:        "runtime",
						Description: "The runtime environment for the Lambda function.",
						Type:        schema.TypeString,
					},
					{
						Name:          "signing_job_arn",
						Description:   "The ARN of the signing job.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "signing_profile_version_arn",
						Description:   "The ARN of the signing profile version.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "state",
						Description: "The current state of the function",
						Type:        schema.TypeString,
					},
					{
						Name:          "state_reason",
						Description:   "The reason for the function's current state.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "state_reason_code",
						Description: "The reason code for the function's current state",
						Type:        schema.TypeString,
					},
					{
						Name:        "timeout",
						Description: "The amount of time in seconds that Lambda allows a function to run before stopping it.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "tracing_config_mode",
						Description: "The tracing mode.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TracingConfig.Mode"),
					},
					{
						Name:        "version",
						Description: "The version of the Lambda function.",
						Type:        schema.TypeString,
					},
					{
						Name:          "vpc_config_security_group_ids",
						Description:   "A list of VPC security groups IDs.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("VpcConfig.SecurityGroupIds"),
						IgnoreInTests: true,
					},
					{
						Name:          "vpc_config_subnet_ids",
						Description:   "A list of VPC subnet IDs.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("VpcConfig.SubnetIds"),
						IgnoreInTests: true,
					},
					{
						Name:          "vpc_config_vpc_id",
						Description:   "The ID of the VPC.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("VpcConfig.VpcId"),
						IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_lambda_function_version_file_system_configs",
						Description:   "Details about the connection between a Lambda function and an Amazon EFS file system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html).",
						Resolver:      fetchLambdaFunctionVersionFileSystemConfigs,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "function_version_cq_id",
								Description: "Unique CloudQuery ID of aws_lambda_function_versions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
								Type:        schema.TypeString,
							},
							{
								Name:        "local_mount_path",
								Description: "The path where the function can access the file system, starting with /mnt/.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_lambda_function_version_layers",
						Description: "An Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).",
						Resolver:    fetchLambdaFunctionVersionLayers,
						Columns: []schema.Column{
							{
								Name:        "function_version_cq_id",
								Description: "Unique CloudQuery ID of aws_lambda_function_versions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) of the function layer.",
								Type:        schema.TypeString,
							},
							{
								Name:        "code_size",
								Description: "The size of the layer archive in bytes.",
								Type:        schema.TypeBigInt,
							},
							{
								Name:          "signing_job_arn",
								Description:   "The Amazon Resource Name (ARN) of a signing job.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "signing_profile_version_arn",
								Description:   "The Amazon Resource Name (ARN) for a signing profile version.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
						},
					},
				},
			},
			{
				Name:          "aws_lambda_function_concurrency_configs",
				Description:   "Details about the provisioned concurrency configuration for a function alias or version.",
				Resolver:      fetchLambdaFunctionConcurrencyConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allocated_provisioned_concurrent_executions",
						Description: "The amount of provisioned concurrency allocated.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "available_provisioned_concurrent_executions",
						Description: "The amount of provisioned concurrency available.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "function_arn",
						Description: "The Amazon Resource Name (ARN) of the alias or version.",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_modified",
						Description: "The date and time that a user last updated the configuration, in ISO 8601 format (https://www.iso.org/iso-8601-date-and-time-format.html).",
						Type:        schema.TypeTimestamp,
						Resolver:    client.ResolveTimestampField("LastModified", TimestampLayout),
					},
					{
						Name:        "requested_provisioned_concurrent_executions",
						Description: "The amount of provisioned concurrency requested.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "status",
						Description: "The status of the allocation process.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_reason",
						Description: "For failed allocations, the reason that provisioned concurrency could not be allocated.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_lambda_function_event_source_mappings",
				Description:   "A mapping between an Amazon Web Services resource and a Lambda function",
				Resolver:      fetchLambdaFunctionEventSourceMappings,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "function_cq_id",
						Description: "Unique CloudQuery ID of aws_lambda_functions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "batch_size",
						Description: "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function",
						Type:        schema.TypeInt,
					},
					{
						Name:        "bisect_batch_on_function_error",
						Description: "(Streams only) If the function returns an error, split the batch in two and retry",
						Type:        schema.TypeBool,
					},
					{
						Name:        "on_failure_destination",
						Description: "The Amazon Resource Name (ARN) of the destination resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationConfig.OnFailure.Destination"),
					},
					{
						Name:        "on_success_destination",
						Description: "The Amazon Resource Name (ARN) of the destination resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DestinationConfig.OnSuccess.Destination"),
					},
					{
						Name:        "event_source_arn",
						Description: "The Amazon Resource Name (ARN) of the event source.",
						Type:        schema.TypeString,
					},
					{
						Name:        "criteria_filters",
						Description: "A list of filters.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveFunctionEventSourceMappingsCriteriaFilters,
					},
					{
						Name:        "function_arn",
						Description: "The ARN of the Lambda function.",
						Type:        schema.TypeString,
					},
					{
						Name:        "function_response_types",
						Description: "(Streams only) A list of current response type enums applied to the event source mapping.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "last_modified",
						Description: "The date that the event source mapping was last updated or that its state changed.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "last_processing_result",
						Description: "The result of the last Lambda invocation of your function.",
						Type:        schema.TypeString,
					},
					{
						Name:        "maximum_batching_window_in_seconds",
						Description: "(Streams and Amazon SQS standard queues) The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function",
						Type:        schema.TypeInt,
					},
					{
						Name:        "maximum_record_age_in_seconds",
						Description: "(Streams only) Discard records older than the specified age",
						Type:        schema.TypeInt,
					},
					{
						Name:        "maximum_retry_attempts",
						Description: "(Streams only) Discard records after the specified number of retries",
						Type:        schema.TypeInt,
					},
					{
						Name:        "parallelization_factor",
						Description: "(Streams only) The number of batches to process concurrently from each shard. The default value is 1.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "queues",
						Description: "(Amazon MQ) The name of the Amazon MQ broker destination queue to consume.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "self_managed_event_source_endpoints",
						Description: "The list of bootstrap servers for your Kafka brokers in the following format: \"KAFKA_BOOTSTRAP_SERVERS\": [\"abc.xyz.com:xxxx\",\"abc2.xyz.com:xxxx\"].",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("SelfManagedEventSource.Endpoints"),
					},
					{
						Name:        "source_access_configurations",
						Description: "An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.",
						Type:        schema.TypeJSON,
						Resolver:    resolveFunctionEventSourceMappingsSourceAccessConfigurations,
					},
					{
						Name:        "starting_position",
						Description: "The position in a stream from which to start reading",
						Type:        schema.TypeString,
					},
					{
						Name:        "starting_position_timestamp",
						Description: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "state",
						Description: "The state of the event source mapping",
						Type:        schema.TypeString,
					},
					{
						Name:        "state_transition_reason",
						Description: "Indicates whether a user or Lambda made the last change to the event source mapping.",
						Type:        schema.TypeString,
					},
					{
						Name:        "topics",
						Description: "The name of the Kafka topic.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "tumbling_window_in_seconds",
						Description: "(Streams only) The duration in seconds of a processing window",
						Type:        schema.TypeInt,
					},
					{
						Name:        "uuid",
						Description: "The identifier of the event source mapping.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UUID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lambda.ListFunctionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	for {
		response, err := svc.ListFunctions(ctx, &input, func(options *lambda.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		for _, f := range response.Functions {
			getFunctionInput := lambda.GetFunctionInput{
				FunctionName: f.FunctionName,
			}
			funcResponse, err := svc.GetFunction(ctx, &getFunctionInput, func(options *lambda.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					return nil
				}
				return diag.WrapError(err)
			}

			res <- funcResponse
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}
func resolvePolicyCodeSigningConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Lambda

	response, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil {
		if client.IsAWSError(err, "ResourceNotFoundException") {
			return nil
		}
		return diag.WrapError(err)
	}

	if response != nil {
		if err := resource.Set("policy_revision_id", response.RevisionId); err != nil {
			return diag.WrapError(err)
		}
		var policyDocument map[string]interface{}
		err = json.Unmarshal([]byte(*response.Policy), &policyDocument)
		if err != nil {
			return diag.WrapError(err)
		}
		if err := resource.Set("policy_document", policyDocument); err != nil {
			return diag.WrapError(err)
		}
	}

	// skip getting CodeSigningConfig since containerized lambda functions does not support this feature
	lambdaType := resource.Get("code_repository_type").(*string)
	if *lambdaType == "ECR" {
		return nil
	}

	functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	if functionSigning.CodeSigningConfigArn == nil {
		return nil
	}

	signing, err := svc.GetCodeSigningConfig(ctx, &lambda.GetCodeSigningConfigInput{
		CodeSigningConfigArn: functionSigning.CodeSigningConfigArn,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	if signing.CodeSigningConfig == nil {
		return nil
	}

	if err := resource.Set("code_signing_allowed_publishers_version_arns", signing.CodeSigningConfig.AllowedPublishers.SigningProfileVersionArns); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("code_signing_config_arn", signing.CodeSigningConfig.CodeSigningConfigArn); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("code_signing_config_id", signing.CodeSigningConfig.CodeSigningConfigId); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("code_signing_policies_untrusted_artifact_on_deployment", signing.CodeSigningConfig.CodeSigningPolicies.UntrustedArtifactOnDeployment); err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("code_signing_description", signing.CodeSigningConfig.Description); err != nil {
		return diag.WrapError(err)
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return diag.WrapError(err)
	}
	codeSigningLastModified, err := time.ParseInLocation(time.RFC3339, *signing.CodeSigningConfig.LastModified, location)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set("code_signing_last_modified", codeSigningLastModified))
}
func fetchLambdaFunctionFileSystemConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}

	res <- r.Configuration.FileSystemConfigs
	return nil
}
func fetchLambdaFunctionLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*lambda.GetFunctionOutput)
	if r.Configuration == nil {
		return nil
	}

	res <- r.Configuration.Layers
	return nil
}
func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListFunctionEventInvokeConfigs(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.FunctionEventInvokeConfigs
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Lambda
	config := lambda.ListAliasesInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListAliases(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		aliases := make([]AliasWrapper, 0, len(output.Aliases))
		for _, a := range output.Aliases {
			alias := a
			urlConfig, err := svc.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
				FunctionName: p.Configuration.FunctionName,
				Qualifier:    alias.Name,
			})
			if err != nil && !c.IsNotFoundError(err) {
				return diag.WrapError(err)
			}
			aliases = append(aliases, AliasWrapper{&alias, urlConfig})
		}
		res <- aliases
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func resolveFunctionAliasesUrlConfigCors(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(AliasWrapper)
	if p.UrlConfig == nil || p.UrlConfig.Cors == nil {
		return nil
	}
	data, err := json.Marshal(p.UrlConfig.Cors)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}
func fetchLambdaFunctionVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListVersionsByFunctionInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListVersionsByFunction(ctx, &config)
		if err != nil {
			if meta.(*client.Client).IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		res <- output.Versions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionVersionFileSystemConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.FunctionConfiguration)

	res <- r.FileSystemConfigs
	return nil
}
func fetchLambdaFunctionVersionLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.FunctionConfiguration)

	res <- r.Layers
	return nil
}
func fetchLambdaFunctionConcurrencyConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListProvisionedConcurrencyConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListProvisionedConcurrencyConfigs(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.ProvisionedConcurrencyConfigs
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionEventSourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListEventSourceMappingsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListEventSourceMappings(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.EventSourceMappings
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func resolveFunctionEventSourceMappingsCriteriaFilters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EventSourceMappingConfiguration)
	if p.FilterCriteria == nil {
		return nil
	}

	filters := make([]string, 0, len(p.FilterCriteria.Filters))
	for _, f := range p.FilterCriteria.Filters {
		filters = append(filters, *f.Pattern)
	}

	return diag.WrapError(resource.Set(c.Name, filters))
}
func resolveFunctionEventSourceMappingsSourceAccessConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EventSourceMappingConfiguration)
	if len(p.SourceAccessConfigurations) == 0 {
		return nil
	}

	data, err := json.Marshal(p.SourceAccessConfigurations)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, data))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================
