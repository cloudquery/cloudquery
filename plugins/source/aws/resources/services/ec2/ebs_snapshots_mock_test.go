package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2EbsSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	sa := ec2.DescribeSnapshotAttributeOutput{}
	err := faker.FakeObject(&sa)
	if err != nil {
		t.Fatal(err)
	}

	s := types.Snapshot{}
	err = faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	s.OwnerId = aws.String("testAccount")
	m.EXPECT().DescribeSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSnapshotsOutput{
			Snapshots: []types.Snapshot{s},
		}, nil)

	m.EXPECT().DescribeSnapshotAttribute(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&sa, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2EbsSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, EbsSnapshots(), buildEc2EbsSnapshots, client.TestOptions{})
}
