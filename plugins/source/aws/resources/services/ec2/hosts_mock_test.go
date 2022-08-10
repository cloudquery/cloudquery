package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2Hosts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := types.Host{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeHosts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeHostsOutput{
			Hosts: []types.Host{g},
		}, nil)

	services := client.Services{
		EC2: m,
	}
	return services
}

func TestEc2Hosts(t *testing.T) {
	client.AwsMockTestHelper(t, Hosts(), buildEc2Hosts, client.TestOptions{})
}
