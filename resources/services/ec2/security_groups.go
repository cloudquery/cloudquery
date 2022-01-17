package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_security_groups",
		Description:  "Describes a security group .",
		Resolver:     fetchEc2SecurityGroups,
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
					return []string{"security-group", *resource.Item.(types.SecurityGroup).GroupId}, nil
				}),
			},
			{
				Name:        "description",
				Description: "A description of the security group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GroupId"),
			},
			{
				Name:          "group_name",
				Description:   "The name of the security group.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the owner of the security group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the security group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2securityGroupTags,
			},
			{
				Name:          "vpc_id",
				Description:   "The ID of the VPC for the security group.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_security_group_ip_permissions",
				Description: "Describes a set of permissions for a security group rule.",
				Resolver:    fetchEc2SecurityGroupIpPermissions,
				Columns: []schema.Column{
					{
						Name:        "security_group_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_security_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "from_port",
						Description: "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "ip_protocol",
						Description: "The IP protocol name (tcp, udp, icmp, icmpv6) or number",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_port",
						Description: "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.",
						Type:        schema.TypeInt,
					}, {
						Name:        "permission_type",
						Description: "egress or ingress",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ec2_security_group_ip_permission_ip_ranges",
						Description: "Details of a cidr range associated with a security group rule",
						Resolver:    fetchEc2SecurityGroupIpPermissionIpRanges,
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permission_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cidr",
								Description: "The CIDR range.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this address range.",
								Type:        schema.TypeString,
							}, {
								Name:        "cidr_type",
								Description: "IP Type: ipv4, or ipv6",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ec2_security_group_ip_permission_prefix_list_ids",
						Description:   "Describes a prefix list ID.",
						Resolver:      fetchEc2SecurityGroupIpPermissionPrefixListIds,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permission_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this prefix list ID.",
								Type:        schema.TypeString,
							},
							{
								Name:        "prefix_list_id",
								Description: "The ID of the prefix.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_security_group_ip_permission_user_id_group_pairs",
						Description: "Describes a security group and AWS account ID pair.",
						Resolver:    fetchEc2SecurityGroupIpPermissionUserIdGroupPairs,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permission_cq_id", "group_id", "user_id"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permission_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this user ID group pair.",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_id",
								Description: "The ID of the security group.",
								Type:        schema.TypeString,
							},
							{
								Name:          "group_name",
								Description:   "The name of the security group.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "peering_status",
								Description:   "The status of a VPC peering connection, if applicable.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "user_id",
								Description: "The ID of an AWS account.",
								Type:        schema.TypeString,
							},
							{
								Name:          "vpc_id",
								Description:   "The ID of the VPC for the referenced security group, if applicable.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:          "vpc_peering_connection_id",
								Description:   "The ID of the VPC peering connection, if applicable.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2SecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeSecurityGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2

	response, err := svc.DescribeSecurityGroups(ctx, &config, func(o *ec2.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.SecurityGroups
	return nil
}
func resolveEc2securityGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.SecurityGroup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2SecurityGroupIpPermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	securityGroup, ok := parent.Item.(types.SecurityGroup)
	if !ok {
		return fmt.Errorf("not ec2 security group")
	}

	capacity := len(securityGroup.IpPermissionsEgress) + len(securityGroup.IpPermissions)

	ipRanges := make([]ipPermission, 0, capacity)
	for _, ip := range securityGroup.IpPermissionsEgress {
		ipRanges = append(ipRanges, ipPermission{ip, "egress"})
	}

	for _, ip := range securityGroup.IpPermissions {
		ipRanges = append(ipRanges, ipPermission{ip, "ingress"})
	}
	res <- ipRanges
	return nil
}

func fetchEc2SecurityGroupIpPermissionIpRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	securityGroupIpPermission, ok := parent.Item.(ipPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}

	type customIpRange struct {
		Cidr        string
		Description string
		CidrType    string
	}

	capacity := len(securityGroupIpPermission.IpRanges) + len(securityGroupIpPermission.Ipv6Ranges)

	ipRanges := make([]customIpRange, 0, capacity)
	for _, ip := range securityGroupIpPermission.IpRanges {
		ipRanges = append(ipRanges, customIpRange{aws.ToString(ip.CidrIp), aws.ToString(ip.Description), "ipv4"})
	}

	for _, ip := range securityGroupIpPermission.Ipv6Ranges {
		ipRanges = append(ipRanges, customIpRange{aws.ToString(ip.CidrIpv6), aws.ToString(ip.Description), "ipv6"})
	}
	res <- ipRanges
	return nil
}
func fetchEc2SecurityGroupIpPermissionPrefixListIds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(ipPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission in ip range")
	}
	res <- securityGroupIpPermission.PrefixListIds
	return nil
}
func fetchEc2SecurityGroupIpPermissionUserIdGroupPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(ipPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission in user id group pair")
	}
	res <- securityGroupIpPermission.UserIdGroupPairs
	return nil
}

type ipPermission struct {
	types.IpPermission
	PermissionType string
}
