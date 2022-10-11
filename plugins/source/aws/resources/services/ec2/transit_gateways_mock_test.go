package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2TransitGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	tgw := types.TransitGateway{}
	err := faker.FakeData(&tgw)
	if err != nil {
		t.Fatal(err)
	}

	tgwvpca := types.TransitGatewayVpcAttachment{}
	err = faker.FakeData(&tgwvpca)
	if err != nil {
		t.Fatal(err)
	}

	tgwpeera := types.TransitGatewayPeeringAttachment{}
	err = faker.FakeData(&tgwpeera)
	if err != nil {
		t.Fatal(err)
	}

	tgwrt := types.TransitGatewayRouteTable{}
	err = faker.FakeData(&tgwrt)
	if err != nil {
		t.Fatal(err)
	}

	tgwmcd := types.TransitGatewayMulticastDomain{}
	err = faker.FakeData(&tgwmcd)
	if err != nil {
		t.Fatal(err)
	}

	tgwa := types.TransitGatewayAttachment{}
	err = faker.FakeData(&tgwa)
	if err != nil {
		t.Fatal(err)
	}

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
		EC2: m,
	}
}

func TestEc2TransitGateways(t *testing.T) {
	client.AwsMockTestHelper(t, TransitGateways(), buildEc2TransitGateways, client.TestOptions{})
}
