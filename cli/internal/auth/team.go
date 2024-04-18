package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/api"
)

// GetTeamForToken returns the team for the given token
// If the token is a bearer token, we need to get the team from the configuration.
// For api keys the team name is not required as the key is bound to a team name already.
func GetTeamForToken(ctx context.Context, token auth.Token) (string, error) {
	switch token.Type {
	case auth.BearerToken:
		team, err := config.GetValue("team")
		if err != nil {
			return "", fmt.Errorf("failed to get team from config: %w", err)
		}
		return team, nil
	case auth.APIKey:
		apiClient, err := api.NewClient(token.Value)
		if err != nil {
			return "", fmt.Errorf("failed to create api client: %w", err)
		}
		resp, err := apiClient.ListTeamsWithResponse(ctx, &cqapi.ListTeamsParams{})
		if err != nil {
			return "", fmt.Errorf("failed to list teams: %w", err)
		}
		if resp.StatusCode() != http.StatusOK {
			return "", fmt.Errorf("failed to list teams for API key, status code: %s", resp.Status())
		}
		return resp.JSON200.Items[0].Name, nil
	default:
		return os.Getenv("_CQ_TEAM_NAME"), nil
	}
}
