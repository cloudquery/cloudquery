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

func buildEc2InstanceStatuses(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.InstanceStatus{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstanceStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceStatusOutput{
			InstanceStatuses: []ec2Types.InstanceStatus{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}
func TestEc2InstanceStatuses(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2InstanceStatuses(), buildEc2InstanceStatuses, client.TestOptions{})
}
