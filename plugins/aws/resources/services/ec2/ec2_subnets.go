package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Subnets() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_subnets",
		Description:  "Describes a subnet.",
		Resolver:     fetchEc2Subnets,
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
				Name:        "customer_owned_ipv4_pool",
				Description: "The customer-owned IPv4 address pool associated with the subnet.",
				Type:        schema.TypeString,
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
				Name:        "outpost_arn",
				Description: "The Amazon Resource Name (ARN) of the Outpost.",
				Type:        schema.TypeString,
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
				Name:        "id",
				Description: "The ID of the subnet.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubnetId"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the subnet.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2SubnetsTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC the subnet is in.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_subnet_ipv6_cidr_block_association_sets",
				Description: "Describes an IPv6 CIDR block associated with a subnet.",
				Resolver:    fetchEc2SubnetIpv6CidrBlockAssociationSets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"subnet_cq_id", "ipv6_cidr_block"}},
				Columns: []schema.Column{
					{
						Name:        "subnet_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_subnets table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "association_id",
						Description: "The association ID for the CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv6_cidr_block",
						Description: "The IPv6 CIDR block.",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv6_cidr_block_state",
						Description: "The state of a CIDR block.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ipv6CidrBlockState.State"),
					},
					{
						Name:        "ipv6_cidr_block_state_status_message",
						Description: "A message about the status of the CIDR block, if applicable.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ipv6CidrBlockState.StatusMessage"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2Subnets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeSubnetsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeSubnets(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
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
func resolveEc2SubnetsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Subnet)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2SubnetIpv6CidrBlockAssociationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.Subnet)
	res <- r.Ipv6CidrBlockAssociationSet
	return nil
}
