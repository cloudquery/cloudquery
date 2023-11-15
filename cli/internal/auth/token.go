package auth

import (
	"strings"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/rs/zerolog"
)

func sourcesNeedToken(sources []*specs.Source) bool {
	for _, source := range sources {
		if source.Registry == specs.RegistryCloudQuery {
			return true
		}
	}
	return false
}

func destinationsNeedToken(destinations []*specs.Destination) bool {
	for _, destination := range destinations {
		if destination.Registry == specs.RegistryCloudQuery {
			return true
		}
	}
	return false
}

func GetAuthTokenIfNeeded(logger zerolog.Logger, sources []*specs.Source, destinations []*specs.Destination) (string, error) {
	needsToken := sourcesNeedToken(sources) || destinationsNeedToken(destinations)
	if !needsToken {
		logger.Debug().Msg("no need to get token since no source or destination uses the CloudQuery registry")
		return "", nil
	}

	tc := cqapiauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		recommendLogin := strings.Contains(err.Error(), "Hint:")
		if recommendLogin {
			logger.Warn().Msg("when using the CloudQuery registry, it's recommended to log in via `cloudquery login`. Logging in allows for better rate limits and downloading of premium plugins")
			return "", nil
		}

		return "", err
	}

	return token.Value, nil
}
