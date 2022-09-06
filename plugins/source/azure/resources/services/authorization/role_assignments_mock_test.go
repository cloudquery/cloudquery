// Auto generated code - DO NOT EDIT.

package authorization

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
)

func TestAuthorizationRoleAssignments(t *testing.T) {
	client.AzureMockTestHelper(t, RoleAssignments(), createRoleAssignmentsMock, client.TestOptions{})
}

func createRoleAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockAuthorizationRoleAssignmentsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleAssignments: mockClient,
		},
	}

	data := authorization.RoleAssignment{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := authorization.NewRoleAssignmentListResultPage(authorization.RoleAssignmentListResult{Value: &[]authorization.RoleAssignment{data}}, func(ctx context.Context, result authorization.RoleAssignmentListResult) (authorization.RoleAssignmentListResult, error) {
		return authorization.RoleAssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").Return(result, nil)
	return s
}
