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

func buildEc2TransitGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	tgw := types.TransitGateway{}
	require.NoError(t, faker.FakeObject(&tgw))

	tgwvpca := types.TransitGatewayVpcAttachment{}
	require.NoError(t, faker.FakeObject(&tgwvpca))

	tgwpeera := types.TransitGatewayPeeringAttachment{}
	require.NoError(t, faker.FakeObject(&tgwpeera))

	tgwrt := types.TransitGatewayRouteTable{}
	require.NoError(t, faker.FakeObject(&tgwrt))

	tgwmcd := types.TransitGatewayMulticastDomain{}
	require.NoError(t, faker.FakeObject(&tgwmcd))

	tgwa := types.TransitGatewayAttachment{}
	require.NoError(t, faker.FakeObject(&tgwa))

	m.EXPECT().DescribeTransitGatewayVpcAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayVpcAttachmentsOutput{
			TransitGatewayVpcAttachments: []types.TransitGatewayVpcAttachment{tgwvpca},
		}, nil)

	m.EXPECT().DescribeTransitGatewayPeeringAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayPeeringAttachmentsOutput{
			TransitGatewayPeeringAttachments: []types.TransitGatewayPeeringAttachment{tgwpeera},
		}, nil)

	m.EXPECT().DescribeTransitGatewayRouteTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayRouteTablesOutput{
			TransitGatewayRouteTables: []types.TransitGatewayRouteTable{tgwrt},
		}, nil)

	m.EXPECT().DescribeTransitGatewayMulticastDomains(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayMulticastDomainsOutput{
			TransitGatewayMulticastDomains: []types.TransitGatewayMulticastDomain{tgwmcd},
		}, nil)
	m.EXPECT().DescribeTransitGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayAttachmentsOutput{
			TransitGatewayAttachments: []types.TransitGatewayAttachment{tgwa},
		}, nil)
	m.EXPECT().DescribeTransitGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewaysOutput{
			TransitGateways: []types.TransitGateway{tgw},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2TransitGateways(t *testing.T) {
	client.AwsMockTestHelper(t, TransitGateways(), buildEc2TransitGateways, client.TestOptions{})
}
