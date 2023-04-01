package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
			ledgerJournalKinesisStreams(),
			ledgerJournalS3Exports(),
		},
	}
}

func fetchQldbLedgers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Qldb
	config := qldb.ListLedgersInput{}
	for {
		response, err := svc.ListLedgers(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Ledgers

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getLedger(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Qldb
	l := resource.Item.(types.LedgerSummary)

	response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name})
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
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
