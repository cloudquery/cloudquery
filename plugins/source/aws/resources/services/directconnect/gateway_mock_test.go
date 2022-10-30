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

func buildDirectconnectGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.DirectConnectGateway{}
	association := types.DirectConnectGatewayAssociation{}
	attachment := types.DirectConnectGatewayAttachment{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&association)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&attachment)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeDirectConnectGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewaysOutput{
			DirectConnectGateways: []types.DirectConnectGateway{l},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAssociationsOutput{
			DirectConnectGatewayAssociations: []types.DirectConnectGatewayAssociation{association},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAttachmentsOutput{
			DirectConnectGatewayAttachments: []types.DirectConnectGatewayAttachment{attachment},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectGateways(t *testing.T) {
	client.AwsMockTestHelper(t, Gateways(), buildDirectconnectGatewaysMock, client.TestOptions{})
}
