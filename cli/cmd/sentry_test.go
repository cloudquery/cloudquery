package cmd

import (
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/secrets"
	"github.com/stretchr/testify/require"
)

func TestRedactSecrets(t *testing.T) {
	previous := secretAwareRedactor
	t.Cleanup(func() { secretAwareRedactor = previous })

	secretAwareRedactor = nil
	require.Equal(t, "panic: super-secret-value", RedactSecrets("panic: super-secret-value"))

	secretAwareRedactor = secrets.NewSecretAwareRedactor()
	secretAwareRedactor.AddSecretEnv([]string{"MY_TOKEN=super-secret-value"})
	require.Equal(t, "panic: MY_TOKEN", RedactSecrets("panic: super-secret-value"))
}
