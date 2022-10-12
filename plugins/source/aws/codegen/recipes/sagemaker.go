package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SagemakerResources() []*Resource {
	resources := []*Resource{

		{
			SubService:          "endpoint_configurations",
			Struct:              &sagemaker.DescribeEndpointConfigOutput{},
			SkipFields:          []string{"EndpointConfigArn"},
			PreResourceResolver: "getEndpointConfiguration",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("EndpointConfigArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The tags associated with the model.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveSagemakerEndpointConfigurationTags`,
					},
				}...),
		},

		{
			SubService:          "models",
			Struct:              &sagemaker.DescribeModelOutput{},
			SkipFields:          []string{"ModelArn"},
			PreResourceResolver: "getModel",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ModelArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The tags associated with the model.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveSagemakerModelTags`,
					},
				}...),
		},

		{
			SubService:          "notebook_instances",
			Struct:              &sagemaker.DescribeNotebookInstanceOutput{},
			SkipFields:          []string{"NotebookInstanceArn"},
			PreResourceResolver: "getNotebookInstance",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("NotebookInstanceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The tags associated with the notebook instance.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveSagemakerNotebookInstanceTags`,
					},
				}...),
		},

		{
			SubService:          "training_jobs",
			Struct:              &sagemaker.DescribeTrainingJobOutput{},
			SkipFields:          []string{"TrainingJobArn"},
			PreResourceResolver: "getTrainingJob",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TrainingJobArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:        "tags",
						Description: "The tags associated with the model.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveSagemakerTrainingJobTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "sagemaker"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("api.sagemaker")`
	}
	return resources
}
