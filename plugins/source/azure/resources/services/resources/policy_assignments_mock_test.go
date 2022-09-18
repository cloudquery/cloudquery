// Auto generated code - DO NOT EDIT.

package resources

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
)

func TestResourcesPolicyAssignments(t *testing.T) {
	client.MockTestHelper(t, PolicyAssignments(), createPolicyAssignmentsMock)
}

func createPolicyAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesPolicyAssignmentsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			PolicyAssignments: mockClient,
		},
	}

	data := policy.Assignment{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := policy.NewAssignmentListResultPage(policy.AssignmentListResult{Value: &[]policy.Assignment{data}}, func(ctx context.Context, result policy.AssignmentListResult) (policy.AssignmentListResult, error) {
		return policy.AssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), gomock.Any(), "", nil).Return(result, nil)
	return s
}
