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

func TestAuthorizationRoleDefinitions(t *testing.T) {
	client.MockTestHelper(t, RoleDefinitions(), createRoleDefinitionsMock)
}

func createRoleDefinitionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockAuthorizationRoleDefinitionsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleDefinitions: mockClient,
		},
	}

	data := authorization.RoleDefinition{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := authorization.NewRoleDefinitionListResultPage(authorization.RoleDefinitionListResult{Value: &[]authorization.RoleDefinition{data}}, func(ctx context.Context, result authorization.RoleDefinitionListResult) (authorization.RoleDefinitionListResult, error) {
		return authorization.RoleDefinitionListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), gomock.Any(), "").Return(result, nil)
	return s
}
