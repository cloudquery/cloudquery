package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEgressOnlyInternetGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	egressOutput := ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	err := faker.FakeData(&egressOutput)
	if err != nil {
		t.Fatal(err)
	}
	egressOutput.NextToken = nil
	m.EXPECT().DescribeEgressOnlyInternetGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(&egressOutput, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEgressOnlyInternetGateways(t *testing.T) {
	client.AwsMockTestHelper(t, EgressOnlyInternetGateways(), buildEgressOnlyInternetGateways, client.TestOptions{})
}
