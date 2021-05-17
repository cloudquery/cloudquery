package resources

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
		Resolver:     fetchEc2SecurityGroups,
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
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "group_id",
				Type: schema.TypeString,
			},
			{
				Name: "group_name",
				Type: schema.TypeString,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2securityGroupTags,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_security_group_ip_permissions",
				Resolver: fetchEc2SecurityGroupIpPermissions,
				Columns: []schema.Column{
					{
						Name:     "security_group_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "from_port",
						Type: schema.TypeInt,
					},
					{
						Name: "ip_protocol",
						Type: schema.TypeString,
					},
					{
						Name: "to_port",
						Type: schema.TypeInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_ec2_security_group_ip_permission_ip_ranges",
						Resolver: fetchEc2SecurityGroupIpPermissionIpRanges,
						Columns: []schema.Column{
							{
								Name:     "security_groupip_permission_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "cidr_ip",
								Type: schema.TypeString,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permission_ipv6_ranges",
						Resolver: fetchEc2SecurityGroupIpPermissionIpv6Ranges,
						Columns: []schema.Column{
							{
								Name:     "security_groupip_permission_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "cidr_ipv6",
								Type: schema.TypeString,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permission_prefix_list_ids",
						Resolver: fetchEc2SecurityGroupIpPermissionPrefixListIds,
						Columns: []schema.Column{
							{
								Name:     "security_groupip_permission_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
							{
								Name: "prefix_list_id",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permission_user_id_group_pairs",
						Resolver: fetchEc2SecurityGroupIpPermissionUserIdGroupPairs,
						Columns: []schema.Column{
							{
								Name:     "security_groupip_permission_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
							{
								Name: "group_id",
								Type: schema.TypeString,
							},
							{
								Name: "group_name",
								Type: schema.TypeString,
							},
							{
								Name: "peering_status",
								Type: schema.TypeString,
							},
							{
								Name: "user_id",
								Type: schema.TypeString,
							},
							{
								Name: "vpc_id",
								Type: schema.TypeString,
							},
							{
								Name: "vpc_peering_connection_id",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_ec2_security_group_ip_permissions_egresses",
				Resolver: fetchEc2SecurityGroupIpPermissionsEgresses,
				Columns: []schema.Column{
					{
						Name:     "security_group_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "from_port",
						Type: schema.TypeInt,
					},
					{
						Name: "ip_protocol",
						Type: schema.TypeString,
					},
					{
						Name: "to_port",
						Type: schema.TypeInt,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_ec2_security_group_ip_permissions_egress_ip_ranges",
						Resolver: fetchEc2SecurityGroupIpPermissionsEgressIpRanges,
						Columns: []schema.Column{
							{
								Name:     "security_group_ip_permissions_egress_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "cidr_ip",
								Type: schema.TypeString,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permissions_egress_ipv6_ranges",
						Resolver: fetchEc2SecurityGroupIpPermissionsEgressIpv6Ranges,
						Columns: []schema.Column{
							{
								Name:     "security_group_ip_permissions_egress_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "cidr_ipv6",
								Type: schema.TypeString,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permissions_egress_prefix_list_ids",
						Resolver: fetchEc2SecurityGroupIpPermissionsEgressPrefixListIds,
						Columns: []schema.Column{
							{
								Name:     "security_group_ip_permissions_egress_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
							{
								Name: "prefix_list_id",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_security_group_ip_permissions_egress_user_group_pairs",
						Resolver: fetchEc2SecurityGroupIpPermissionsEgressUserIdGroupPairs,
						Columns: []schema.Column{
							{
								Name:     "security_group_ip_permissions_egress_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "description",
								Type: schema.TypeString,
							},
							{
								Name: "group_id",
								Type: schema.TypeString,
							},
							{
								Name: "group_name",
								Type: schema.TypeString,
							},
							{
								Name: "peering_status",
								Type: schema.TypeString,
							},
							{
								Name: "user_id",
								Type: schema.TypeString,
							},
							{
								Name: "vpc_id",
								Type: schema.TypeString,
							},
							{
								Name: "vpc_peering_connection_id",
								Type: schema.TypeString,
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
func fetchEc2SecurityGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
func resolveEc2securityGroupTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.SecurityGroup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2SecurityGroupIpPermissions(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroup, ok := parent.Item.(types.SecurityGroup)
	if !ok {
		return fmt.Errorf("not ec2 security group")
	}
	res <- securityGroup.IpPermissions
	return nil
}
func fetchEc2SecurityGroupIpPermissionIpRanges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.IpRanges
	return nil
}
func fetchEc2SecurityGroupIpPermissionIpv6Ranges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.Ipv6Ranges
	return nil
}
func fetchEc2SecurityGroupIpPermissionPrefixListIds(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.PrefixListIds
	return nil
}
func fetchEc2SecurityGroupIpPermissionUserIdGroupPairs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermission, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission")
	}
	res <- securityGroupIpPermission.UserIdGroupPairs
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroup, ok := parent.Item.(types.SecurityGroup)
	if !ok {
		return fmt.Errorf("not ec2 security group")
	}
	res <- securityGroup.IpPermissionsEgress
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressIpRanges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.IpRanges
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressIpv6Ranges(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.Ipv6Ranges
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressPrefixListIds(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.PrefixListIds
	return nil
}
func fetchEc2SecurityGroupIpPermissionsEgressUserIdGroupPairs(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	securityGroupIpPermissionEgress, ok := parent.Item.(types.IpPermission)
	if !ok {
		return fmt.Errorf("not ec2 security group ip permission egress")
	}
	res <- securityGroupIpPermissionEgress.UserIdGroupPairs
	return nil
}
