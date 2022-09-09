package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2NetworkAcls() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_network_acls",
		Description: "Describes a network ACL.",
		Resolver:    fetchEc2NetworkAcls,
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
					return []string{"network-acl", *resource.Item.(types.NetworkAcl).NetworkAclId}, nil
				}),
			},
			{
				Name:        "is_default",
				Description: "Indicates whether this is the default network ACL for the VPC.",
				Type:        schema.TypeBool,
			},
			{
				Name:            "id",
				Description:     "The ID of the network ACL.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("NetworkAclId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the network ACL.",
				Type:        schema.TypeString,
			},
			{
				Name:        "associations",
				Description: "Describes an association between a network ACL and a subnet.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Associations"),
			},
			{
				Name:        "entries",
				Description: "Describes an entry in a network ACL.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Entries"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2NetworkAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeNetworkAclsInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeNetworkAcls(ctx, &config)
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
