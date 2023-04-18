package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
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

	var eacc types.EndpointAccess
	if err := faker.FakeObject(&eacc); err != nil {
		t.Fatal(err)
	}
	eacc.ClusterIdentifier = g.ClusterIdentifier

	m.EXPECT().DescribeEndpointAccess(
		gomock.Any(),
		&redshift.DescribeEndpointAccessInput{
			ClusterIdentifier: g.ClusterIdentifier,
			MaxRecords:        aws.Int32(100),
		},
		gomock.Any(),
	).Return(
		&redshift.DescribeEndpointAccessOutput{EndpointAccessList: []types.EndpointAccess{eacc}},
		nil,
	)

	var eauth types.EndpointAuthorization
	if err := faker.FakeObject(&eauth); err != nil {
		t.Fatal(err)
	}
	eauth.ClusterIdentifier = g.ClusterIdentifier

	m.EXPECT().DescribeEndpointAuthorization(
		gomock.Any(),
		&redshift.DescribeEndpointAuthorizationInput{
			Account:           aws.String("testAccount"),
			ClusterIdentifier: g.ClusterIdentifier,
			MaxRecords:        aws.Int32(100),
		},
		gomock.Any(),
	).Return(
		&redshift.DescribeEndpointAuthorizationOutput{EndpointAuthorizationList: []types.EndpointAuthorization{eauth}},
		nil,
	)

	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildClustersMock, client.TestOptions{})
}
