package appconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppconfigClient(ctrl)

	var app types.Application
	require.NoError(t, faker.FakeObject(&app))

	m.EXPECT().ListApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListApplicationsOutput{
			Items: []types.Application{app},
		},
		nil,
	)

	var env types.Environment
	require.NoError(t, faker.FakeObject(&env))

	m.EXPECT().ListEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListEnvironmentsOutput{
			Items: []types.Environment{env},
		},
		nil,
	)

	var configProfile types.ConfigurationProfileSummary
	require.NoError(t, faker.FakeObject(&configProfile))

	m.EXPECT().ListConfigurationProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListConfigurationProfilesOutput{
			Items: []types.ConfigurationProfileSummary{configProfile},
		},
		nil,
	)
	var configProfileOutput appconfig.GetConfigurationProfileOutput
	require.NoError(t, faker.FakeObject(&configProfileOutput))

	m.EXPECT().GetConfigurationProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configProfileOutput,
		nil,
	)

	var hostedConfigurationVersionSummary types.HostedConfigurationVersionSummary
	require.NoError(t, faker.FakeObject(&hostedConfigurationVersionSummary))

	m.EXPECT().ListHostedConfigurationVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListHostedConfigurationVersionsOutput{
			Items: []types.HostedConfigurationVersionSummary{hostedConfigurationVersionSummary},
		},
		nil,
	)
	var hostedConfigurationVersionOutput appconfig.GetHostedConfigurationVersionOutput
	require.NoError(t, faker.FakeObject(&hostedConfigurationVersionOutput))

	m.EXPECT().GetHostedConfigurationVersion(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&hostedConfigurationVersionOutput,
		nil,
	)

	return client.Services{Appconfig: m}
}

func TestApps(t *testing.T) {
	client.AwsMockTestHelper(t, Applications(), buildApps, client.TestOptions{})
}
