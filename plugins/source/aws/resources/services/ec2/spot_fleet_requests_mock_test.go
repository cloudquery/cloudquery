package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSpotFleetRequests(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	item := types.SpotFleetRequestConfig{}
	require.NoError(t, faker.FakeObject(&item))

	ins := types.ActiveInstance{}
	require.NoError(t, faker.FakeObject(&ins))

	m.EXPECT().DescribeSpotFleetRequests(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSpotFleetRequestsOutput{
			SpotFleetRequestConfigs: []types.SpotFleetRequestConfig{item},
		}, nil)

	m.EXPECT().DescribeSpotFleetInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSpotFleetInstancesOutput{
			ActiveInstances: []types.ActiveInstance{ins},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestSpotFleetRequests(t *testing.T) {
	client.AwsMockTestHelper(t, SpotFleetRequests(), buildSpotFleetRequests, client.TestOptions{})
}
