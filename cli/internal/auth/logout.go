package auth

import (
	"fmt"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
)

func Logout() error {
	err := cqapiauth.RemoveRefreshToken()
	if err != nil {
		return fmt.Errorf("failed to remove refresh token: %w", err)
	}
	return nil
}
