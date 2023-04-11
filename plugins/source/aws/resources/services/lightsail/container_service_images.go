package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func containerServiceImages() *schema.Table {
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

func fetchLightsailContainerServiceImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.ContainerService)
	input := lightsail.GetContainerImagesInput{
		ServiceName: r.ContainerServiceName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	deployments, err := svc.GetContainerImages(ctx, &input)
	if err != nil {
		return err
	}
	res <- deployments.ContainerImages
	return nil
}
