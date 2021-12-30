//go:build !integration

package authorization

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAuthorizationRoleDefinitions(t *testing.T, ctrl *gomock.Controller) services.Services {
	defs := mocks.NewMockRoleDefinitionsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleDefinitions: defs,
		},
	}

	var def authorization.RoleDefinition
	if err := faker.FakeData(&def); err != nil {
		t.Fatal(err)
	}
	defs.EXPECT().List(gomock.Any(), gomock.Any(), "").Return(
		authorization.NewRoleDefinitionListResultPage(
			authorization.RoleDefinitionListResult{Value: &[]authorization.RoleDefinition{def}},
			func(context.Context, authorization.RoleDefinitionListResult) (authorization.RoleDefinitionListResult, error) {
				return authorization.RoleDefinitionListResult{}, nil
			},
		), nil,
	)
	return s
}

func TestAuthorizationRoleDefinitions(t *testing.T) {
	client.AzureMockTestHelper(t, AuthorizationRoleDefinitions(), buildAuthorizationRoleDefinitions, client.TestOptions{})
}
