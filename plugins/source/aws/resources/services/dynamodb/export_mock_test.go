package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDynamodbExportsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var es types.ExportSummary
	require.NoError(t, faker.FakeObject(&es))

	listOutput := &dynamodb.ListExportsOutput{
		ExportSummaries: []types.ExportSummary{es},
	}
	m.EXPECT().ListExports(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)

	var ed types.ExportDescription
	require.NoError(t, faker.FakeObject(&ed))

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
