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

func buildEc2RouteTables(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.RouteTable{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRouteTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeRouteTablesOutput{
			RouteTables: []types.RouteTable{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2RouteTables(t *testing.T) {
	client.AwsMockTestHelper(t, RouteTables(), buildEc2RouteTables, client.TestOptions{})
}
