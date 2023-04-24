package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/v2/transformers"
)

func gatewayAssociations() *schema.Table {
	tableName := "aws_directconnect_gateway_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAssociation.html`,
		Resolver:    fetchDirectconnectGatewayAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.DirectConnectGatewayAssociation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchDirectconnectGatewayAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	gateway := parent.Item.(types.DirectConnectGateway)
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	config := directconnect.DescribeDirectConnectGatewayAssociationsInput{DirectConnectGatewayId: gateway.DirectConnectGatewayId}
	// No paginator available
	for {
		output, err := svc.DescribeDirectConnectGatewayAssociations(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.DirectConnectGatewayAssociations
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
