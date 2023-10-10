package auth

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cli/internal/config"
)

func Logout() error {
	err := removeRefreshToken()
	if err != nil {
		return fmt.Errorf("failed to remove refresh token: %w", err)
	}
	err = config.UnsetValue("team")
	if err != nil {
		return fmt.Errorf("failed to reset team value: %w", err)
	}
	return nil
}
