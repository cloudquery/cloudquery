package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRDSClient(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBSnapshot
	if err := faker.FakeData(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshots(
		gomock.Any(),
		&rds.DescribeDBSnapshotsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotsOutput{DBSnapshots: []types.DBSnapshot{s}},
		nil,
	)

	var attrs []types.DBSnapshotAttribute
	if err := faker.FakeData(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotAttributesOutput{
			DBSnapshotAttributesResult: &types.DBSnapshotAttributesResult{DBSnapshotAttributes: attrs},
		},
		nil,
	)
	return client.Services{RDS: mock}
}

func TestRDSDBSnapshots(t *testing.T) {
	awsTestHelper(t, RdsDbSnapshots(), buildRDSClient, TestOptions{})
}
