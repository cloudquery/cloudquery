package resources_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAuthorizationRoleAssignments(t *testing.T, ctrl *gomock.Controller) services.Services {
	assignments := mocks.NewMockRoleAssignmentsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleAssignments: assignments,
		},
	}

	var a authorization.RoleAssignment
	if err := faker.FakeData(&a); err != nil {
		t.Fatal(err)
	}
	assignments.EXPECT().List(gomock.Any(), "").Return(
		authorization.NewRoleAssignmentListResultPage(
			authorization.RoleAssignmentListResult{Value: &[]authorization.RoleAssignment{a}},
			func(context.Context, authorization.RoleAssignmentListResult) (authorization.RoleAssignmentListResult, error) {
				return authorization.RoleAssignmentListResult{}, nil
			},
		), nil,
	)

	return s
}

func TestAuthorizationRoleAssignments(t *testing.T) {
	azureTestHelper(t, resources.AuthorizationRoleAssignments(), buildAuthorizationRoleAssignments)
}
