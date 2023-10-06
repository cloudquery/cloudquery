// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"net/http"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

//go:generate mockgen -package=mocks -destination=../mocks/users_api.go -source=users_api.go UsersAPIClient
type UsersAPIClient interface {
	ListUserOrganizations(context.Context, string) (datadogV2.UserResponse, *http.Response, error)
	ListUserPermissions(context.Context, string) (datadogV2.PermissionsResponse, *http.Response, error)
	ListUsers(context.Context, ...datadogV2.ListUsersOptionalParameters) (datadogV2.UsersResponse, *http.Response, error)
	ListUsersWithPagination(context.Context, ...datadogV2.ListUsersOptionalParameters) (<-chan datadog.PaginationResult[datadogV2.User], func())
}
