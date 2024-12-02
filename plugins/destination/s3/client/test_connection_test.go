package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/s3/v7/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
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
			name: "ok",
			spec: []byte(`{"bucket": "test", "region": "test", "path": "test", "format": "csv"}`),
			clientBuilder: func() (plugin.Client, error) {
				return &Client{}, nil
			},
		},
		{
			name: "error/unauthorized",
			spec: []byte(`{"bucket": "test", "region": "test", "path": "test", "format": "csv"}`),
			err:  plugin.NewTestConnError(codeUnauthorized, assert.AnError),
			clientBuilder: func() (plugin.Client, error) {
				return nil, errTestWriteFailed
			},
		},
		{
			name: "error/spec",
			spec: []byte(`{null}`),
			err:  plugin.NewTestConnError(codeInvalidSpec, assert.AnError),
			clientBuilder: func() (plugin.Client, error) {
				return &Client{}, nil
			},
		},
	}

	for idx := range cases {
		tc := cases[idx]

		t.Run(tc.name, func(t *testing.T) {
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				spec := &spec.Spec{}
				if err := json.Unmarshal(specBytes, spec); err != nil {
					return nil, err
				}
				spec.SetDefaults()
				if err := spec.Validate(); err != nil {
					return nil, err
				}

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
