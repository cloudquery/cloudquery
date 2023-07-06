package amp

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWorkspaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAmpClient(ctrl)

	var summary types.WorkspaceSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListWorkspaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.ListWorkspacesOutput{
			Workspaces: []types.WorkspaceSummary{summary},
		},
		nil,
	)

	var description types.WorkspaceDescription
	require.NoError(t, faker.FakeObject(&description))

	m.EXPECT().DescribeWorkspace(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.DescribeWorkspaceOutput{
			Workspace: &description,
		},
		nil,
	)

	var alertManagerDefinition types.AlertManagerDefinitionDescription
	require.NoError(t, faker.FakeObject(&alertManagerDefinition))

	m.EXPECT().DescribeAlertManagerDefinition(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.DescribeAlertManagerDefinitionOutput{
			AlertManagerDefinition: &alertManagerDefinition,
		},
		nil,
	)

	var loggingConfiguration types.LoggingConfigurationMetadata
	require.NoError(t, faker.FakeObject(&loggingConfiguration))

	m.EXPECT().DescribeLoggingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amp.DescribeLoggingConfigurationOutput{
			LoggingConfiguration: &loggingConfiguration,
		},
		nil,
	)

	buildRuleGroupsNamespaces(t, m)

	return client.Services{Amp: m}
}

func TestWorkspaces(t *testing.T) {
	client.AwsMockTestHelper(t, Workspaces(), buildWorkspaces, client.TestOptions{})
}
