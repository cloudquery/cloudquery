package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
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
			containerServiceDeployments(),
			containerServiceImages(),
		},
	}
}

func fetchLightsailContainerServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetContainerServicesInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetContainerServices(ctx, &input)
	if err != nil {
		return err
	}
	res <- response.ContainerServices
	return nil
}
