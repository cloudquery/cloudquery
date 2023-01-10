package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var clusters docdb.DescribeDBClustersOutput
	if err := faker.FakeObject(&clusters); err != nil {
		t.Fatal(err)
	}
	clusters.Marker = nil
	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&clusters,
		nil,
	)

	var clusterSnapshots docdb.DescribeDBClusterSnapshotsOutput
	if err := faker.FakeObject(&clusterSnapshots); err != nil {
		t.Fatal(err)
	}
	clusterSnapshots.Marker = nil
	m.EXPECT().DescribeDBClusterSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&clusterSnapshots,
		nil,
	)

	var snapshotAttributes docdb.DescribeDBClusterSnapshotAttributesOutput
	if err := faker.FakeObject(&snapshotAttributes); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeDBClusterSnapshotAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&snapshotAttributes,
		nil,
	)

	var ev docdb.DescribeDBInstancesOutput
	if err := faker.FakeObject(&ev); err != nil {
		t.Fatal(err)
	}
	ev.Marker = nil
	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var tags docdb.ListTagsForResourceOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags,
		nil,
	).AnyTimes()

	return services
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildClustersMock, client.TestOptions{})
}
