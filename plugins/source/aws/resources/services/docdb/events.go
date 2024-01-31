package docdb

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Events() *schema.Table {
	tableName := "aws_docdb_events"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html`,
		Resolver:    fetchDocdbEvents,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform: transformers.TransformWithStruct(&types.Event{},
			transformers.WithPrimaryKeyComponents("SourceArn", "SourceIdentifier", "Date"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name: "categories_concat",
				Type: arrow.BinaryTypes.String,
				Resolver: func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, strings.Join(r.Item.(types.Event).EventCategories, ","))
				},
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDocdbEvents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDocdb).Docdb

	input := &docdb.DescribeEventsInput{}

	p := docdb.NewDescribeEventsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *docdb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Events
	}
	return nil
}
