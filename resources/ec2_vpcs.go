package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
)

func Ec2Vpcs() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_vpcs",
		Resolver:     fetchEc2Vpcs,
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
				Name: "cidr_block",
				Type: schema.TypeString,
			},
			{
				Name: "dhcp_options_id",
				Type: schema.TypeString,
			},
			{
				Name: "instance_tenancy",
				Type: schema.TypeString,
			},
			{
				Name: "is_default",
				Type: schema.TypeBool,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2vpcTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_vpc_cidr_block_association_sets",
				Resolver: fetchEc2VpcCidrBlockAssociationSets,
				Columns: []schema.Column{
					{
						Name:     "vpc_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "association_id",
						Type: schema.TypeString,
					},
					{
						Name: "cidr_block",
						Type: schema.TypeString,
					},
					{
						Name:     "cidr_block_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CidrBlockState.State"),
					},
					{
						Name:     "cidr_block_state_status_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CidrBlockState.StatusMessage"),
					},
				},
			},
			{
				Name:     "aws_ec2_vpc_ipv6_cidr_block_association_sets",
				Resolver: fetchEc2VpcIpv6CidrBlockAssociationSets,
				Columns: []schema.Column{
					{
						Name:     "vpc_id",
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
					{
						Name: "ipv6_pool",
						Type: schema.TypeString,
					},
					{
						Name: "network_border_group",
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
	resource.Set("tags", tags)
	return nil
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
