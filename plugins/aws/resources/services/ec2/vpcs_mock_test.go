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

func buildEc2Vpcs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Vpc{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcsOutput{
			Vpcs: []ec2Types.Vpc{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}
func TestEc2Vpcs(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2Vpcs(), buildEc2Vpcs, client.TestOptions{})
}
