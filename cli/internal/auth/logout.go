package auth

import (
	"fmt"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
)

func Logout() error {
	err := cqapiauth.RemoveRefreshToken()
	if err != nil {
		return fmt.Errorf("failed to remove refresh token: %w", err)
	}

	err = config.UnsetValue("team")
	if err != nil {
		return fmt.Errorf("failed to unset team: %w", err)
	}

	return nil
}
