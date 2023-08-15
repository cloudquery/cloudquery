package docdb

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ClusterParameterGroups() *schema.Table {
	tableName := "aws_docdb_cluster_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html`,
		Resolver:    fetchDocdbClusterParameterGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBClusterParameterGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDBClusterParameterGroupTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DBClusterParameterGroupArn"),
				PrimaryKey: true,
			},
			{
				Name:     "parameters",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDocdbClusterParameterGroupParameters,
			},
			{
				Name:     "db_cluster_parameter_group_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DBClusterParameterGroupName"),
			},
			{
				Name:     "db_parameter_group_family",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DBParameterGroupFamily"),
			},
		},
	}
}

func fetchDocdbClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Docdb

	input := &docdb.DescribeDBClusterParameterGroupsInput{}

	p := docdb.NewDescribeDBClusterParameterGroupsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *docdb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBClusterParameterGroups
	}
	return nil
}

func resolveDocdbClusterParameterGroupParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterParameterGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Docdb

	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}

	var params []types.Parameter
	p := docdb.NewDescribeDBClusterParametersPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *docdb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		params = append(params, response.Parameters...)
	}
	return resource.Set(c.Name, params)
}

func resolveDBClusterParameterGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBClusterParameterGroup)
	return resolveDocDBTags(ctx, meta, resource, *item.DBClusterParameterGroupArn, c.Name)
}
