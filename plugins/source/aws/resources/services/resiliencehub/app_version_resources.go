package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func appVersionResources() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_app_version_resources",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_PhysicalResource.html`,
		Resolver:    fetchAppVersionResources,
		Transform:   transformers.TransformWithStruct(&types.PhysicalResource{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, appVersion,
			{
				Name:            "physical_resource_identifier",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("PhysicalResourceId.Identifier"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
