package docdb

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EventCategories() *schema.Table {
	tableName := "aws_docdb_event_categories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventCategoriesMap.html`,
		Resolver:    fetchDocdbEventCategories,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.EventCategoriesMap{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "event_categories",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: schema.PathResolver("EventCategories"),
			},
			{
				Name:     "source_type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("SourceType"),
			},
		},
	}
}

func fetchDocdbEventCategories(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Docdb

	input := &docdb.DescribeEventCategoriesInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"docdb"}}},
	}

	response, err := svc.DescribeEventCategories(ctx, input, func(options *docdb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.EventCategoriesMapList
	return nil
}
