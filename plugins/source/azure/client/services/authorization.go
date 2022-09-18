//go:generate mockgen -destination=./mocks/authorization.go -package=mocks . AuthorizationRoleAssignmentsClient,AuthorizationRoleDefinitionsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/Azure/go-autorest/autorest"
)

type AuthorizationClient struct {
	RoleAssignments AuthorizationRoleAssignmentsClient
	RoleDefinitions AuthorizationRoleDefinitionsClient
}

type AuthorizationRoleAssignmentsClient interface {
	List(ctx context.Context, filter string) (result authorization.RoleAssignmentListResultPage, err error)
}

type AuthorizationRoleDefinitionsClient interface {
	List(ctx context.Context, scope string, filter string) (result authorization.RoleDefinitionListResultPage, err error)
}

func NewAuthorizationClient(subscriptionId string, auth autorest.Authorizer) AuthorizationClient {
	assignments := authorization.NewRoleAssignmentsClient(subscriptionId)
	assignments.Authorizer = auth
	definitions := authorization.NewRoleDefinitionsClient(subscriptionId)
	definitions.Authorizer = auth
	return AuthorizationClient{
		RoleAssignments: assignments,
		RoleDefinitions: definitions,
	}
}
