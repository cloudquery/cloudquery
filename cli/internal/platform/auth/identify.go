package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cqapi "github.com/cloudquery/cloudquery-platform-api-go"
	"github.com/cloudquery/cloudquery-platform-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform/api"
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

	if resp.JSON200 == nil {
		return nil, errors.New("failed to get current user: no response data")
	}

	// Convert the response to User type
	user := &cqapi.User{
		ID:              resp.JSON200.ID,
		Email:           resp.JSON200.Email,
		Name:            resp.JSON200.Name,
		CreatedAt:       resp.JSON200.CreatedAt,
		UpdatedAt:       resp.JSON200.UpdatedAt,
		LastLoginAt:     resp.JSON200.LastLoginAt,
		ProfileImageURL: resp.JSON200.ProfileImageURL,
	}

	return user, nil
}
