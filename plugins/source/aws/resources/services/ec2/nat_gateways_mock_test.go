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

func buildEc2NatGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.NatGateway{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeNatGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNatGatewaysOutput{
			NatGateways: []types.NatGateway{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2NatGateways(t *testing.T) {
	client.AwsMockTestHelper(t, NatGateways(), buildEc2NatGateways, client.TestOptions{})
}
