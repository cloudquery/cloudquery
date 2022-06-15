package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2VpcEndpointServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sd := ec2Types.ServiceDetail{}
	if err := faker.FakeData(&sd); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcEndpointServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServicesOutput{
			ServiceDetails: []ec2Types.ServiceDetail{sd},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2VpcEndpointServices(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2VpcEndpointServices(), buildEc2VpcEndpointServices, client.TestOptions{})
}
