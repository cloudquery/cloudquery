// +build mock

package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRdsCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.Certificate{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeCertificatesOutput{
			Certificates: []rdsTypes.Certificate{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBCluster{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBClustersOutput{
			DBClusters: []rdsTypes.DBCluster{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBInstance{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBInstancesOutput{
			DBInstances: []rdsTypes.DBInstance{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func buildRdsDBSubnetGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBSubnetGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rds.DescribeDBSubnetGroupsOutput{
			DBSubnetGroups: []rdsTypes.DBSubnetGroup{l},
		}, nil)
	return client.Services{
		RDS: m,
	}
}

func TestRdsCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, RdsCertificates(), buildRdsCertificates, client.TestOptions{})
}
func TestRdsInstances(t *testing.T) {
	client.AwsMockTestHelper(t, RdsInstances(), buildRdsDBInstances, client.TestOptions{})
}
func TestRdsClusters(t *testing.T) {
	client.AwsMockTestHelper(t, RdsClusters(), buildRdsDBClusters, client.TestOptions{})
}
func TestRdsSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RdsSubnetGroups(), buildRdsDBSubnetGroups, client.TestOptions{})
}
