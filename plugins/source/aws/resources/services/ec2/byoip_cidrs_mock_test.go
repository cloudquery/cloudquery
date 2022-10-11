package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2ByoipCidrsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.ByoipCidr{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeByoipCidrs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeByoipCidrsOutput{
			ByoipCidrs: []types.ByoipCidr{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2ByoipCidrs(t *testing.T) {
	client.AwsMockTestHelper(t, ByoipCidrs(), buildEc2ByoipCidrsMock, client.TestOptions{})
}
