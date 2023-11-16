package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
)

// GetTeamForToken returns the team for the given token
// If the token is a bearer token, we need to get the team from the configuration.
// For api keys the team name is not required as the key is bound to a team name already.
func GetTeamForToken(token auth.Token) (string, error) {
	if token.Type == auth.BearerToken {
		team, err := config.GetValue("team")
		if err != nil {
			return "", fmt.Errorf("failed to get team from config: %w", err)
		}
		return team, nil
	}
	return "", nil
}

// RequireTeamForToken returns the team for the given token. If the team is not set, it returns an error.
func RequireTeamForToken(token auth.Token) (string, error) {
	if token.Type == auth.BearerToken {
		team, err := config.GetValue("team")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("failed to get team from config: %w", err)
		}
		if team == "" { // err is os.ErrNotExist or team is empty
			return "", fmt.Errorf("team is required. Hint: use `cloudquery switch` to set a team")
		}
		return team, nil
	}
	return "", nil
}
