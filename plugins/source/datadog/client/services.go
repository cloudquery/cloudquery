package client

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"net/http"
)

type DatadogServices struct {
	Users UsersService
}

//go:generate mockgen -package=mocks -destination=../mocks/users_mock.go . UsersService
type UsersService interface {
	ListUsers(ctx context.Context, o ...datadogV2.ListUsersOptionalParameters) (datadogV2.UsersResponse, *http.Response, error)
	ListUserPermissions(ctx context.Context, userID string) (datadogV2.PermissionsResponse, *http.Response, error)
	ListUserOrganizations(ctx context.Context, userId string) (datadogV2.UserResponse, *http.Response, error)
}
