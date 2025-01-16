package api

import (
	"context"
	"net/http"
	"testing"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/stretchr/testify/require"
)

const (
	testKey       = "testkey"
	unexpectedKey = "whatkey"
)

func TestLocalClientDefaultAPI(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", testKey)
	t.Setenv("CLOUDQUERY_API_URL", "")
	t.Setenv("CQ_CLOUD", "")
	t.Setenv("CLOUDQUERY_CLI_API_URL", "")
	t.Setenv("CLOUDQUERY_CLI_KEY", "")

	r := require.New(t)
	c, err := NewLocalClient()
	r.NoError(err)

	cc := c.ClientInterface.(*cloudquery_api.Client)
	r.Equal("https://api.cloudquery.io/", cc.Server)
	r.Len(cc.RequestEditors, 1)
	ensureRequestAuthHeader(t, cc.RequestEditors[0], "Bearer "+testKey)
}

func TestLocalClientLocalAPI(t *testing.T) {
	// Both keys are set
	t.Setenv("CLOUDQUERY_API_KEY", unexpectedKey)
	t.Setenv("CLOUDQUERY_API_URL", "")
	t.Setenv("CQ_CLOUD", "1")
	t.Setenv("CLOUDQUERY_CLI_API_URL", "http://localhost:8080")
	t.Setenv("CLOUDQUERY_CLI_KEY", testKey)

	r := require.New(t)
	c, err := NewLocalClient()
	r.NoError(err)

	cc := c.ClientInterface.(*cloudquery_api.Client)
	r.Equal("http://localhost:8080/", cc.Server)
	r.Len(cc.RequestEditors, 1)
	ensureRequestAuthHeader(t, cc.RequestEditors[0], "Bearer "+testKey)
}

func TestLocalClientNotLocalAPI(t *testing.T) {
	// Both env var pairs are set but CQ_CLOUD is falsy
	t.Setenv("CLOUDQUERY_API_KEY", testKey)
	t.Setenv("CLOUDQUERY_API_URL", "http://localhost:3333")
	t.Setenv("CQ_CLOUD", "0")
	t.Setenv("CLOUDQUERY_CLI_API_URL", "http://localhost:8080")
	t.Setenv("CLOUDQUERY_CLI_KEY", unexpectedKey)

	r := require.New(t)
	c, err := NewLocalClient()
	r.NoError(err)

	cc := c.ClientInterface.(*cloudquery_api.Client)
	r.Equal("http://localhost:3333/", cc.Server)
	r.Len(cc.RequestEditors, 1)
	ensureRequestAuthHeader(t, cc.RequestEditors[0], "Bearer "+testKey)
}

func ensureRequestAuthHeader(t *testing.T, f cloudquery_api.RequestEditorFn, expectedValue string) {
	t.Helper()

	req, err := http.NewRequest(http.MethodGet, "https://api.cloudquery.io", nil)
	require.NoError(t, err)
	err = f(context.TODO(), req)
	require.NoError(t, err)

	val := req.Header.Get("Authorization")
	require.Equal(t, expectedValue, val)
}
