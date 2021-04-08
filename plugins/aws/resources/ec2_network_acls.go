package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2NetworkAcls() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_network_acls",
		Resolver:     fetchEc2NetworkAcls,
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
				Name: "is_default",
				Type: schema.TypeBool,
			},
			{
				Name: "network_acl_id",
				Type: schema.TypeString,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2networkACLTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_network_acl_associations",
				Resolver: fetchEc2NetworkAclAssociations,
				Columns: []schema.Column{
					{
						Name:     "network_acl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "network_acl_association_id",
						Type: schema.TypeString,
					},
					{
						Name: "subnet_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_network_acl_entries",
				Resolver: fetchEc2NetworkAclEntries,
				Columns: []schema.Column{
					{
						Name:     "network_acl_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "cidr_block",
						Type: schema.TypeString,
					},
					{
						Name: "egress",
						Type: schema.TypeBool,
					},
					{
						Name:     "icmp_type_code",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("IcmpTypeCode.Code"),
					},
					{
						Name:     "icmp_type_code_type",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("IcmpTypeCode.Type"),
					},
					{
						Name: "ipv6_cidr_block",
						Type: schema.TypeString,
					},
					{
						Name:     "port_range_from",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("PortRange.From"),
					},
					{
						Name:     "port_range_to",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("PortRange.To"),
					},
					{
						Name: "protocol",
						Type: schema.TypeString,
					},
					{
						Name: "rule_action",
						Type: schema.TypeString,
					},
					{
						Name: "rule_number",
						Type: schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2NetworkAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config ec2.DescribeNetworkAclsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeNetworkAcls(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.NetworkAcls
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2networkACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.NetworkAcl)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchEc2NetworkAclAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.NetworkAcl)
	res <- r.Associations
	return nil
}
func fetchEc2NetworkAclEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r := parent.Item.(types.NetworkAcl)
	res <- r.Entries
	return nil
}
