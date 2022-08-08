package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2NetworkInterfaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	niOutput := ec2.DescribeNetworkInterfacesOutput{}
	err := faker.FakeData(&niOutput)
	if err != nil {
		t.Fatal(err)
	}
	niOutput.NextToken = nil
	m.EXPECT().DescribeNetworkInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(&niOutput, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2NetworkInterfaces(t *testing.T) {
	client.AwsMockTestHelper(t, NetworkInterfaces(), buildEc2NetworkInterfaces, client.TestOptions{})
}
