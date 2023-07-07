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

func buildEc2VpcEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	e := types.VpcEndpoint{}
	require.NoError(t, faker.FakeObject(&e))

	m.EXPECT().DescribeVpcEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointsOutput{
			VpcEndpoints: []types.VpcEndpoint{e},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpoints(), buildEc2VpcEndpoints, client.TestOptions{})
}
