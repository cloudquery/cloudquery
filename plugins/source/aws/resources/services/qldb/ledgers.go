package qldb

import (
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Ledgers() *schema.Table {
	return &schema.Table{
		Name:                "aws_qldb_ledgers",
		Description:         `https://docs.aws.amazon.com/qldb/latest/developerguide/API_DescribeLedger.html`,
		Resolver:            fetchQldbLedgers,
		PreResourceResolver: getLedger,
		Transform:           transformers.TransformWithStruct(&qldb.DescribeLedgerOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("qldb"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveQldbLedgerTags,
				Description: `The tags associated with the pipeline.`,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			LedgerJournalKinesisStreams(),
			LedgerJournalS3Exports(),
		},
	}
}
