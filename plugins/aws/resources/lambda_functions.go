package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LambdaFunctions() *schema.Table {
	return &schema.Table{
		Name:                 "aws_lambda_functions",
		Resolver:             fetchLambdaFunctions,
		Multiplex:            client.AccountRegionMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: resolvePolicyCodeSigningConfig,
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
				Name: "policy_document",
				Type: schema.TypeJSON,
			},
			{
				Name: "policy_revision_id",
				Type: schema.TypeString,
			},
			{
				Name: "code_signing_allowed_publishers_version_arns",
				Type: schema.TypeStringArray,
			},
			{
				Name: "code_signing_config_arn",
				Type: schema.TypeString,
			},
			{
				Name: "code_signing_config_id",
				Type: schema.TypeString,
			},
			{
				Name: "code_signing_policies_untrusted_artifact_on_deployment",
				Type: schema.TypeString,
			},
			{
				Name: "code_signing_description",
				Type: schema.TypeString,
			},
			{
				Name: "code_signing_last_modified",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "code_image_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.ImageUri"),
			},
			{
				Name:     "code_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.Location"),
			},
			{
				Name:     "code_repository_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name:     "code_resolved_image_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.ResolvedImageUri"),
			},
			{
				Name:     "concurrency_reserved_concurrent_executions",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Concurrency.ReservedConcurrentExecutions"),
			},
			{
				Name:     "code_sha256",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.CodeSha256"),
			},
			{
				Name:     "code_size",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Configuration.CodeSize"),
			},
			{
				Name:     "dead_letter_config_target_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.DeadLetterConfig.TargetArn"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Description"),
			},
			{
				Name:     "environment_error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Environment.Error.ErrorCode"),
			},
			{
				Name:     "environment_error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Environment.Error.Message"),
			},
			{
				Name:     "environment_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configuration.Environment.Variables"),
			},
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.FunctionArn"),
			},
			{
				Name:     "function_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.FunctionName"),
			},
			{
				Name:     "handler",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Handler"),
			},
			{
				Name:     "error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ImageConfigResponse.Error.ErrorCode"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ImageConfigResponse.Error.Message"),
			},
			{
				Name:     "image_config_command",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.Command"),
			},
			{
				Name:     "image_config_entry_point",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.EntryPoint"),
			},
			{
				Name:     "image_config_working_directory",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.ImageConfigResponse.ImageConfig.WorkingDirectory"),
			},
			{
				Name:     "kms_key_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.KMSKeyArn"),
			},
			{
				Name:     "last_modified",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.LastModified"),
			},
			{
				Name:     "last_update_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.LastUpdateStatus"),
			},
			{
				Name:     "last_update_status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.LastUpdateStatusReason"),
			},
			{
				Name:     "last_update_status_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.LastUpdateStatusReasonCode"),
			},
			{
				Name:     "master_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.MasterArn"),
			},
			{
				Name:     "memory_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Configuration.MemorySize"),
			},
			{
				Name:     "package_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.PackageType"),
			},
			{
				Name:     "revision_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.RevisionId"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Role"),
			},
			{
				Name:     "runtime",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Runtime"),
			},
			{
				Name:     "signing_job_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.SigningJobArn"),
			},
			{
				Name:     "signing_profile_version_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.SigningProfileVersionArn"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.State"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.StateReason"),
			},
			{
				Name:     "state_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.StateReasonCode"),
			},
			{
				Name:     "timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Configuration.Timeout"),
			},
			{
				Name:     "tracing_config_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.TracingConfig.Mode"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.Version"),
			},
			{
				Name:     "vpc_config_security_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Configuration.VpcConfig.SecurityGroupIds"),
			},
			{
				Name:     "vpc_config_subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Configuration.VpcConfig.SubnetIds"),
			},
			{
				Name:     "vpc_config_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.VpcConfig.VpcId"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_lambda_function_file_system_configs",
				Resolver: fetchLambdaFunctionFileSystemConfigs,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "arn",
						Type: schema.TypeString,
					},
					{
						Name: "local_mount_path",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_lambda_function_layers",
				Resolver: fetchLambdaFunctionLayers,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "arn",
						Type: schema.TypeString,
					},
					{
						Name: "code_size",
						Type: schema.TypeBigInt,
					},
					{
						Name: "signing_job_arn",
						Type: schema.TypeString,
					},
					{
						Name: "signing_profile_version_arn",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_lambda_function_aliases",
				Resolver: fetchLambdaFunctionAliases,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "alias_arn",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "function_version",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "revision_id",
						Type: schema.TypeString,
					},
					{
						Name:     "routing_config_additional_version_weights",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("RoutingConfig.AdditionalVersionWeights"),
					},
				},
			},
			{
				Name:     "aws_lambda_function_event_invoke_configs",
				Resolver: fetchLambdaFunctionEventInvokeConfigs,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "on_failure_destination",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DestinationConfig.OnFailure.Destination"),
					},
					{
						Name:     "on_success_destination",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DestinationConfig.OnSuccess.Destination"),
					},
					{
						Name: "function_arn",
						Type: schema.TypeString,
					},
					{
						Name: "last_modified",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "maximum_event_age_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "maximum_retry_attempts",
						Type: schema.TypeInt,
					},
				},
			},
			{
				Name:     "aws_lambda_function_versions",
				Resolver: fetchLambdaFunctionVersions,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "code_sha256",
						Type: schema.TypeString,
					},
					{
						Name: "code_size",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "dead_letter_config_target_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DeadLetterConfig.TargetArn"),
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name:     "environment_error_error_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Environment.Error.ErrorCode"),
					},
					{
						Name:     "environment_error_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Environment.Error.Message"),
					},
					{
						Name:     "environment_variables",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Environment.Variables"),
					},
					{
						Name: "function_arn",
						Type: schema.TypeString,
					},
					{
						Name: "function_name",
						Type: schema.TypeString,
					},
					{
						Name: "handler",
						Type: schema.TypeString,
					},
					{
						Name:     "error_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ImageConfigResponse.Error.ErrorCode"),
					},
					{
						Name:     "error_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ImageConfigResponse.Error.Message"),
					},
					{
						Name:     "image_config_command",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("ImageConfigResponse.ImageConfig.Command"),
					},
					{
						Name:     "image_config_entry_point",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("ImageConfigResponse.ImageConfig.EntryPoint"),
					},
					{
						Name:     "image_config_working_directory",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ImageConfigResponse.ImageConfig.WorkingDirectory"),
					},
					{
						Name:     "kms_key_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KMSKeyArn"),
					},
					{
						Name: "last_modified",
						Type: schema.TypeString,
					},
					{
						Name: "last_update_status",
						Type: schema.TypeString,
					},
					{
						Name: "last_update_status_reason",
						Type: schema.TypeString,
					},
					{
						Name: "last_update_status_reason_code",
						Type: schema.TypeString,
					},
					{
						Name: "master_arn",
						Type: schema.TypeString,
					},
					{
						Name: "memory_size",
						Type: schema.TypeInt,
					},
					{
						Name: "package_type",
						Type: schema.TypeString,
					},
					{
						Name: "revision_id",
						Type: schema.TypeString,
					},
					{
						Name: "role",
						Type: schema.TypeString,
					},
					{
						Name: "runtime",
						Type: schema.TypeString,
					},
					{
						Name: "signing_job_arn",
						Type: schema.TypeString,
					},
					{
						Name: "signing_profile_version_arn",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "state_reason",
						Type: schema.TypeString,
					},
					{
						Name: "state_reason_code",
						Type: schema.TypeString,
					},
					{
						Name: "timeout",
						Type: schema.TypeInt,
					},
					{
						Name:     "tracing_config_mode",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TracingConfig.Mode"),
					},
					{
						Name: "version",
						Type: schema.TypeString,
					},
					{
						Name:     "vpc_config_security_group_ids",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("VpcConfig.SecurityGroupIds"),
					},
					{
						Name:     "vpc_config_subnet_ids",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("VpcConfig.SubnetIds"),
					},
					{
						Name:     "vpc_config_vpc_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VpcConfig.VpcId"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_lambda_function_version_file_system_configs",
						Resolver: fetchLambdaFunctionVersionFileSystemConfigs,
						Columns: []schema.Column{
							{
								Name:     "function_version_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "arn",
								Type: schema.TypeString,
							},
							{
								Name: "local_mount_path",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_lambda_function_version_layers",
						Resolver: fetchLambdaFunctionVersionLayers,
						Columns: []schema.Column{
							{
								Name:     "function_version_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "arn",
								Type: schema.TypeString,
							},
							{
								Name: "code_size",
								Type: schema.TypeBigInt,
							},
							{
								Name: "signing_job_arn",
								Type: schema.TypeString,
							},
							{
								Name: "signing_profile_version_arn",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_lambda_function_concurrency_configs",
				Resolver: fetchLambdaFunctionConcurrencyConfigs,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "allocated_provisioned_concurrent_executions",
						Type: schema.TypeInt,
					},
					{
						Name: "available_provisioned_concurrent_executions",
						Type: schema.TypeInt,
					},
					{
						Name: "function_arn",
						Type: schema.TypeString,
					},
					{
						Name: "last_modified",
						Type: schema.TypeString,
					},
					{
						Name: "requested_provisioned_concurrent_executions",
						Type: schema.TypeInt,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "status_reason",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_lambda_function_event_source_mappings",
				Resolver: fetchLambdaFunctionEventSourceMappings,
				Columns: []schema.Column{
					{
						Name:     "function_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "batch_size",
						Type: schema.TypeInt,
					},
					{
						Name: "bisect_batch_on_function_error",
						Type: schema.TypeBool,
					},
					{
						Name:     "on_failure_destination",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DestinationConfig.OnFailure.Destination"),
					},
					{
						Name:     "on_success_destination",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DestinationConfig.OnSuccess.Destination"),
					},
					{
						Name: "event_source_arn",
						Type: schema.TypeString,
					},
					{
						Name: "function_arn",
						Type: schema.TypeString,
					},
					{
						Name: "function_response_types",
						Type: schema.TypeStringArray,
					},
					{
						Name: "last_modified",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "last_processing_result",
						Type: schema.TypeString,
					},
					{
						Name: "maximum_batching_window_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "maximum_record_age_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "maximum_retry_attempts",
						Type: schema.TypeInt,
					},
					{
						Name: "parallelization_factor",
						Type: schema.TypeInt,
					},
					{
						Name: "queues",
						Type: schema.TypeStringArray,
					},
					{
						Name:     "self_managed_event_source_endpoints",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("SelfManagedEventSource.Endpoints"),
					},
					{
						Name: "starting_position",
						Type: schema.TypeString,
					},
					{
						Name: "starting_position_timestamp",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "state_transition_reason",
						Type: schema.TypeString,
					},
					{
						Name: "topics",
						Type: schema.TypeStringArray,
					},
					{
						Name: "tumbling_window_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name:     "uuid",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("UUID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_lambda_function_event_source_mapping_access_configurations",
						Resolver: fetchLambdaFunctionEventSourceMappingAccessConfigurations,
						Columns: []schema.Column{
							{
								Name:     "function_event_source_mapping_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "type",
								Type: schema.TypeString,
							},
							{
								Name:     "uri",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("URI"),
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var input lambda.ListFunctionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lambda
	for {
		response, err := svc.ListFunctions(ctx, &input, func(options *lambda.Options) {
			options.Region = c.Region
		})
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
	r, ok := resource.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Lambda

	response, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	var ae smithy.APIError
	if err != nil {
		if !errors.As(err, &ae) || ae.ErrorCode() != "ResourceNotFoundException" {
			return err
		}
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

	functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
		FunctionName: r.Configuration.FunctionName,
	}, func(options *lambda.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
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
	if err := resource.Set("code_signing_last_modified", codeSigningLastModified); err != nil {
		return err
	}

	return nil
}
func fetchLambdaFunctionFileSystemConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", r)
	}

	res <- r.Configuration.FileSystemConfigs
	return nil
}
func fetchLambdaFunctionLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", r)
	}

	res <- r.Configuration.Layers
	return nil
}
func fetchLambdaFunctionAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListAliasesInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListAliases(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Aliases
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListFunctionEventInvokeConfigs(ctx, &config)
		if err != nil {
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
func fetchLambdaFunctionVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListVersionsByFunctionInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListVersionsByFunction(ctx, &config)
		if err != nil {
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
func fetchLambdaFunctionVersionFileSystemConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.FunctionConfiguration)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of FunctionConfiguration", r)
	}

	res <- r.FileSystemConfigs
	return nil
}
func fetchLambdaFunctionVersionLayers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.FunctionConfiguration)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of FunctionConfiguration", r)
	}

	res <- r.Layers
	return nil
}
func fetchLambdaFunctionConcurrencyConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListProvisionedConcurrencyConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListProvisionedConcurrencyConfigs(ctx, &config)
		if err != nil {
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
func fetchLambdaFunctionEventSourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*lambda.GetFunctionOutput)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of *GetFunctionOutput", p)
	}
	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListEventSourceMappingsInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListEventSourceMappings(ctx, &config)
		if err != nil {
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
func fetchLambdaFunctionEventSourceMappingAccessConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(types.EventSourceMappingConfiguration)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of EventSourceMappingConfiguration", p)
	}
	res <- p.SourceAccessConfigurations
	return nil
}
