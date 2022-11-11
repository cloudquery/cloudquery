package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2FlowLogsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := types.FlowLog{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeFlowLogs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeFlowLogsOutput{
			FlowLogs: []types.FlowLog{g},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}
func TestEc2FlowLogs(t *testing.T) {
	client.AwsMockTestHelper(t, FlowLogs(), buildEc2FlowLogsMock, client.TestOptions{})
}
