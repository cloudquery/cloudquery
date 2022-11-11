package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingSheduledActionMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := client.Services{
		Autoscaling: m,
	}
	l := types.ScheduledUpdateGroupAction{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	autoscalingLaunchConfigurations := &autoscaling.DescribeScheduledActionsOutput{
		ScheduledUpdateGroupActions: []types.ScheduledUpdateGroupAction{l},
	}
	m.EXPECT().DescribeScheduledActions(gomock.Any(), gomock.Any(), gomock.Any()).Return(autoscalingLaunchConfigurations, nil)
	return services
}

func TestAutoscalingSheduledActions(t *testing.T) {
	client.AwsMockTestHelper(t, ScheduledActions(), buildAutoscalingSheduledActionMock, client.TestOptions{})
}
