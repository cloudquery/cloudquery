package lightsail

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v3/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	response, err := svc.GetContainerServices(ctx, &input, func(options *lightsail.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.ContainerServices
	return nil
}
