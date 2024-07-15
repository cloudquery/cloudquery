package client

import (
	"context"
	"net"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	cases := []struct {
		name          string
		spec          []byte
		err           *plugin.TestConnError
		clientBuilder func() (plugin.Client, error)
	}{
		{
			name:          "ok",
			spec:          []byte(`{}`),
			clientBuilder: func() (plugin.Client, error) { return &Client{}, nil },
		},
		{
			name: "error/unauthorized",
			spec: []byte(`{}`),
			err:  plugin.NewTestConnError(codeUnauthorized, assert.AnError),
			clientBuilder: func() (plugin.Client, error) {
				return nil, &mssql.Error{Number: 1}
			},
		},
		{
			name: "error/unreachable",
			spec: []byte(`{}`),
			err:  plugin.NewTestConnError(codeUnreachable, assert.AnError),
			clientBuilder: func() (plugin.Client, error) {
				return nil, &net.OpError{}
			},
		},
		{
			name:          "error/spec",
			spec:          []byte(`{null}`),
			err:           plugin.NewTestConnError(codeInvalidSpec, assert.AnError),
			clientBuilder: func() (plugin.Client, error) { return nil, errInvalidSpec },
		},
		{
			name:          "error/connection_failed",
			spec:          []byte(`{}`),
			err:           plugin.NewTestConnError(codeConnectionFailed, assert.AnError),
			clientBuilder: func() (plugin.Client, error) { return nil, assert.AnError },
		},
	}

	for idx := range cases {
		tc := cases[idx]

		t.Run(tc.name, func(t *testing.T) {
			tester := NewConnectionTester(func(context.Context, zerolog.Logger, []byte, plugin.NewClientOptions) (plugin.Client, error) {
				return tc.clientBuilder()
			})

			err := tester(context.Background(), zerolog.Nop(), tc.spec)
			if tc.err == nil {
				require.NoError(t, err)
				return
			}

			var e *plugin.TestConnError
			require.ErrorAs(t, err, &e)
			require.Equal(t, tc.err.Code, err.(*plugin.TestConnError).Code)
		})
	}
}
