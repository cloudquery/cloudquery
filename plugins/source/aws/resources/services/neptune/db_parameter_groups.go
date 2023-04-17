package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func DbParameterGroups() *schema.Table {
	tableName := "aws_neptune_db_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBClusterParameterGroups`,
		Resolver:    fetchNeptuneDbParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
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
				Resolver: resolveNeptuneDbParameterGroupTags,
			},
		},

		Relations: []*schema.Table{
			DbParameterGroupDbParameters(),
		},
	}
}

func fetchNeptuneDbParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	input := neptune.DescribeDBParameterGroupsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
	paginator := neptune.NewDescribeDBParameterGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.DBParameterGroups
	}
	return nil
}

func fetchNeptuneDbParameterGroupDbParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	g := parent.Item.(types.DBParameterGroup)
	input := neptune.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName}
	paginator := neptune.NewDescribeDBParametersPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if client.IsAWSError(err, "DBParameterGroupNotFound") {
				cl.Logger().Warn().Err(err).Msg("received DBParameterGroupNotFound on DescribeDBParameters")
				return nil
			}
			return err
		}
		res <- page.Parameters
	}
	return nil
}

func resolveNeptuneDbParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	g := resource.Item.(types.DBParameterGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
