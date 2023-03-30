package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ContainerServiceImages() *schema.Table {
	tableName := "aws_lightsail_container_service_images"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerImage.html`,
		Resolver:    fetchLightsailContainerServiceImages,
		Transform:   transformers.TransformWithStruct(&types.ContainerImage{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "container_service_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
