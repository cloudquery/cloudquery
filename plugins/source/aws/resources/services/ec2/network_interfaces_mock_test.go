package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2NetworkInterfaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	niOutput := ec2.DescribeNetworkInterfacesOutput{}
	err := faker.FakeObject(&niOutput)
	if err != nil {
		t.Fatal(err)
	}
	niOutput.NextToken = nil
	m.EXPECT().DescribeNetworkInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(&niOutput, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2NetworkInterfaces(t *testing.T) {
	client.AwsMockTestHelper(t, NetworkInterfaces(), buildEc2NetworkInterfaces, client.TestOptions{})
}
