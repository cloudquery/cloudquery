package platform

import (
	"encoding/base64"
	"testing"

	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/stretchr/testify/require"
)

func mintTestToken(t *testing.T, apiURL, team string) string {
	t.Helper()
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"u":"` + apiURL + `","tm":"` + team + `"}`))
	sig := base64.RawURLEncoding.EncodeToString([]byte("sig"))
	return cqpdPrefix + payload + "." + sig
}

func setupTokenStore(t *testing.T) {
	t.Helper()
	require.NoError(t, config.SetDataHome(t.TempDir()))
	t.Cleanup(func() {
		require.NoError(t, config.UnsetDataHome())
	})
}

func TestPlatformTokenStoreRoundtrip(t *testing.T) {
	setupTokenStore(t)

	require.Empty(t, ReadPlatformToken())

	token := mintTestToken(t, "https://acme.mycloudquery.com", "team-x")
	require.NoError(t, SavePlatformToken(token))
	require.Equal(t, token, ReadPlatformToken())

	require.NoError(t, RemovePlatformToken())
	require.Empty(t, ReadPlatformToken())
	// removing again is not an error
	require.NoError(t, RemovePlatformToken())
}

func TestSavePlatformTokenRejectsNonPlatformToken(t *testing.T) {
	setupTokenStore(t)

	require.Error(t, SavePlatformToken("some-firebase-refresh-token"))
	require.Empty(t, ReadPlatformToken())
}

func TestPlatformTokenEnvPrecedence(t *testing.T) {
	setupTokenStore(t)

	stored := mintTestToken(t, "https://stored.mycloudquery.com", "stored-team")
	require.NoError(t, SavePlatformToken(stored))

	t.Setenv(EnvPlatformToken, "")
	t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, "")
	require.Equal(t, stored, platformToken())

	fromKey := mintTestToken(t, "https://key.mycloudquery.com", "key-team")
	t.Setenv(cqapiauth.EnvVarCloudQueryAPIKey, fromKey)
	require.Equal(t, fromKey, platformToken())

	fromEnv := mintTestToken(t, "https://env.mycloudquery.com", "env-team")
	t.Setenv(EnvPlatformToken, fromEnv)
	require.Equal(t, fromEnv, platformToken())
}
