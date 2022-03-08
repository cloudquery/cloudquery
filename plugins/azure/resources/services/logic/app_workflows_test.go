package logic

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildLogicAppWorkflowMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	diagnosticSettings := mocks.NewMockMonitorDiagnosticSettingsClient(ctrl)
	workflows := mocks.NewMockWorkflowsClient(ctrl)
	var diagnosticSettingsResource insights.DiagnosticSettingsResource
	if err := faker.FakeDataSkipFields(&diagnosticSettingsResource, []string{""}); err != nil {
		t.Fatal(err)
	}
	var workflow logic.Workflow
	if err := faker.FakeDataSkipFields(&workflow, []string{"WorkflowProperties"}); err != nil {
		t.Fatal(err)
	}
	var workflowProperties logic.WorkflowProperties
	if err := faker.FakeDataSkipFields(&workflowProperties, []string{"ProvisioningState", "State", "Definition", "Parameters"}); err != nil {
		t.Fatal(err)
	}
	workflowProperties.ProvisioningState = logic.WorkflowProvisioningStateAccepted
	workflowProperties.State = logic.WorkflowStateCompleted
	// workflow.Definition = nil
	workflow.WorkflowProperties = &workflowProperties
	// workflow.Parameters = make(map[string]*logic.WorkflowParameter)
	diagnosticSettings.EXPECT().List(gomock.Any(), gomock.Any()).Return(
		insights.DiagnosticSettingsResourceCollection{
			Value: &[]insights.DiagnosticSettingsResource{diagnosticSettingsResource},
		}, nil,
	)
	workflows.EXPECT().ListBySubscription(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		logic.NewWorkflowListResultPage(
			logic.WorkflowListResult{Value: &[]logic.Workflow{workflow}},
			func(context.Context, logic.WorkflowListResult) (logic.WorkflowListResult, error) {
				return logic.WorkflowListResult{}, nil
			},
		), nil,
	)
	return services.Services{
		Logic: services.LogicClient{
			DiagnosticSettings: diagnosticSettings,
			Workflows:          workflows,
		},
	}
}

func TestLogicAppWorkflows(t *testing.T) {
	client.AzureMockTestHelper(t, LogicAppWorkflows(), buildLogicAppWorkflowMock, client.TestOptions{})
}
