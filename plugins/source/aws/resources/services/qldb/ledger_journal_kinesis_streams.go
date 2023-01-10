package qldb

import (
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LedgerJournalKinesisStreams() *schema.Table {
	return &schema.Table{
		Name:        "aws_qldb_ledger_journal_kinesis_streams",
		Description: `https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalKinesisStreamDescription.html`,
		Resolver:    fetchQldbLedgerJournalKinesisStreams,
		Transform:   transformers.TransformWithStruct(&types.JournalKinesisStreamDescription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("qldb"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "ledger_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
