package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	redshiftTypes "github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRedshiftClustersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	g := redshiftTypes.Cluster{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClustersOutput{
			Clusters: []redshiftTypes.Cluster{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func buildRedshiftSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	g := redshiftTypes.ClusterSubnetGroup{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusterSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClusterSubnetGroupsOutput{
			ClusterSubnetGroups: []redshiftTypes.ClusterSubnetGroup{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftClusters(t *testing.T) {
	awsTestHelper(t, RedshiftClusters(), buildRedshiftClustersMock, TestOptions{})
}

func TestRedshiftSubnetGroups(t *testing.T) {
	awsTestHelper(t, RedshiftSubnetGroups(), buildRedshiftSubnetGroupsMock, TestOptions{})
}
