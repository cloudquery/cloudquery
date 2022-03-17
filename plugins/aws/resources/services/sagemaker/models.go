package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"

	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SagemakerModels() *schema.Table {
	return &schema.Table{
		Name:          "aws_sagemaker_models",
		Description:   "Provides summary information about a model.",
		Resolver:      fetchSagemakerModels,
		Multiplex:     client.ServiceAccountRegionMultiplexer("api.sagemaker"),
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
				Name:        "enable_network_isolation",
				Description: "If True, no inbound or outbound network calls can be made to or from the model container.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "execution_role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role that you specified for the model.",
				Type:        schema.TypeString,
			},
			{
				Name:        "inference_execution_config",
				Description: "Specifies details of how containers in a multi-container endpoint are called.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "primary_container",
				Description: "The location of the primary inference code, associated artifacts, and custom environment map that the inference code uses when it is deployed in production.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the model.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerModelTags,
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that indicates when the model was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the model.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ModelArn"),
			},
			{
				Name:        "name",
				Description: "The name of the model that you want a summary for.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ModelName"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_sagemaker_model_containers",
				Description:   "Describes the container, as part of model definition.",
				Resolver:      fetchSagemakerModelContainers,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "model_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_model table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "container_hostname",
						Description: "This parameter is ignored for models that contain only a PrimaryContainer",
						Type:        schema.TypeString,
					},
					{
						Name:        "environment",
						Description: "The environment variables to set in the Docker container",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "image",
						Description: "The path where inference code is stored",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_config_repository_access_mode",
						Description: "Set this to one of the following values:  * Platform - The model image is hosted in Amazon ECR.  * Vpc - The model image is hosted in a private Docker registry in your VPC. ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageConfig.RepositoryAccessMode"),
					},
					{
						Name:        "image_config_repository_auth_config_repo_creds_provider_arn",
						Description: "The Amazon Resource Name (ARN) of an Amazon Web Services Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderArn"),
					},
					{
						Name:        "mode",
						Description: "Whether the container hosts a single model or multiple models.",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_data_url",
						Description: "The S3 path where the model artifacts, which result from model training, are stored",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_package_name",
						Description: "The name or Amazon Resource Name (ARN) of the model package to use to create the model.",
						Type:        schema.TypeString,
					},
					{
						Name:        "multi_model_config_model_cache_setting",
						Description: "Whether to cache models for a multi-model endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MultiModelConfig.ModelCacheSetting"),
					},
				},
			},
			{
				Name:          "aws_sagemaker_model_vpc_config",
				Description:   "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC",
				Resolver:      fetchSagemakerModelVpcConfigs,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "model_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_model table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "security_group_ids",
						Description: "The VPC security group IDs, in the form sg-xxxxxxxx",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "subnets",
						Description: "The ID of the subnets in the VPC to which you want to connect your training job or model",
						Type:        schema.TypeStringArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSagemakerModels(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListModelsInput{}
	for {
		response, err := svc.ListModels(ctx, &config, func(options *sagemaker.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		// get more details about the notebook instance
		for _, n := range response.Models {

			config := sagemaker.DescribeModelInput{
				ModelName: n.ModelName,
			}
			response, err := svc.DescribeModel(ctx, &config, func(options *sagemaker.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}

			model := WrappedSageMakerModel{
				DescribeModelOutput: response,
				ModelArn:            n.ModelArn,
				ModelName:           n.ModelName,
			}

			res <- model
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveSagemakerModelTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r, ok := resource.Item.(WrappedSageMakerModel)

	if !ok {
		return fmt.Errorf("expected ModelSummary but got %T", r)
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.ModelArn,
	}
	response, err := svc.ListTags(ctx, &config, func(options *sagemaker.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}

func fetchSagemakerModelContainers(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(WrappedSageMakerModel)
	if !ok {
		return fmt.Errorf("expected WrappedModel but got %T", r)
	}
	res <- r.Containers
	return nil
}

func fetchSagemakerModelVpcConfigs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(WrappedSageMakerModel)
	if !ok {
		return fmt.Errorf("expected WrappedModel but got %T", r)
	}
	res <- r.VpcConfig
	return nil
}

type WrappedSageMakerModel struct {
	*sagemaker.DescribeModelOutput
	ModelArn  *string
	ModelName *string
}
