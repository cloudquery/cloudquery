package client

import (
	"context"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_user.go . UserService
type UserService interface {
	ListUsers(ctx context.Context, qp *query.Params) ([]*okta.User, *okta.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_group.go . GroupService
type GroupService interface {
	ListGroups(ctx context.Context, qp *query.Params) ([]*okta.Group, *okta.Response, error)
	ListGroupUsers(ctx context.Context, groupId string, qp *query.Params) ([]*okta.User, *okta.Response, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_client_application.go . ApplicationService
type ApplicationService interface {
	ListApplications(ctx context.Context, qp *query.Params) ([]okta.App, *okta.Response, error)
	ListApplicationUsers(ctx context.Context, appId string, qp *query.Params) ([]*okta.AppUser, *okta.Response, error)
	ListApplicationGroupAssignments(ctx context.Context, appId string, qp *query.Params) ([]*okta.ApplicationGroupAssignment, *okta.Response, error)
}
