package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WrappedSageMakerModel struct {
	*sagemaker.DescribeModelOutput
	ModelArn  *string
	ModelName *string
}

func SagemakerModels() *schema.Table {
	return &schema.Table{
		Name:          "aws_sagemaker_models",
		Description:   "Provides summary information about a model.",
		Resolver:      fetchSagemakerModels,
		Multiplex:     client.ServiceAccountRegionMultiplexer("api.sagemaker"),
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
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of the model that you want a summary for.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ModelName"),
			},
			{
				Name:        "containers",
				Description: "Describes the container, as part of model definition.",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "vpc_config",
				Description: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC",
				Type: 			schema.TypeJSON,
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
		response, err := svc.ListModels(ctx, &config)
		if err != nil {
			return err
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
				return err
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
	r := resource.Item.(WrappedSageMakerModel)
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.ModelArn,
	}
	response, err := svc.ListTags(ctx, &config)
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return diag.WrapError(resource.Set("tags", tags))
}

