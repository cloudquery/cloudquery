package auth

import (
	"errors"
	"fmt"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
)

func Logout() error {
	var errs []error
	if err := cqapiauth.RemoveRefreshToken(); err != nil {
		errs = append(errs, fmt.Errorf("failed to remove refresh token: %w", err))
	}

	if err := config.UnsetValue("team"); err != nil {
		errs = append(errs, fmt.Errorf("failed to unset team: %w", err))
	}

	return errors.Join(errs...)
}
