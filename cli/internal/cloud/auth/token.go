package auth

import (
	"strings"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/rs/zerolog"
)

func tokenNeeded(registry specs.Registry, path string) bool {
	return registry == specs.RegistryCloudQuery || (registry == specs.RegistryDocker && strings.HasPrefix(path, "docker.cloudquery.io"))
}

func transformerNeedsToken(transformers []*specs.Transformer) bool {
	for _, transformer := range transformers {
		if tokenNeeded(transformer.Registry, transformer.Path) {
			return true
		}
	}
	return false
}

func sourcesNeedToken(sources []*specs.Source) bool {
	for _, source := range sources {
		if tokenNeeded(source.Registry, source.Path) {
			return true
		}
	}
	return false
}

func destinationsNeedToken(destinations []*specs.Destination) bool {
	for _, destination := range destinations {
		if tokenNeeded(destination.Registry, destination.Path) {
			return true
		}
	}
	return false
}

func GetAuthTokenIfNeeded(logger zerolog.Logger, sources []*specs.Source, destinations []*specs.Destination, transformers []*specs.Transformer) (cqapiauth.Token, error) {
	needsToken := sourcesNeedToken(sources) || destinationsNeedToken(destinations) || transformerNeedsToken(transformers)
	if !needsToken {
		logger.Debug().Msg("no need to get token since no source or destination uses the CloudQuery registry")
		return cqapiauth.UndefinedToken, nil
	}

	tc := cqapiauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		hasLoginHint := strings.Contains(err.Error(), "Hint:")
		if hasLoginHint {
			return cqapiauth.UndefinedToken, nil
		}

		return cqapiauth.UndefinedToken, err
	}

	return token, nil
}

func GetTokenType() cqapiauth.TokenType {
	tc := cqapiauth.NewTokenClient()
	return tc.GetTokenType()
}
