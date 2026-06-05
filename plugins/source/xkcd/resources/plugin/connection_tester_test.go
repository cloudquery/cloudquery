package plugin

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/xkcd/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	cases := []struct {
		name     string
		spec     string
		status   int
		wantCode string
	}{
		{name: "ok", spec: `{}`, status: http.StatusOK},
		{name: "invalid spec", spec: `{"concurrency":"nope"}`, wantCode: codeInvalidSpec},
		{name: "connection failed", spec: `{}`, status: http.StatusInternalServerError, wantCode: codeConnectionFailed},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				if tc.status != http.StatusOK {
					w.WriteHeader(tc.status)
					return
				}
				_, _ = w.Write([]byte(`{"num":1,"title":"x"}`))
			}))
			defer ts.Close()

			tester := connectionTester(func() (*xkcd.Client, error) {
				return xkcd.NewClient(xkcd.WithBaseURL(ts.URL), xkcd.WithHTTPClient(ts.Client()))
			})
			err := tester(context.Background(), zerolog.Nop(), []byte(tc.spec))
			if tc.wantCode == "" {
				require.NoError(t, err)
				return
			}
			var connErr *plugin.TestConnError
			require.ErrorAs(t, err, &connErr)
			require.Equal(t, tc.wantCode, connErr.Code)
		})
	}
}
