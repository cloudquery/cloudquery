package qldb

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Ledgers() *schema.Table {
	tableName := "aws_qldb_ledgers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/qldb/latest/developerguide/API_DescribeLedger.html`,
		Resolver:            fetchQldbLedgers,
		PreResourceResolver: getLedger,
		Transform:           transformers.TransformWithStruct(&qldb.DescribeLedgerOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "qldb"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveQldbLedgerTags,
				Description: `The tags associated with the pipeline.`,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			ledgerJournalKinesisStreams(),
			ledgerJournalS3Exports(),
		},
	}
}

func fetchQldbLedgers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Qldb
	config := qldb.ListLedgersInput{}
	paginator := qldb.NewListLedgersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *qldb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Ledgers
	}
	return nil
}

func getLedger(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Qldb
	l := resource.Item.(types.LedgerSummary)

	response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name}, func(options *qldb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response
	return nil
}

func resolveQldbLedgerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ledger := resource.Item.(*qldb.DescribeLedgerOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().Qldb
	response, err := svc.ListTagsForResource(ctx, &qldb.ListTagsForResourceInput{
		ResourceArn: ledger.Arn,
	}, func(options *qldb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
