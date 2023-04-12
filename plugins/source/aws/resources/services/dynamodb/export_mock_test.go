package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDynamodbExportsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var es types.ExportSummary
	if err := faker.FakeObject(&es); err != nil {
		t.Fatal(err)
	}
	listOutput := &dynamodb.ListExportsOutput{
		ExportSummaries: []types.ExportSummary{es},
	}
	m.EXPECT().ListExports(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)

	var ed types.ExportDescription
	if err := faker.FakeObject(&ed); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeExport(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodb.DescribeExportOutput{
			ExportDescription: &ed,
		},
		nil,
	)

	return services
}

func TestDynamodbExports(t *testing.T) {
	client.AwsMockTestHelper(t, Exports(), buildDynamodbExportsMock, client.TestOptions{})
}
