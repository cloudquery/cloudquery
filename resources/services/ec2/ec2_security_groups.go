package ec2

import (
	"context"
	"fmt"

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
				Description: "The Amazon Resource Name (ARN) for the security group",
				Type:        schema.TypeString,
				Resolver:    resolveSGArn,
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
				Name:        "group_name",
				Description: "The name of the security group.",
				Type:        schema.TypeString,
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
				Name:        "vpc_id",
				Description: "The ID of the VPC for the security group.",
				Type:        schema.TypeString,
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
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ec2_security_group_ip_permission_ip_ranges",
						Description: "Describes an IPv4 range.",
						Resolver:    fetchEc2SecurityGroupIpPermissionIpRanges,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permission_cq_id", "cidr_ip"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permission_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cidr_ip",
								Description: "The IPv4 CIDR range.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this IPv4 address range.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_security_group_ip_permission_ipv6_ranges",
						Description: "[EC2-VPC only] Describes an IPv6 range.",
						Resolver:    fetchEc2SecurityGroupIpPermissionIpv6Ranges,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permission_cq_id", "cidr_ipv6"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permission_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cidr_ipv6",
								Description: "The IPv6 CIDR range.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this IPv6 address range.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_security_group_ip_permission_prefix_list_ids",
						Description: "Describes a prefix list ID.",
						Resolver:    fetchEc2SecurityGroupIpPermissionPrefixListIds,
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
								Name:        "group_name",
								Description: "The name of the security group.",
								Type:        schema.TypeString,
							},
							{
								Name:        "peering_status",
								Description: "The status of a VPC peering connection, if applicable.",
								Type:        schema.TypeString,
							},
							{
								Name:        "user_id",
								Description: "The ID of an AWS account.",
								Type:        schema.TypeString,
							},
							{
								Name:        "vpc_id",
								Description: "The ID of the VPC for the referenced security group, if applicable.",
								Type:        schema.TypeString,
							},
							{
								Name:        "vpc_peering_connection_id",
								Description: "The ID of the VPC peering connection, if applicable.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_ec2_security_group_ip_permissions_egresses",
				Description: "Describes a set of permissions for a security group rule.",
				Resolver:    fetchEc2SecurityGroupIpPermissionsEgresses,
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
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ec2_security_group_ip_permissions_egress_ip_ranges",
						Description: "Describes an IPv4 range.",
						Resolver:    fetchEc2SecurityGroupIpPermissionsEgressIpRanges,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permissions_egress_cq_id", "cidr_ip"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permissions_egress_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions_egresses table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cidr_ip",
								Description: "The IPv4 CIDR range.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this IPv4 address range.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_security_group_ip_permissions_egress_ipv6_ranges",
						Description: "[EC2-VPC only] Describes an IPv6 range.",
						Resolver:    fetchEc2SecurityGroupIpPermissionsEgressIpv6Ranges,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permissions_egress_cq_id", "cidr_ipv6"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permissions_egress_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions_egresses table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "cidr_ipv6",
								Description: "The IPv6 CIDR range.",
								Type:        schema.TypeString,
							},
							{
								Name:        "description",
								Description: "A description for the security group rule that references this IPv6 address range.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_security_group_ip_permissions_egress_prefix_list_ids",
						Description: "Describes a prefix list ID.",
						Resolver:    fetchEc2SecurityGroupIpPermissionsEgressPrefixListIds,
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permissions_egress_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions_egresses table (FK)",
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
						Name:        "aws_ec2_security_group_ip_permissions_egress_user_group_pairs",
						Description: "Describes a security group and AWS account ID pair.",
						Resolver:    fetchEc2SecurityGroupIpPermissionsEgressUserIdGroupPairs,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"security_group_ip_permissions_egress_cq_id", "group_id", "user_id"}},
						Columns: []schema.Column{
							{
								Name:        "security_group_ip_permissions_egress_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_security_group_ip_permissions_egresses table (FK)",
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
								Name:        "group_name",
								Description: "The name of the security group.",
								Type:        schema.TypeString,
							},
							{
								Name:        "peering_status",
								Description: "The status of a VPC peering connection, if applicable.",
								Type:        schema.TypeString,
							},
							{
								Name:        "user_id",
								Description: "The ID of an AWS account.",
								Type:        schema.TypeString,
							},
							{
								Name:        "vpc_id",
								Description: "The ID of the VPC for the referenced security group, if applicable.",
								Type:        schema.TypeString,
							},
							{
								Name:        "vpc_peering_connection_id",
								Description: "The ID of the VPC peering connection, if applicable.",
								Type:        schema.TypeString,
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
	res <- securityGroup.IpPermissions
	return nil
}
func fetchEc2SecurityGroupIpPermissionIpRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.IpRanges
	return nil
}
func fetchEc2SecurityGroupIpPermissionIpv6Ranges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.Ipv6Ranges
	return nil
}
func fetchEc2SecurityGroupIpPermissionPrefixListIds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.PrefixListIds
	return nil
}
func fetchEc2SecurityGroupIpPermissionUserIdGroupPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.UserIdGroupPairs
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroup, ok := parent.Item.(types.SecurityGroup)
	if !ok {
		return fmt.Errorf("not ec2 security group")
	}
	res <- securityGroup.IpPermissionsEgress
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressIpRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.IpRanges
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressIpv6Ranges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.Ipv6Ranges
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressPrefixListIds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.PrefixListIds
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressUserIdGroupPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.UserIdGroupPairs
	return nil
}

func resolveSGArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	sg, ok := resource.Item.(types.SecurityGroup)
	if !ok {
		return fmt.Errorf("not ec2 security group")
	}
	return resource.Set(c.Name, client.GenerateResourceARN("ec2", "security-group", *sg.GroupId, cl.Region, cl.AccountID))
}
