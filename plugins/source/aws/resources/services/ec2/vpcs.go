package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2Vpcs() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_vpcs",
		Description: "Describes a VPC.",
		Resolver:    fetchEc2Vpcs,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"vpc", *resource.Item.(types.Vpc).VpcId}, nil
				}),
			},
			{
				Name:        "cidr_block",
				Description: "The primary IPv4 CIDR block for the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dhcp_options_id",
				Description: "The ID of the set of DHCP options you've associated with the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_tenancy",
				Description: "The allowed tenancy of instances launched into the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_default",
				Description: "Indicates whether the VPC is the default VPC.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the VPC.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the VPC.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:            "id",
				Description:     "The ID of the VPC.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("VpcId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "cidr_block_association_set",
				Description: "Describes an IPv4 CIDR block associated with a VPC.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("CidrBlockAssociationSet"),
			},
			{
				Name:        "ipv6_cidr_block_association_set",
				Description: "Describes an IPv6 CIDR block associated with a VPC.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Ipv6CidrBlockAssociationSet"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2Vpcs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeVpcsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeVpcs(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Vpcs
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
