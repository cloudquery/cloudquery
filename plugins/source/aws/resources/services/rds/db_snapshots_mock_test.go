package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRDSClient(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBSnapshot
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeDBSnapshots(
		gomock.Any(),
		&rds.DescribeDBSnapshotsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotsOutput{DBSnapshots: []types.DBSnapshot{s}},
		nil,
	)

	var attrs []types.DBSnapshotAttribute
	require.NoError(t, faker.FakeObject(&attrs))

	mock.EXPECT().DescribeDBSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotAttributesOutput{
			DBSnapshotAttributesResult: &types.DBSnapshotAttributesResult{DBSnapshotAttributes: attrs},
		},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSDBSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, DbSnapshots(), buildRDSClient, client.TestOptions{})
}
