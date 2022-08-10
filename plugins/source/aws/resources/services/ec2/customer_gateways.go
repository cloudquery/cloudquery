package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2CustomerGateways() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_customer_gateways",
		Description:   "Describes a customer gateway.",
		Resolver:      fetchEc2CustomerGateways,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "id",
				Description: "The ID of the customer gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CustomerGatewayId"),
			},
			{
				Name:        "bgp_asn",
				Description: "The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_arn",
				Description: "The Amazon Resource Name (ARN) for the customer gateway certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the customer gateway",
				Type:        schema.TypeString,
				Resolver:    resolveCustomerGatewayArn,
			},
			{
				Name:        "device_name",
				Description: "The name of customer gateway device.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ip_address",
				Description: "The Internet-routable IP address of the customer gateway's outside interface.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the customer gateway (pending | available | deleting | deleted).",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the customer gateway.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "type",
				Description: "The type of VPN connection the customer gateway supports (ipsec.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2CustomerGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	response, err := svc.DescribeCustomerGateways(ctx, nil, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- response.CustomerGateways
	return nil
}

func resolveCustomerGatewayArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cg := resource.Item.(types.CustomerGateway)
	return diag.WrapError(diag.WrapError(resource.Set(c.Name, cl.ARN(client.EC2Service, "customer-gateway", *cg.CustomerGatewayId))))
}
