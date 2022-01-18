//go:build mock
// +build mock

package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2TransitGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	tgw := ec2Types.TransitGateway{}
	err := faker.FakeData(&tgw)
	if err != nil {
		t.Fatal(err)
	}

	tgwvpca := ec2Types.TransitGatewayVpcAttachment{}
	err = faker.FakeData(&tgwvpca)
	if err != nil {
		t.Fatal(err)
	}

	tgwpeera := ec2Types.TransitGatewayPeeringAttachment{}
	err = faker.FakeData(&tgwpeera)
	if err != nil {
		t.Fatal(err)
	}

	tgwrt := ec2Types.TransitGatewayRouteTable{}
	err = faker.FakeData(&tgwrt)
	if err != nil {
		t.Fatal(err)
	}

	tgwmcd := ec2Types.TransitGatewayMulticastDomain{}
	err = faker.FakeData(&tgwmcd)
	if err != nil {
		t.Fatal(err)
	}

	tgwa := ec2Types.TransitGatewayAttachment{}
	err = faker.FakeData(&tgwa)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTransitGatewayVpcAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayVpcAttachmentsOutput{
			TransitGatewayVpcAttachments: []ec2Types.TransitGatewayVpcAttachment{tgwvpca},
		}, nil)

	m.EXPECT().DescribeTransitGatewayPeeringAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayPeeringAttachmentsOutput{
			TransitGatewayPeeringAttachments: []ec2Types.TransitGatewayPeeringAttachment{tgwpeera},
		}, nil)

	m.EXPECT().DescribeTransitGatewayRouteTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayRouteTablesOutput{
			TransitGatewayRouteTables: []ec2Types.TransitGatewayRouteTable{tgwrt},
		}, nil)

	m.EXPECT().DescribeTransitGatewayMulticastDomains(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayMulticastDomainsOutput{
			TransitGatewayMulticastDomains: []ec2Types.TransitGatewayMulticastDomain{tgwmcd},
		}, nil)
	m.EXPECT().DescribeTransitGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewayAttachmentsOutput{
			TransitGatewayAttachments: []ec2Types.TransitGatewayAttachment{tgwa},
		}, nil)
	m.EXPECT().DescribeTransitGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeTransitGatewaysOutput{
			TransitGateways: []ec2Types.TransitGateway{tgw},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2TransitGateways(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2TransitGateways(), buildEc2TransitGateways, client.TestOptions{})
}
