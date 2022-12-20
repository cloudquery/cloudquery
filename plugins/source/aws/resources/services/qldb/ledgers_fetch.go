package qldb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
func fetchQldbLedgerJournalKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalKinesisStreamsForLedgerInput{
		LedgerName: ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().Qldb.ListJournalKinesisStreamsForLedger(ctx, config)
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

func fetchQldbLedgerJournalS3Exports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	config := &qldb.ListJournalS3ExportsForLedgerInput{
		Name:       ledger.Name,
		MaxResults: aws.Int32(100),
	}
	for {
		response, err := cl.Services().Qldb.ListJournalS3ExportsForLedger(ctx, config)
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
