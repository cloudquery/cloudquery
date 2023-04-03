package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DbParameterGroups() *schema.Table {
	tableName := "aws_rds_db_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html`,
		Resolver:    fetchRdsDbParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsDbParameterGroupTags,
			},
		},

		Relations: []*schema.Table{
			DbParameterGroupDbParameters(),
		},
	}
}

func fetchRdsDbParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	var input rds.DescribeDBParameterGroupsInput
	for {
		output, err := svc.DescribeDBParameterGroups(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.DBParameterGroups
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func fetchRdsDbParameterGroupDbParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Rds
	g := parent.Item.(types.DBParameterGroup)
	input := rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName}
	for {
		output, err := svc.DescribeDBParameters(ctx, &input)
		if err != nil {
			if client.IsAWSError(err, "DBParameterGroupNotFound") {
				cl.Logger().Warn().Err(err).Msg("received DBParameterGroupNotFound on DescribeDBParameters")
				return nil
			}
			return err
		}
		res <- output.Parameters
		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
	}
	return nil
}

func resolveRdsDbParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBParameterGroup)
	svc := meta.(*client.Client).Services().Rds
	out, err := svc.ListTagsForResource(ctx, &rds.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
