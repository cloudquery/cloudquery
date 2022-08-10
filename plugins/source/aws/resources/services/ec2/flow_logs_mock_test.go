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

func buildEc2FlowLogsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := ec2Types.FlowLog{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFlowLogs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeFlowLogsOutput{
			FlowLogs: []ec2Types.FlowLog{g},
		}, nil)
	return client.Services{
		EC2: m,
	}
}
func TestEc2FlowLogs(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2FlowLogs(), buildEc2FlowLogsMock, client.TestOptions{})
}
