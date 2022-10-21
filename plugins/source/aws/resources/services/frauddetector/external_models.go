// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ExternalModels() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_external_models",
		Description: "https://docs.aws.amazon.com/frauddetector/latest/api/API_ExternalModel.html",
		Resolver:    fetchFrauddetectorExternalModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "input_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InputConfiguration"),
			},
			{
				Name:     "invoke_model_endpoint_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvokeModelEndpointRoleArn"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "model_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelEndpoint"),
			},
			{
				Name:     "model_endpoint_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelEndpointStatus"),
			},
			{
				Name:     "model_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelSource"),
			},
			{
				Name:     "output_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutputConfiguration"),
			},
		},
	}
}
