package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Vpcs() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpcs",
		Description:  "Describes a VPC.",
		Resolver:     fetchEc2Vpcs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Resolver:    resolveEc2vpcTags,
			},
			{
				Name:        "id",
				Description: "The ID of the VPC.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VpcId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_vpc_cidr_block_association_sets",
				Description: "Describes an IPv4 CIDR block associated with a VPC.",
				Resolver:    fetchEc2VpcCidrBlockAssociationSets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vpc_cq_id", "association_id"}},
				Columns: []schema.Column{
					{
						Name:        "vpc_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_vpcs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "association_id",
						Description: "The association ID for the IPv4 CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cidr_block",
						Description: "The IPv4 CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cidr_block_state",
						Description: "The state of the CIDR block.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CidrBlockState.State"),
					},
					{
						Name:        "cidr_block_state_status_message",
						Description: "A message about the status of the CIDR block, if applicable.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CidrBlockState.StatusMessage"),
					},
				},
			},
			{
				Name:        "aws_ec2_vpc_ipv6_cidr_block_association_sets",
				Description: "Describes an IPv6 CIDR block associated with a VPC.",
				Resolver:    fetchEc2VpcIpv6CidrBlockAssociationSets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vpc_cq_id", "association_id"}},
				Columns: []schema.Column{
					{
						Name:        "vpc_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_vpcs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "association_id",
						Description: "The association ID for the IPv6 CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv6_cidr_block",
						Description: "The IPv6 CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv6_cidr_block_state",
						Description: "The state of the CIDR block.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ipv6CidrBlockState.State"),
					},
					{
						Name:        "ipv6_cidr_block_state_status_message",
						Description: "A message about the status of the CIDR block, if applicable.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ipv6CidrBlockState.StatusMessage"),
					},
					{
						Name:        "ipv6_pool",
						Description: "The ID of the IPv6 address pool from which the IPv6 CIDR block is allocated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_border_group",
						Description: "The name of the unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses, for example, us-east-1-wl1-bos-wlz-1.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2Vpcs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeVpcsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeVpcs(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
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
func resolveEc2vpcTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Vpc)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2VpcCidrBlockAssociationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Vpc)
	res <- r.CidrBlockAssociationSet
	return nil
}
func fetchEc2VpcIpv6CidrBlockAssociationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Vpc)
	res <- r.Ipv6CidrBlockAssociationSet
	return nil
}
