package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApprunnerAutosScalingConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppRunnerClient(ctrl)
	as := types.AutoScalingConfiguration{}
	err := faker.FakeObject(&as)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAutoScalingConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListAutoScalingConfigurationsOutput{
			AutoScalingConfigurationSummaryList: []types.AutoScalingConfigurationSummary{
				{AutoScalingConfigurationArn: as.AutoScalingConfigurationArn},
			},
		}, nil)

	m.EXPECT().DescribeAutoScalingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.DescribeAutoScalingConfigurationOutput{
			AutoScalingConfiguration: &as,
		}, nil)

	return client.Services{
		Apprunner: m,
	}
}

func TestApprunnerAutosScalingConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, AutosScalingConfiguration(), buildApprunnerAutosScalingConfigurationsMock, client.TestOptions{})
}
