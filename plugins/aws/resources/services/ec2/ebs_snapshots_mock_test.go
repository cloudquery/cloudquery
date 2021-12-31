// +build mock

package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2EbsSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	s := ec2Types.Snapshot{}
	userId := "userId"
	sa := ec2Types.CreateVolumePermission{
		Group:  "test",
		UserId: &userId,
	}
	err := faker.FakeData(&s)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSnapshotsOutput{
			Snapshots: []ec2Types.Snapshot{s},
		}, nil)
	m.EXPECT().DescribeSnapshotAttribute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSnapshotAttributeOutput{
			CreateVolumePermissions: []ec2Types.CreateVolumePermission{sa},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2EbsSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2EbsSnapshots(), buildEc2EbsSnapshots, client.TestOptions{})
}
