package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SagemakerEndpointConfigurations() *schema.Table {
	return &schema.Table{
		Name:         "aws_sagemaker_endpoint_configurations",
		Description:  "Provides summary information for an endpoint configuration.",
		Resolver:     fetchSagemakerEndpointConfigurations,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "kms_key_id",
				Description: "Amazon Web Services KMS key ID Amazon SageMaker uses to encrypt data when storing it on the ML storage volume attached to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name: "data_capture_config",
				Type: schema.TypeJSON,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the model.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerEndpointConfigurationTags,
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that indicates when the endpoint configuration was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the endpoint configuration.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointConfigArn"),
			},
			{
				Name:        "name",
				Description: "Name of the Amazon SageMaker endpoint configuration.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EndpointConfigName"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_sagemaker_endpoint_configuration_production_variants",
				Description: "Identifies a model that you want to host and the resources chosen to deploy for hosting it",
				Resolver:    fetchSagemakerEndpointConfigurationProductionVariants,
				Columns: []schema.Column{
					{
						Name:        "endpoint_configuration_cq_id",
						Description: "Unique CloudQuery ID of aws_sagemaker_endpoint_configurations table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "initial_instance_count",
						Description: "Number of instances to launch initially.  This member is required.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "instance_type",
						Description: "The ML compute instance type.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_name",
						Description: "The name of the model that you want to host",
						Type:        schema.TypeString,
					},
					{
						Name:        "variant_name",
						Description: "The name of the production variant.  This member is required.",
						Type:        schema.TypeString,
					},
					{
						Name:        "accelerator_type",
						Description: "The size of the Elastic Inference (EI) instance to use for the production variant",
						Type:        schema.TypeString,
					},
					{
						Name:        "core_dump_config_destination_s3_uri",
						Description: "The Amazon S3 bucket to send the core dump to.  This member is required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CoreDumpConfig.DestinationS3Uri"),
					},
					{
						Name:        "core_dump_config_kms_key_id",
						Description: "The Amazon Web Services Key Management Service (Amazon Web Services KMS) key that Amazon SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CoreDumpConfig.KmsKeyId"),
					},
					{
						Name:        "initial_variant_weight",
						Description: "Determines initial traffic distribution among all of the models that you specify in the endpoint configuration",
						Type:        schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSagemakerEndpointConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListEndpointConfigsInput{}
	for {
		response, err := svc.ListEndpointConfigs(ctx, &config, func(options *sagemaker.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		// get more details about the notebook instance
		for _, n := range response.EndpointConfigs {

			config := sagemaker.DescribeEndpointConfigInput{
				EndpointConfigName: n.EndpointConfigName,
			}
			response, err := svc.DescribeEndpointConfig(ctx, &config, func(options *sagemaker.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			res <- response
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchSagemakerEndpointConfigurationProductionVariants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(*sagemaker.DescribeEndpointConfigOutput)
	if !ok {
		return fmt.Errorf("expected DescribeEndpointConfigOutput but got %T", r)
	}
	res <- r.ProductionVariants
	return nil
}
func resolveSagemakerEndpointConfigurationTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r, ok := resource.Item.(*sagemaker.DescribeEndpointConfigOutput)

	if !ok {
		return fmt.Errorf("expected DescribeEndpointConfigOutput but got %T", r)
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.EndpointConfigArn,
	}
	response, err := svc.ListTags(ctx, &config, func(options *sagemaker.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}
