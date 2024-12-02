package auth

import (
	"context"
	"fmt"
	"net/http"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
)

func GetUser(ctx context.Context, token auth.Token) (*cqapi.User, error) {
	apiClient, err := api.NewClient(token.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}
	resp, err := apiClient.GetCurrentUserWithResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to get current user, status code: %s", resp.Status())
	}

	return resp.JSON200, nil
}
