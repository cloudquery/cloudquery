package rds

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsDbSecurityGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_db_security_groups",
		Description:  "Contains the details for an Amazon RDS DB security group",
		Resolver:     fetchRdsDbSecurityGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Resolver:    resolveRdsDbSecurityGroupJSONField(func(g types.DBSecurityGroup) interface{} { return g.EC2SecurityGroups }),
			},
			{
				Name:        "ip_ranges",
				Description: "Contains a list of IP range elements.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRdsDbSecurityGroupJSONField(func(g types.DBSecurityGroup) interface{} { return g.IPRanges }),
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
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsDbSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	var input rds.DescribeDBSecurityGroupsInput
	for {
		output, err := svc.DescribeDBSecurityGroups(ctx, &input, func(o *rds.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.DBSecurityGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRdsDbSecurityGroupJSONField(getter func(g types.DBSecurityGroup) interface{}) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		g := resource.Item.(types.DBSecurityGroup)
		b, err := json.Marshal(getter(g))
		if err != nil {
			return diag.WrapError(err)
		}
		return diag.WrapError(resource.Set(c.Name, b))
	}
}

func resolveRdsDbSecurityGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBSecurityGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().RDS
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBSecurityGroupArn}, func(o *rds.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.TagList)))
}
