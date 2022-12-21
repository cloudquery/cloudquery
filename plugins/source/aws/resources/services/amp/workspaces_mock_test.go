package amp

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildWorkspaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAmpClient(ctrl)

	var summary types.WorkspaceSummary
	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListWorkspaces(gomock.Any(), gomock.Any()).Return(
		&amp.ListWorkspacesOutput{
			Workspaces: []types.WorkspaceSummary{summary},
		},
		nil,
	)

	var description types.WorkspaceDescription
	if err := faker.FakeObject(&description); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeWorkspace(gomock.Any(), gomock.Any()).Return(
		&amp.DescribeWorkspaceOutput{
			Workspace: &description,
		},
		nil,
	)

	var alertManagerDefinition types.AlertManagerDefinitionDescription
	if err := faker.FakeObject(&alertManagerDefinition); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAlertManagerDefinition(gomock.Any(), gomock.Any()).Return(
		&amp.DescribeAlertManagerDefinitionOutput{
			AlertManagerDefinition: &alertManagerDefinition,
		},
		nil,
	)

	var loggingConfiguration types.LoggingConfigurationMetadata
	if err := faker.FakeObject(&loggingConfiguration); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoggingConfiguration(gomock.Any(), gomock.Any()).Return(
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
