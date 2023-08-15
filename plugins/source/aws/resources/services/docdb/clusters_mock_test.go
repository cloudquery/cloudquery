package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var clusters docdb.DescribeDBClustersOutput
	require.NoError(t, faker.FakeObject(&clusters))

	clusters.Marker = nil
	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&clusters,
		nil,
	)

	var clusterSnapshots docdb.DescribeDBClusterSnapshotsOutput
	require.NoError(t, faker.FakeObject(&clusterSnapshots))

	clusterSnapshots.Marker = nil
	m.EXPECT().DescribeDBClusterSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&clusterSnapshots,
		nil,
	)

	var snapshotAttributes docdb.DescribeDBClusterSnapshotAttributesOutput
	require.NoError(t, faker.FakeObject(&snapshotAttributes))

	m.EXPECT().DescribeDBClusterSnapshotAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&snapshotAttributes,
		nil,
	)

	var ev docdb.DescribeDBInstancesOutput
	require.NoError(t, faker.FakeObject(&ev))

	ev.Marker = nil
	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var tags docdb.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags,
		nil,
	).AnyTimes()

	return services
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildClustersMock, client.TestOptions{})
}
