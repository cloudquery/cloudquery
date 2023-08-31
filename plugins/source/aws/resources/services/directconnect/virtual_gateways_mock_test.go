package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDirectconnectVirtualGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualGateway{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().DescribeVirtualGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualGatewaysOutput{
			VirtualGateways: []types.VirtualGateway{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}
func TestDirectconnectVirtualGateways(t *testing.T) {
	client.AwsMockTestHelper(t, VirtualGateways(), buildDirectconnectVirtualGatewaysMock, client.TestOptions{})
}
