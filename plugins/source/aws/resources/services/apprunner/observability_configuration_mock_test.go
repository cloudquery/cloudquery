package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildObservabilityConfiguration(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApprunnerClient(ctrl)
	s := types.ObservabilityConfiguration{}
	require.NoError(t, faker.FakeObject(&s))

	m.EXPECT().ListObservabilityConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListObservabilityConfigurationsOutput{
			ObservabilityConfigurationSummaryList: []types.ObservabilityConfigurationSummary{
				{ObservabilityConfigurationArn: s.ObservabilityConfigurationArn},
			},
		}, nil)

	m.EXPECT().DescribeObservabilityConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.DescribeObservabilityConfigurationOutput{
			ObservabilityConfiguration: &s,
		}, nil)
	tags := types.Tag{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)
	return client.Services{
		Apprunner: m,
	}
}

func TestObservabilityConfiguration(t *testing.T) {
	client.AwsMockTestHelper(t, ObservabilityConfigurations(), buildObservabilityConfiguration, client.TestOptions{})
}
