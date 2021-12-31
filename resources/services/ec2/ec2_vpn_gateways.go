package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2VpnGateways() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpn_gateways",
		Resolver:     fetchEc2VpnGateways,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "amazon_side_asn",
				Type: schema.TypeBigInt,
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
				Resolver: resolveEc2VpnGatewayTags,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpnGatewayId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_vpc_attachment",
				Resolver: fetchEc2VpcAttachments,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"vpn_gateway_cq_id", "vpc_id"}},
				Columns: []schema.Column{
					{
						Name:     "vpn_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2VpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpnGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	output, err := svc.DescribeVpnGateways(ctx, &config, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- output.VpnGateways
	return nil
}
func fetchEc2VpcAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.VpnGateway)
	res <- r.VpcAttachments
	return nil
}
func resolveEc2VpnGatewayTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.VpnGateway)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
