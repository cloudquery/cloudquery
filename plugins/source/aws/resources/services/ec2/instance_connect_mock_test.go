package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildInstanceConnectEndpoint(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	endpoint := types.Ec2InstanceConnectEndpoint{}
	err := faker.FakeObject(&endpoint)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstanceConnectEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceConnectEndpointsOutput{
			InstanceConnectEndpoints: []types.Ec2InstanceConnectEndpoint{endpoint},
			NextToken:                nil,
			ResultMetadata:           middleware.Metadata{},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestInstanceConnect(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceConnectEndpoints(), buildInstanceConnectEndpoint, client.TestOptions{})
}
