package ec2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func transitGatewayRouteTables() *schema.Table {
	tableName := "aws_ec2_transit_gateway_route_tables"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteTable.html`,
		Resolver:    fetchEc2TransitGatewayRouteTables,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.TransitGatewayRouteTable{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "transit_gateway_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2TransitGatewayRouteTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	paginator := ec2.NewDescribeTransitGatewayRouteTablesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TransitGatewayRouteTables
	}
	return nil
}
