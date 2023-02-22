package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRdsCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.Certificate{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeCertificatesOutput{
			Certificates: []rdsTypes.Certificate{l},
		}, nil)
	return client.Services{
		Rds: m,
	}
}

func buildRdsDBClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBCluster{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBClustersOutput{
			DBClusters: []rdsTypes.DBCluster{l},
		}, nil)
	return client.Services{
		Rds: m,
	}
}

func buildRdsDBInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBInstance{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBInstancesOutput{
			DBInstances: []rdsTypes.DBInstance{l},
		}, nil)
	return client.Services{
		Rds: m,
	}
}

func buildRdsDBSubnetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBSubnetGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBSubnetGroupsOutput{
			DBSubnetGroups: []rdsTypes.DBSubnetGroup{l},
		}, nil)
	return client.Services{
		Rds: m,
	}
}

func buildRdsDBReservedInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	ri := rdsTypes.ReservedDBInstance{}
	if err := faker.FakeObject(&ri); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeReservedDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeReservedDBInstancesOutput{
			ReservedDBInstances: []rdsTypes.ReservedDBInstance{ri},
		}, nil)

	tagOutput := rds.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tagOutput); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagOutput, nil)

	return client.Services{
		Rds: m,
	}
}

func TestRdsCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildRdsCertificates, client.TestOptions{})
}
func TestRdsInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildRdsDBInstances, client.TestOptions{})
}
func TestRdsClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildRdsDBClusters, client.TestOptions{})
}
func TestRdsSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildRdsDBSubnetGroups, client.TestOptions{})
}

func TestRdsReservedInstances(t *testing.T) {
	client.AwsMockTestHelper(t, ReservedInstances(), buildRdsDBReservedInstances, client.TestOptions{})
}
