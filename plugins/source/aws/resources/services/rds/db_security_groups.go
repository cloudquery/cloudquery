package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RdsDbSecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_db_security_groups",
		Description: "Contains the details for an Amazon RDS DB security group",
		Resolver:    fetchRdsDbSecurityGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
				Description: "The Amazon Resource Name (ARN) for the DB security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSecurityGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "description",
				Description: "Provides the description of the DB security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSecurityGroupDescription"),
			},
			{
				Name:        "name",
				Description: "Specifies the name of the DB security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DBSecurityGroupName"),
			},
			{
				Name:        "ec2_security_groups",
				Description: "Contains a list of EC2 Security Group elements.",
				Type:        schema.TypeJSON,
				Resolver: schema.PathResolver("DBSecurityGroup"),
			},
			{
				Name:        "owner_id",
				Description: "Provides the AWS ID of the owner of a specific DB security group.",
				Type:        schema.TypeString,
			},
			{
				Name:          "vpc_id",
				Description:   "Provides the VpcId of the DB security group.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "List of tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsDbSecurityGroupTags,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchRdsDbSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	var input rds.DescribeDBSecurityGroupsInput
	for {
		output, err := svc.DescribeDBSecurityGroups(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.DBSecurityGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}


func resolveRdsDbSecurityGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBSecurityGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBSecurityGroupArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
