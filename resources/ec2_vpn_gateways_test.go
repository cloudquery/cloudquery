package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2VpnGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.VpnGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpnGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpnGatewaysOutput{
			VpnGateways: []ec2Types.VpnGateway{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2VpnGateways(t *testing.T) {
	awsTestHelper(t, Ec2VpnGateways(), buildEc2VpnGateways, TestOptions{})
}
