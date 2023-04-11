package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func VirtualGateways() *schema.Table {
	tableName := "aws_directconnect_virtual_gateways"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html`,
		Resolver:    fetchDirectconnectVirtualGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.VirtualGateway{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchDirectconnectVirtualGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeVirtualGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualGateways(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.VirtualGateways
	return nil
}
