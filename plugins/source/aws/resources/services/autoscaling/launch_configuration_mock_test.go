package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAutoscalingLaunchConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := client.Services{
		Autoscaling: m,
	}
	l := types.LaunchConfiguration{}
	require.NoError(t, faker.FakeObject(&l))
	autoscalingLaunchConfigurations := &autoscaling.DescribeLaunchConfigurationsOutput{
		LaunchConfigurations: []types.LaunchConfiguration{l},
	}
	m.EXPECT().DescribeLaunchConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(autoscalingLaunchConfigurations, nil)
	return services
}

func TestAutoscalingLaunchConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, LaunchConfigurations(), buildAutoscalingLaunchConfigurationsMock, client.TestOptions{})
}
