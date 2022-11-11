package qldb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildLedgersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQldbClient(ctrl)

	ledger := types.LedgerSummary{}
	if err := faker.FakeObject(&ledger); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListLedgers(gomock.Any(), &qldb.ListLedgersInput{}, gomock.Any()).Return(
		&qldb.ListLedgersOutput{Ledgers: []types.LedgerSummary{ledger}}, nil)

	var resource qldb.DescribeLedgerOutput
	if err := faker.FakeObject(&resource); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLedger(
		gomock.Any(),
		&qldb.DescribeLedgerInput{Name: ledger.Name},
		gomock.Any(),
	).Return(
		&resource,
		nil,
	)

	tags := &qldb.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)

	s3 := types.JournalS3ExportDescription{}
	if err := faker.FakeObject(&s3); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListJournalS3ExportsForLedger(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&qldb.ListJournalS3ExportsForLedgerOutput{
			JournalS3Exports: []types.JournalS3ExportDescription{s3},
		}, nil)

	ke := types.JournalKinesisStreamDescription{}
	if err := faker.FakeObject(&ke); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListJournalKinesisStreamsForLedger(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&qldb.ListJournalKinesisStreamsForLedgerOutput{
			Streams: []types.JournalKinesisStreamDescription{ke},
		}, nil)

	return client.Services{Qldb: m}
}

func TestQldbLedgers(t *testing.T) {
	client.AwsMockTestHelper(t, Ledgers(), buildLedgersMock, client.TestOptions{})
}
