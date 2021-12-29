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

func buildDirectconnectGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.DirectConnectGateway{}
	association := types.DirectConnectGatewayAssociation{}
	attachment := types.DirectConnectGatewayAttachment{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeData(&association)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeData(&attachment)
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

func TestAutoscalingLaunchConfiguration(t *testing.T) {
	client.AwsMockTestHelper(t, DirectconnectGateways(), buildDirectconnectGatewaysMock, client.TestOptions{})
}
