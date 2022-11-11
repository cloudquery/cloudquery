package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectVirtualGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualGateway{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualGatewaysOutput{
			VirtualGateways: []types.VirtualGateway{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}
func TestDirectconnecVirtualGateways(t *testing.T) {
	client.AwsMockTestHelper(t, VirtualGateways(), buildDirectconnectVirtualGatewaysMock, client.TestOptions{})
}
