package resources

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
		Resolver:     fetchEc2Subnets,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "assign_ipv6_address_on_creation",
				Type: schema.TypeBool,
			},
			{
				Name: "availability_zone",
				Type: schema.TypeString,
			},
			{
				Name: "availability_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "available_ip_address_count",
				Type: schema.TypeInt,
			},
			{
				Name: "cidr_block",
				Type: schema.TypeString,
			},
			{
				Name: "customer_owned_ipv4_pool",
				Type: schema.TypeString,
			},
			{
				Name: "default_for_az",
				Type: schema.TypeBool,
			},
			{
				Name: "map_customer_owned_ip_on_launch",
				Type: schema.TypeBool,
			},
			{
				Name: "map_public_ip_on_launch",
				Type: schema.TypeBool,
			},
			{
				Name: "outpost_arn",
				Type: schema.TypeString,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_arn",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2subnetTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_subnet_ipv6_cidr_block_association_sets",
				Resolver: fetchEc2SubnetIpv6CidrBlockAssociationSets,
				Columns: []schema.Column{
					{
						Name:     "subnet_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "association_id",
						Type: schema.TypeString,
					},
					{
						Name: "ipv6_cidr_block",
						Type: schema.TypeString,
					},
					{
						Name:     "ipv6_cidr_block_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ipv6CidrBlockState.State"),
					},
					{
						Name:     "ipv6_cidr_block_state_status_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ipv6CidrBlockState.StatusMessage"),
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
func resolveEc2subnetTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
