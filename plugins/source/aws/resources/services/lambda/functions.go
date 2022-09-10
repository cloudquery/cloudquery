package lambda

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type AliasWrapper struct {
	*types.AliasConfiguration
	UrlConfig *lambda.GetFunctionUrlConfigOutput
}

func Functions() *schema.Table {
	return &schema.Table{
		Name:                 "aws_lambda_functions",
		Description:          "AWS Lambda is a serverless compute service that lets you run code without provisioning or managing servers, creating workload-aware cluster scaling logic, maintaining event integrations, or managing runtimes",
		Resolver:             fetchLambdaFunctions,
		Multiplex:            client.ServiceAccountRegionMultiplexer("lambda"),
		PostResourceResolver: resolvePolicyCodeSigningConfig,
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
				Name:          "code",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Code"),
				IgnoreInTests: true,
			},
			{
				Name:        "code_repository_type",
				Description: "The service that's hosting the file.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name:          "concurrency",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Concurrency"),
				IgnoreInTests: true,
			},
			{
				Name:     "configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configuration"),
			},
			{
				Name:        "description",
				Description: "The function's description.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.Description"),
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
				Name:          "aws_lambda_function_event_invoke_configs",
				Description:   "A configuration object that specifies the destination of an event after Lambda processes it. ",
				Resolver:      fetchLambdaFunctionEventInvokeConfigs,
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
						Name:     "alias_configuration",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("AliasConfiguration"),
					},
					{
						Name:     "url_config",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("UrlConfig"),
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
						Type:        schema.TypeInt,
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
						Name:          "environment",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Environment"),
						IgnoreInTests: true,
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
						Name:          "image_config_response",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("ImageConfigResponse"),
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
						Resolver:    schema.PathResolver("LastModified"),
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
						Name:     "tracing_config",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("TracingConfig"),
					},
					{
						Name:        "version",
						Description: "The version of the Lambda function.",
						Type:        schema.TypeString,
					},
					{
						Name:          "vpc_config",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("VpcConfig"),
						IgnoreInTests: true,
					},
					{
						Name: "file_system_configs",
						Type: schema.TypeJSON,
					},
					{
						Name: "layers",
						Type: schema.TypeJSON,
					},
				},
			},
			{
				Name:          "aws_lambda_function_concurrency_configs",
				Description:   "Details about the provisioned concurrency configuration for a function alias or version.",
				Resolver:      fetchLambdaFunctionConcurrencyConfigs,
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
						Resolver:    schema.PathResolver("LastModified"),
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
						Name:     "filter_criteria",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("FilterCriteria"),
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
						Resolver:    schema.PathResolver("SourceAccessConfigurations"),
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
		response, err := svc.ListFunctions(ctx, &input)
		if err != nil {
			return err
		}

		for _, f := range response.Functions {
			getFunctionInput := lambda.GetFunctionInput{
				FunctionName: f.FunctionName,
			}
			funcResponse, err := svc.GetFunction(ctx, &getFunctionInput, func(options *lambda.Options) {
				options.Region = c.Region
			})

			if err != nil {
				if c.IsNotFoundError(err) || c.IsAccessDeniedError(err) {
					c.Logger().Warn().Err(err).Msg("Failed to get function")
					res <- &lambda.GetFunctionOutput{
						Configuration: &f,
					}
					continue
				}
				return err
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
	})
	if err != nil {
		if client.IsAWSError(err, "ResourceNotFoundException") {
			return nil
		}
		return err
	}

	if response != nil {
		if err := resource.Set("policy_revision_id", response.RevisionId); err != nil {
			return err
		}
		var policyDocument map[string]interface{}
		err = json.Unmarshal([]byte(*response.Policy), &policyDocument)
		if err != nil {
			return err
		}
		if err := resource.Set("policy_document", policyDocument); err != nil {
			return err
		}
	}

	// skip getting CodeSigningConfig since containerized lambda functions does not support this feature
	// value can be nil if the caller doesn't have GetFunctionConfiguration permission and only has List*
	lambdaType, ok := resource.Get("code_repository_type").(*string)
	if !ok || *lambdaType == "ECR" {
		return nil
	}

	functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
		FunctionName: r.Configuration.FunctionName,
	})
	if err != nil {
		return err
	}
	if functionSigning.CodeSigningConfigArn == nil {
		return nil
	}

	signing, err := svc.GetCodeSigningConfig(ctx, &lambda.GetCodeSigningConfigInput{
		CodeSigningConfigArn: functionSigning.CodeSigningConfigArn,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if signing.CodeSigningConfig == nil {
		return nil
	}

	if err := resource.Set("code_signing_allowed_publishers_version_arns", signing.CodeSigningConfig.AllowedPublishers.SigningProfileVersionArns); err != nil {
		return err
	}
	if err := resource.Set("code_signing_config_arn", signing.CodeSigningConfig.CodeSigningConfigArn); err != nil {
		return err
	}
	if err := resource.Set("code_signing_config_id", signing.CodeSigningConfig.CodeSigningConfigId); err != nil {
		return err
	}
	if err := resource.Set("code_signing_policies_untrusted_artifact_on_deployment", signing.CodeSigningConfig.CodeSigningPolicies.UntrustedArtifactOnDeployment); err != nil {
		return err
	}
	if err := resource.Set("code_signing_description", signing.CodeSigningConfig.Description); err != nil {
		return err
	}

	location, err := time.LoadLocation("UTC")
	if err != nil {
		return err
	}
	codeSigningLastModified, err := time.ParseInLocation(time.RFC3339, *signing.CodeSigningConfig.LastModified, location)
	if err != nil {
		return err
	}
	return resource.Set("code_signing_last_modified", codeSigningLastModified)
}

func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListFunctionEventInvokeConfigs(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
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
			return err
		}
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		aliases := make([]AliasWrapper, 0, len(output.Aliases))
		for _, a := range output.Aliases {
			alias := a
			urlConfig, err := svc.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
				FunctionName: p.Configuration.FunctionName,
				Qualifier:    alias.Name,
			})
			if err != nil && !c.IsNotFoundError(err) {
				return err
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
			return err
		}
		res <- output.Versions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}

func fetchLambdaFunctionConcurrencyConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListProvisionedConcurrencyConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListProvisionedConcurrencyConfigs(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
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

	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListEventSourceMappingsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListEventSourceMappings(ctx, &config)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.EventSourceMappings
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
