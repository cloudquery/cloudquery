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

func buildSpotInstanceRequests(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	item := types.SpotInstanceRequest{}
	if err := faker.FakeObject(&item); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeSpotInstanceRequests(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSpotInstanceRequestsOutput{
			SpotInstanceRequests: []types.SpotInstanceRequest{item},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestSpotInstanceRequests(t *testing.T) {
	client.AwsMockTestHelper(t, SpotInstanceRequests(), buildSpotInstanceRequests, client.TestOptions{})
}
