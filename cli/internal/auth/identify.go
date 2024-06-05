package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/api"
)

func GetUserId(ctx context.Context, token auth.Token) (string, error) {
	apiClient, err := api.NewClient(token.Value)
	if err != nil {
		return "", fmt.Errorf("failed to create api client: %w", err)
	}
	resp, err := apiClient.GetCurrentUserWithResponse(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("failed to get current user, status code: %s", resp.Status())
	}

	return resp.JSON200.ID.String(), nil
}
