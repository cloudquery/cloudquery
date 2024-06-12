package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	teamapi "github.com/cloudquery/cloudquery/cli/internal/team"
)

func getAvailableUserTeams(ctx context.Context, token auth.Token) []string {
	cl, err := teamapi.NewClient(token.Value)
	if err != nil {
		return nil
	}
	teams, err := cl.ListAllTeams(ctx)
	if err != nil {
		return nil
	}
	return teams
}

func configFileMissing(err error) bool {
	return err != nil && errors.Is(err, os.ErrNotExist)
}

func emptyTeam(err error, team string) bool {
	return err == nil && team == ""
}

// GetTeamForToken returns the team for the given token
// If the token is a bearer token, we need to get the team from the configuration.
// For api keys the team name is not required as the key is bound to a team name already.
func GetTeamForToken(ctx context.Context, token auth.Token) (string, error) {
	switch token.Type {
	case auth.BearerToken:
		team, err := config.GetValue("team")
		if configFileMissing(err) || emptyTeam(err, team) {
			teams := getAvailableUserTeams(ctx, token)
			message := "team is not set.\n\n. Hint: use `cloudquery login` and/or `cloudquery switch <team>`"
			if len(teams) > 0 {
				message = fmt.Sprintf("team is not set.\n\nAvailable teams: %s.\n\nTo set your team, run `cloudquery switch <team>`", strings.Join(teams, ", "))
			}
			return "", fmt.Errorf(message)
		}
		if err != nil {
			return "", fmt.Errorf("failed to get team name from config: %w. Hint: use `cloudquery login` and/or `cloudquery switch <team>`", err)
		}
		return team, nil
	case auth.APIKey:
		cl, err := teamapi.NewClient(token.Value)
		if err != nil {
			return "", fmt.Errorf("failed to create API client for api key: %w", err)
		}
		teams, err := cl.ListAllTeams(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to list teams for api key: %w", err)
		}
		switch l := len(teams); l {
		case 0:
			return "", errors.New("team api key has no assigned team")
		case 1:
			return teams[0], nil
		default:
			return "", fmt.Errorf("team api key has more than one team: %s", strings.Join(teams, ", "))
		}
	default:
		return os.Getenv("_CQ_TEAM_NAME"), nil
	}
}
