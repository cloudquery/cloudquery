package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2EbsSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	s := types.Snapshot{}
	userId := "userId"
	sa := types.CreateVolumePermission{
		Group:  "test",
		UserId: &userId,
	}
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSnapshotsOutput{
			Snapshots: []types.Snapshot{s},
		}, nil)
	m.EXPECT().DescribeSnapshotAttribute(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSnapshotAttributeOutput{
			CreateVolumePermissions: []types.CreateVolumePermission{sa},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2EbsSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, EbsSnapshots(), buildEc2EbsSnapshots, client.TestOptions{})
}
