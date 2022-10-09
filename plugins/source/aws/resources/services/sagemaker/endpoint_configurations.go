// Code generated by codegen; DO NOT EDIT.

package sagemaker

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EndpointConfigurations() *schema.Table {
	return &schema.Table{
		Name:                "aws_sagemaker_endpoint_configurations",
		Resolver:            fetchSagemakerEndpointConfigurations,
		PreResourceResolver: getEndpointConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer("api.sagemaker"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointConfigArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerEndpointConfigurationTags,
				Description: `The tags associated with the model.`,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "endpoint_config_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointConfigName"),
			},
			{
				Name:     "production_variants",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProductionVariants"),
			},
			{
				Name:     "async_inference_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AsyncInferenceConfig"),
			},
			{
				Name:     "data_capture_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataCaptureConfig"),
			},
			{
				Name:     "explainer_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExplainerConfig"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
