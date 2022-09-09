package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2VpnGateways() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_vpn_gateways",
		Resolver:      fetchEc2VpnGateways,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpn-gateway", *resource.Item.(types.VpnGateway).VpnGatewayId}, nil
				}),
			},
			{
				Name: "amazon_side_asn",
				Type: schema.TypeInt,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("VpnGatewayId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "vpc_attachments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcAttachments"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2VpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpnGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	output, err := svc.DescribeVpnGateways(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.VpnGateways
	return nil
}
