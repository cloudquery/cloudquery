// +build mock

package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2InternetGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.InternetGateway{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInternetGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInternetGatewaysOutput{
			InternetGateways: []types.InternetGateway{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2InternetGateways(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2InternetGateways(), buildEc2InternetGateways, client.TestOptions{})
}
