package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRedshiftClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	g := types.Cluster{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	p := types.Parameter{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	logging := redshift.DescribeLoggingStatusOutput{}
	err = faker.FakeObject(&logging)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClustersOutput{
			Clusters: []types.Cluster{g},
		}, nil)
	m.EXPECT().DescribeClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClusterParametersOutput{
			Parameters: []types.Parameter{p},
		}, nil)
	m.EXPECT().DescribeLoggingStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&logging, nil)

	var snap types.Snapshot
	if err := faker.FakeObject(&snap); err != nil {
		t.Fatal(err)
	}
	snap.ClusterIdentifier = g.ClusterIdentifier
	snap.ClusterCreateTime = g.ClusterCreateTime
	m.EXPECT().DescribeClusterSnapshots(
		gomock.Any(),
		&redshift.DescribeClusterSnapshotsInput{
			ClusterExists:     aws.Bool(true),
			ClusterIdentifier: g.ClusterIdentifier,
			MaxRecords:        aws.Int32(100),
		},
		gomock.Any(),
	).Return(
		&redshift.DescribeClusterSnapshotsOutput{Snapshots: []types.Snapshot{snap}},
		nil,
	)

	return client.Services{
		Redshift: m,
	}
}

func buildRedshiftSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	g := types.ClusterSubnetGroup{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusterSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClusterSubnetGroupsOutput{
			ClusterSubnetGroups: []types.ClusterSubnetGroup{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildRedshiftClustersMock, client.TestOptions{})
}

func TestRedshiftSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildRedshiftSubnetGroupsMock, client.TestOptions{})
}
