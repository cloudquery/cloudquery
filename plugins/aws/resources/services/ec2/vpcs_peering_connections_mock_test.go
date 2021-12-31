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

func buildEc2VpcsPeeringConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.VpcPeeringConnection{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcPeeringConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcPeeringConnectionsOutput{
			VpcPeeringConnections: []ec2Types.VpcPeeringConnection{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2VpcsPeeringConnections(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2VpcPeeringConnections(), buildEc2VpcsPeeringConnections, client.TestOptions{})
}
