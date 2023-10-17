package team

import (
	"fmt"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
)

func Logout() error {
	err := auth.RemoveRefreshToken()
	if err != nil {
		return fmt.Errorf("failed to remove refresh token: %w", err)
	}
	err = config.UnsetValue("team")
	if err != nil {
		return fmt.Errorf("failed to reset team value: %w", err)
	}
	return nil
}
