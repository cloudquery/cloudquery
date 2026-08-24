package platform

import (
	"errors"
	"strings"

	"github.com/cloudquery/cloudquery-api-go/config"
)

// tokenStorePath is a CLI-owned data file, separate from cloudquery/token: that
// file is reserved for Firebase refresh tokens, which the api-go token client
// posts to Firebase on every read — a cqpd_ token stored there would fail auth.
const tokenStorePath = "cloudquery/platform_token"

// IsPlatformToken reports whether the value is a platform-destination (cqpd_)
// token, e.g. one returned by a tenant's browser-login callback.
func IsPlatformToken(token string) bool {
	return strings.HasPrefix(token, cqpdPrefix)
}

// SavePlatformToken persists a cqpd_ token from a browser login so later CLI
// invocations pick it up without env vars.
func SavePlatformToken(token string) error {
	if !IsPlatformToken(token) {
		return errors.New("not a platform token")
	}
	return config.SaveDataString(tokenStorePath, token)
}

// ReadPlatformToken returns the saved cqpd_ token, or "" when none is stored.
func ReadPlatformToken() string {
	token, err := config.ReadDataString(tokenStorePath)
	if err != nil || !IsPlatformToken(token) {
		return ""
	}
	return token
}

// RemovePlatformToken deletes the saved token. Missing file is not an error.
func RemovePlatformToken() error {
	return config.DeleteDataString(tokenStorePath)
}
