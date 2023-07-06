package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNeptuneClientForClusterSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)

	var s types.DBClusterSnapshot
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeDBClusterSnapshots(
		gomock.Any(),
		&neptune.DescribeDBClusterSnapshotsInput{
			Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
		},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBClusterSnapshotsOutput{DBClusterSnapshots: []types.DBClusterSnapshot{s}},
		nil,
	)

	var attrs []types.DBClusterSnapshotAttribute
	require.NoError(t, faker.FakeObject(&attrs))

	mock.EXPECT().DescribeDBClusterSnapshotAttributes(
		gomock.Any(),
		&neptune.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBClusterSnapshotAttributesOutput{
			DBClusterSnapshotAttributesResult: &types.DBClusterSnapshotAttributesResult{DBClusterSnapshotAttributes: attrs},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: s.DBClusterSnapshotArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return client.Services{Neptune: mock}
}

func TestNeptuneDBClusterSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterSnapshots(), buildNeptuneClientForClusterSnapshots, client.TestOptions{})
}
