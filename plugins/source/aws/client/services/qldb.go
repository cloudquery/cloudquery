// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
)

//go:generate mockgen -package=mocks -destination=../mocks/qldb.go . QldbClient
type QldbClient interface {
	DescribeJournalKinesisStream(context.Context, *qldb.DescribeJournalKinesisStreamInput, ...func(*qldb.Options)) (*qldb.DescribeJournalKinesisStreamOutput, error)
	DescribeJournalS3Export(context.Context, *qldb.DescribeJournalS3ExportInput, ...func(*qldb.Options)) (*qldb.DescribeJournalS3ExportOutput, error)
	DescribeLedger(context.Context, *qldb.DescribeLedgerInput, ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error)
	GetBlock(context.Context, *qldb.GetBlockInput, ...func(*qldb.Options)) (*qldb.GetBlockOutput, error)
	GetDigest(context.Context, *qldb.GetDigestInput, ...func(*qldb.Options)) (*qldb.GetDigestOutput, error)
	GetRevision(context.Context, *qldb.GetRevisionInput, ...func(*qldb.Options)) (*qldb.GetRevisionOutput, error)
	ListJournalKinesisStreamsForLedger(context.Context, *qldb.ListJournalKinesisStreamsForLedgerInput, ...func(*qldb.Options)) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error)
	ListJournalS3Exports(context.Context, *qldb.ListJournalS3ExportsInput, ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsOutput, error)
	ListJournalS3ExportsForLedger(context.Context, *qldb.ListJournalS3ExportsForLedgerInput, ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsForLedgerOutput, error)
	ListLedgers(context.Context, *qldb.ListLedgersInput, ...func(*qldb.Options)) (*qldb.ListLedgersOutput, error)
	ListTagsForResource(context.Context, *qldb.ListTagsForResourceInput, ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error)
}
