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

func ledgerJournalS3Exports() *schema.Table {
	tableName := "aws_qldb_ledger_journal_s3_exports"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalS3ExportDescription.html`,
		Resolver:    fetchQldbLedgerJournalS3Exports,
		Transform:   transformers.TransformWithStruct(&types.JournalS3ExportDescription{}),
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

func fetchQldbLedgerJournalS3Exports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	svc := meta.(*client.Client).Services().Qldb
	config := &qldb.ListJournalS3ExportsForLedgerInput{
		Name:       ledger.Name,
		MaxResults: aws.Int32(100),
	}
	paginator := qldb.NewListJournalS3ExportsForLedgerPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- page.JournalS3Exports
	}
	return nil
}
