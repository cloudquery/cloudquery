package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ledgerJournalKinesisStreams() *schema.Table {
	tableName := "aws_qldb_ledger_journal_kinesis_streams"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalKinesisStreamDescription.html`,
		Resolver:    fetchQldbLedgerJournalKinesisStreams,
		Transform:   transformers.TransformWithStruct(&types.JournalKinesisStreamDescription{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "qldb"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "ledger_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchQldbLedgerJournalKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	svc := meta.(*client.Client).Services().Qldb
	config := &qldb.ListJournalKinesisStreamsForLedgerInput{
		LedgerName: ledger.Name,
		MaxResults: aws.Int32(100),
	}
	paginator := qldb.NewListJournalKinesisStreamsForLedgerPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- page.Streams
	}
	return nil
}
