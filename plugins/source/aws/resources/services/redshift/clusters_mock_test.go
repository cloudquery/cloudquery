package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	g := types.Cluster{}
	require.NoError(t, faker.FakeObject(&g))
	p := types.Parameter{}
	require.NoError(t, faker.FakeObject(&p))
	logging := redshift.DescribeLoggingStatusOutput{}
	require.NoError(t, faker.FakeObject(&logging))

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
	require.NoError(t, faker.FakeObject(&snap))

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
	require.NoError(t, faker.FakeObject(&eacc))

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
	require.NoError(t, faker.FakeObject(&eauth))

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
