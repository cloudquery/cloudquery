package ec2

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
		Description:  "Describes a network ACL.",
		Resolver:     fetchEc2NetworkAcls,
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"network-acl", *resource.Item.(types.NetworkAcl).NetworkAclId}, nil
				}),
			},
			{
				Name:        "is_default",
				Description: "Indicates whether this is the default network ACL for the VPC.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The ID of the network ACL.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkAclId"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the AWS account that owns the network ACL.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the network ACL.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2networkACLTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the network ACL.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_network_acl_associations",
				Description: "Describes an association between a network ACL and a subnet.",
				Resolver:    fetchEc2NetworkAclAssociations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"network_acl_cq_id", "network_acl_association_id"}},
				Columns: []schema.Column{
					{
						Name:        "network_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_network_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "network_acl_association_id",
						Description: "The ID of the association between a network ACL and a subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_ec2_network_acl_entries",
				Description: "Describes an entry in a network ACL.",
				Resolver:    fetchEc2NetworkAclEntries,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"network_acl_cq_id", "egress", "rule_number"}},
				Columns: []schema.Column{
					{
						Name:        "network_acl_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_network_acls table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cidr_block",
						Description: "The IPv4 network range to allow or deny, in CIDR notation.",
						Type:        schema.TypeString,
					},
					{
						Name:        "egress",
						Description: "Indicates whether the rule is an egress rule (applied to traffic leaving the subnet).",
						Type:        schema.TypeBool,
					},
					{
						Name:        "icmp_type_code",
						Description: "The ICMP code.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("IcmpTypeCode.Code"),
					},
					{
						Name:        "icmp_type_code_type",
						Description: "The ICMP type.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("IcmpTypeCode.Type"),
					},
					{
						Name:        "ipv6_cidr_block",
						Description: "The IPv6 network range to allow or deny, in CIDR notation.",
						Type:        schema.TypeString,
					},
					{
						Name:        "port_range_from",
						Description: "The first port in the range.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("PortRange.From"),
					},
					{
						Name:        "port_range_to",
						Description: "The last port in the range.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("PortRange.To"),
					},
					{
						Name:        "protocol",
						Description: "The protocol number.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_action",
						Description: "Indicates whether to allow or deny the traffic that matches the rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "rule_number",
						Description: "The rule number for the entry.",
						Type:        schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2NetworkAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
	return resource.Set("tags", tags)
}
func fetchEc2NetworkAclAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.NetworkAcl)
	res <- r.Associations
	return nil
}
func fetchEc2NetworkAclEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.NetworkAcl)
	res <- r.Entries
	return nil
}
