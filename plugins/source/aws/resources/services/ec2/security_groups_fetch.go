package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_security_groups",
		Description: "Describes a security group .",
		Resolver:    fetchEc2SecurityGroups,
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
					return []string{"security-group", *resource.Item.(types.SecurityGroup).GroupId}, nil
				}),
			},
			{
				Name:        "description",
				Description: "A description of the security group.",
				Type:        schema.TypeString,
			},
			{
				Name:            "id",
				Description:     "The ID of the security group.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("GroupId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:          "vpc_id",
				Description:   "The ID of the VPC for the security group.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "ip_permissions",
				Description: "Describes a set of permissions for a security group rule.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("IpPermissions"),
			},
			{
				Name:        "ip_permissions_egress",
				Description: "Describes a set of egress permissions for a security group rule.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("IpPermissionsEgress"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2SecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeSecurityGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeSecurityGroups(ctx, &config, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.SecurityGroups
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveSecurityGroupArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.SecurityGroup)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "security_group/" + aws.ToString(item.GroupId),
	}
	return resource.Set(c.Name, a.String())
}
