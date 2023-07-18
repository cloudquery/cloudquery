package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBackupJobsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var job types.BackupJob
	require.NoError(t, faker.FakeObject(&job))

	m.EXPECT().ListBackupJobs(
		gomock.Any(),
		&backup.ListBackupJobsInput{ByAccountId: aws.String("testAccount"), MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&backup.ListBackupJobsOutput{BackupJobs: []types.BackupJob{job}},
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestJobs(t *testing.T) {
	client.AwsMockTestHelper(t, Jobs(), buildBackupJobsMock, client.TestOptions{})
}
