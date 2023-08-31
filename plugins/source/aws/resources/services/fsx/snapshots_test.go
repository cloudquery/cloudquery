package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSnapshotsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var s types.Snapshot
	require.NoError(t, faker.FakeObject(&s))

	s.Lifecycle = types.SnapshotLifecycleAvailable
	m.EXPECT().DescribeSnapshots(
		gomock.Any(),
		&fsx.DescribeSnapshotsInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeSnapshotsOutput{Snapshots: []types.Snapshot{s}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, Snapshots(), buildSnapshotsMock, client.TestOptions{})
}
