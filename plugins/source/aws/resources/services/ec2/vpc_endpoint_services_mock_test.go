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

func buildEc2VpcEndpointServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sd := types.ServiceDetail{}
	if err := faker.FakeObject(&sd); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcEndpointServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServicesOutput{
			ServiceDetails: []types.ServiceDetail{sd},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpointServices(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpointServices(), buildEc2VpcEndpointServices, client.TestOptions{})
}
