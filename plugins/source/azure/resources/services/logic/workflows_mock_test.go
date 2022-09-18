// Auto generated code - DO NOT EDIT.

package logic

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
)

func TestLogicWorkflows(t *testing.T) {
	client.MockTestHelper(t, Workflows(), createWorkflowsMock)
}

func createWorkflowsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockLogicWorkflowsClient(ctrl)
	s := services.Services{
		Logic: services.LogicClient{
			Workflows:          mockClient,
			DiagnosticSettings: createDiagnosticSettingsMock(t, ctrl).Logic.DiagnosticSettings,
		},
	}

	data := logic.Workflow{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := logic.NewWorkflowListResultPage(logic.WorkflowListResult{Value: &[]logic.Workflow{data}}, func(ctx context.Context, result logic.WorkflowListResult) (logic.WorkflowListResult, error) {
		return logic.WorkflowListResult{}, nil
	})

	var top int32 = 100
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &top, "").Return(result, nil)
	return s
}
