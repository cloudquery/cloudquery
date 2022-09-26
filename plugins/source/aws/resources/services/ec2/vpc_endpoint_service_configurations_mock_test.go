package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2VpcEndpointServiceConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sc := ec2Types.ServiceConfiguration{}
	if err := faker.FakeData(&sc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcEndpointServiceConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServiceConfigurationsOutput{
			ServiceConfigurations: []ec2Types.ServiceConfiguration{sc},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2VpcEndpointServiceConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpointServiceConfigurations(), buildEc2VpcEndpointServiceConfigurations, client.TestOptions{})
}
