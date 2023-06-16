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

func SubnetGroups() *schema.Table {
	tableName := "aws_docdb_subnet_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBSubnetGroup.html`,
		Resolver:    fetchDocdbSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBSubnetGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDBSubnetGroupTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("DBSubnetGroupArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchDocdbSubnetGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Docdb

	input := &docdb.DescribeDBSubnetGroupsInput{}

	p := docdb.NewDescribeDBSubnetGroupsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *docdb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DBSubnetGroups
	}
	return nil
}

func resolveDBSubnetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.DBSubnetGroup)
	return resolveDocDBTags(ctx, meta, resource, *item.DBSubnetGroupArn, c.Name)
}
