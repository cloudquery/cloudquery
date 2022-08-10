package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRDSClientForClusterSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBClusterSnapshot
	if err := faker.FakeData(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshots(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterSnapshotsOutput{DBClusterSnapshots: []types.DBClusterSnapshot{s}},
		nil,
	)

	var attrs []types.DBClusterSnapshotAttribute
	if err := faker.FakeData(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterSnapshotAttributesOutput{
			DBClusterSnapshotAttributesResult: &types.DBClusterSnapshotAttributesResult{DBClusterSnapshotAttributes: attrs},
		},
		nil,
	)
	return client.Services{RDS: mock}
}

func TestRDSDBClusterSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, RdsClusterSnapshots(), buildRDSClientForClusterSnapshots, client.TestOptions{})
}
