package qldb

import (
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ledgers() *schema.Table {
	tableName := "aws_qldb_ledgers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/qldb/latest/developerguide/API_DescribeLedger.html`,
		Resolver:            fetchQldbLedgers,
		PreResourceResolver: getLedger,
		Transform:           client.TransformWithStruct(&qldb.DescribeLedgerOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "qldb"),
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
