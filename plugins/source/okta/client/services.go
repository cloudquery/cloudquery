package client

import (
	"context"

	"github.com/okta/okta-sdk-golang/v3/okta"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_user.go . UserService
type UserService interface {
	ListUsers(ctx context.Context) okta.ApiListUsersRequest
	ListUsersExecute(r okta.ApiListUsersRequest) ([]okta.User, *okta.APIResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_group.go . GroupService
type GroupService interface {
	ListGroups(ctx context.Context) okta.ApiListGroupsRequest
	ListGroupsExecute(r okta.ApiListGroupsRequest) ([]okta.Group, *okta.APIResponse, error)

	ListGroupUsers(ctx context.Context, groupId string) okta.ApiListGroupUsersRequest
	ListGroupUsersExecute(r okta.ApiListGroupUsersRequest) ([]okta.User, *okta.APIResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_application.go . ApplicationService
type ApplicationService interface {
	ListApplications(ctx context.Context) okta.ApiListApplicationsRequest
	ListApplicationsExecute(r okta.ApiListApplicationsRequest) ([]okta.ListApplications200ResponseInner, *okta.APIResponse, error)

	ListApplicationUsers(ctx context.Context, appId string) okta.ApiListApplicationUsersRequest
	ListApplicationUsersExecute(r okta.ApiListApplicationUsersRequest) ([]okta.AppUser, *okta.APIResponse, error)

	ListApplicationGroupAssignments(ctx context.Context, appId string) okta.ApiListApplicationGroupAssignmentsRequest
	ListApplicationGroupAssignmentsExecute(r okta.ApiListApplicationGroupAssignmentsRequest) ([]okta.ApplicationGroupAssignment, *okta.APIResponse, error)
}
