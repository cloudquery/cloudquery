package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSnapshotsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var s types.Snapshot
	require.NoError(t, faker.FakeDataSkipFields(&s, []string{
		"AdministrativeActions",
		"Lifecycle",
	}))
	s.Lifecycle = types.SnapshotLifecycleAvailable
	m.EXPECT().DescribeSnapshots(
		gomock.Any(),
		&fsx.DescribeSnapshotsInput{MaxResults: aws.Int32(1000)},
	).Return(
		&fsx.DescribeSnapshotsOutput{Snapshots: []types.Snapshot{s}},
		nil,
	)

	return client.Services{
		FSX: m,
	}
}

func TestSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, Snapshots(), buildSnapshotsMock, client.TestOptions{})
}
