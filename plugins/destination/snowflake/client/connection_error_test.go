package client

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewPingErrorOmitsSecrets(t *testing.T) {
	const password = "s3cr3t-pa55w0rd"

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	keyDER, err := x509.MarshalPKCS8PrivateKey(key)
	require.NoError(t, err)
	keyPEM := string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer server.Close()
	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	connectionString := "user:" + password + "@" + serverURL.Host + "/db/schema?account=account&protocol=http&loginTimeout=1"

	cases := []struct {
		name string
		spec Spec
	}{
		{name: "password", spec: Spec{ConnectionString: connectionString}},
		{name: "private key", spec: Spec{ConnectionString: connectionString, PrivateKey: keyPEM}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			specBytes, err := json.Marshal(tc.spec)
			require.NoError(t, err)

			_, err = New(context.Background(), zerolog.Nop(), specBytes, plugin.NewClientOptions{})
			require.Error(t, err)

			message := err.Error()
			require.Contains(t, message, "ping failed")
			require.NotContains(t, message, password)
			require.NotContains(t, message, base64.URLEncoding.EncodeToString(keyDER))
			require.NotContains(t, message, "privateKey=")
		})
	}
}
