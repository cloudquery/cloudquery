package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectVirtualGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualGateway{}
	err := faker.FakeData(&l)
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
	client.AwsMockTestHelper(t, DirectconnectVirtualGateways(), buildDirectconnectVirtualGatewaysMock, client.TestOptions{})
}
