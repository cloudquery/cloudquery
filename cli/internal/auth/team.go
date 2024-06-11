package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/api"
	"github.com/cloudquery/cloudquery/cli/internal/team"
)

func getAvailableUserTeams(ctx context.Context, token auth.Token) []string {
	cl, err := team.NewClient(token.Value)
	if err != nil {
		return nil
	}
	teams, err := cl.ListAllTeams(ctx)
	if err != nil {
		return nil
	}
	return teams
}

// GetTeamForToken returns the team for the given token
// If the token is a bearer token, we need to get the team from the configuration.
// For api keys the team name is not required as the key is bound to a team name already.
func GetTeamForToken(ctx context.Context, token auth.Token) (string, error) {
	switch token.Type {
	case auth.BearerToken:
		team, err := config.GetValue("team")
		if errors.Is(err, os.ErrNotExist) {
			teams := getAvailableUserTeams(ctx, token)
			message := "your current team is not set.\n\nTo set your current team, run `cloudquery switch <team>`"
			if len(teams) > 0 {
				message = fmt.Sprintf("your current team is not set.\n\nAvailable teams: %s.\n\nTo set your current team, run `cloudquery switch <team>`", strings.Join(teams, ", "))
			}
			return "", fmt.Errorf(message)
		}
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
