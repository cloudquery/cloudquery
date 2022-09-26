package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchQldbLedgers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().QLDB
	config := qldb.ListLedgersInput{}
	for {
		response, err := svc.ListLedgers(ctx, &config)
		if err != nil {
			return err
		}
		ledgers := make([]*qldb.DescribeLedgerOutput, 0, len(response.Ledgers))
		for _, l := range response.Ledgers {
			response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name}, func(o *qldb.Options) {
				o.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			ledgers = append(ledgers, response)
		}
		res <- ledgers
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveQldbLedgerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ledger := resource.Item.(*qldb.DescribeLedgerOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().QLDB
	response, err := svc.ListTagsForResource(ctx, &qldb.ListTagsForResourceInput{
		ResourceArn: ledger.Arn,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
func fetchQldbLedgerJournalKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalKinesisStreamsForLedgerInput{
		LedgerName: ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().QLDB.ListJournalKinesisStreamsForLedger(ctx, config)
		if err != nil {
			return err
		}

		res <- response.Streams
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchQldbLedgerJournalS3Exports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalS3ExportsForLedgerInput{
		Name:       ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().QLDB.ListJournalS3ExportsForLedger(ctx, config)
		if err != nil {
			return err
		}

		res <- response.JournalS3Exports
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
