// +build mock

package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAutoscalingLaunchConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := client.Services{
		Autoscaling: m,
	}
	l := types.LaunchConfiguration{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	autoscalingLaunchConfigurations := &autoscaling.DescribeLaunchConfigurationsOutput{
		LaunchConfigurations: []types.LaunchConfiguration{l},
	}
	m.EXPECT().DescribeLaunchConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(autoscalingLaunchConfigurations, nil)
	return services
}

func TestAutoscalingLaunchConfiguration(t *testing.T) {
	client.AwsMockTestHelper(t, AutoscalingLaunchConfigurations(), buildAutoscalingLaunchConfigurationsMock, client.TestOptions{})
}
