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

func buildDynamodbBackupMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var bs types.BackupSummary
	require.NoError(t, faker.FakeObject(&bs))

	listOutput := &dynamodb.ListBackupsOutput{
		BackupSummaries: []types.BackupSummary{bs},
	}
	m.EXPECT().ListBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)

	var bd types.BackupDescription
	require.NoError(t, faker.FakeObject(&bd))

	m.EXPECT().DescribeBackup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodb.DescribeBackupOutput{
			BackupDescription: &bd,
		},
		nil,
	)

	return services
}

func TestDynamodbBackups(t *testing.T) {
	client.AwsMockTestHelper(t, Backups(), buildDynamodbBackupMock, client.TestOptions{})
}
