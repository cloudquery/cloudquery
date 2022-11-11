package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func QLDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "ledgers",
			Struct:              &qldb.DescribeLedgerOutput{},
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getLedger",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "tags",
						Description: "The tags associated with the pipeline.",
						Type:        schema.TypeJSON,
						Resolver:    `resolveQldbLedgerTags`,
					},
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"LedgerJournalKinesisStreams()",
				"LedgerJournalS3Exports()",
			},
		},
		{
			SubService:  "ledger_journal_kinesis_streams",
			Struct:      &types.JournalKinesisStreamDescription{},
			Description: "https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalKinesisStreamDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "ledger_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "ledger_journal_s3_exports",
			Struct:      &types.JournalS3ExportDescription{},
			Description: "https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalS3ExportDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "ledger_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "qldb"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("qldb")`
	}
	return resources
}
