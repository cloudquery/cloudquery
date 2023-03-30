package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ContainerServices() *schema.Table {
	tableName := "aws_lightsail_container_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html`,
		Resolver:    fetchLightsailContainerServices,
		Transform:   transformers.TransformWithStruct(&types.ContainerService{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			ContainerServiceDeployments(),
			ContainerServiceImages(),
		},
	}
}
