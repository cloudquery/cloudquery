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

func buildEc2VpcEndpointConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sc := types.VpcEndpointConnection{}
	require.NoError(t, faker.FakeObject(&sc))

	m.EXPECT().DescribeVpcEndpointConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointConnectionsOutput{
			VpcEndpointConnections: []types.VpcEndpointConnection{sc},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpointConnections(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpointConnections(), buildEc2VpcEndpointConnections, client.TestOptions{})
}
