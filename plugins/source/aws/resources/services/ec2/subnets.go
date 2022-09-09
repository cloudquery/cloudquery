package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2Subnets() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_subnets",
		Description: "Describes a subnet.",
		Resolver:    fetchEc2Subnets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "assign_ipv6_address_on_creation",
				Description: "Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives an IPv6 address.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone of the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone_id",
				Description: "The AZ ID of the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "available_ip_address_count",
				Description: "The number of unused private IPv4 addresses in the subnet",
				Type:        schema.TypeInt,
			},
			{
				Name:        "cidr_block",
				Description: "The IPv4 CIDR block assigned to the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:          "customer_owned_ipv4_pool",
				Description:   "The customer-owned IPv4 address pool associated with the subnet.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "default_for_az",
				Description: "Indicates whether this is the default subnet for the Availability Zone.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "map_customer_owned_ip_on_launch",
				Description: "Indicates whether a network interface created in this subnet (including a network interface created by RunInstances) receives a customer-owned IPv4 address.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "map_public_ip_on_launch",
				Description: "Indicates whether instances launched in this subnet receive a public IPv4 address.",
				Type:        schema.TypeBool,
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the Amazon Web Services account that owns the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the subnet.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the subnet.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubnetArn"),
			},
			{
				Name:            "id",
				Description:     "The ID of the subnet.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("SubnetId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the subnet.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC the subnet is in.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_cidr_block_association_sets",
				Description: "Describes an IPv6 CIDR block associated with a subnet.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Ipv6CidrBlockAssociationSet"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2Subnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeSubnetsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeSubnets(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Subnets
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
