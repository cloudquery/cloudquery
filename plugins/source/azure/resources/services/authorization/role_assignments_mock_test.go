// Auto generated code - DO NOT EDIT.

package authorization

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
)

func TestAuthorizationRoleAssignments(t *testing.T) {
	client.MockTestHelper(t, RoleAssignments(), createRoleAssignmentsMock)
}

func createRoleAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockAuthorizationRoleAssignmentsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleAssignments: mockClient,
		},
	}

	data := authorization.RoleAssignment{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := authorization.NewRoleAssignmentListResultPage(authorization.RoleAssignmentListResult{Value: &[]authorization.RoleAssignment{data}}, func(ctx context.Context, result authorization.RoleAssignmentListResult) (authorization.RoleAssignmentListResult, error) {
		return authorization.RoleAssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").Return(result, nil)
	return s
}
