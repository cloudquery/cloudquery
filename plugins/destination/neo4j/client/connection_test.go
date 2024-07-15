package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	cases := []struct {
		name        string
		spec        []byte
		err         *plugin.TestConnError
		builderFunc func() (plugin.Client, error)
	}{
		{
			name: "ok",
			spec: []byte(`{"connection_string": "file", "username": "user", "password": "pass"}`),
			builderFunc: func() (plugin.Client, error) {
				return &Client{}, nil
			},
		},
		{
			name: "error/invalid_spec_unmarshall",
			spec: []byte(`{null}`),
			err:  plugin.NewTestConnError(codeInvalidSpec, assert.AnError),
			builderFunc: func() (plugin.Client, error) {
				return &Client{}, nil
			},
		},
		{
			name: "error/invalid_spec_validate",
			spec: []byte(`{"connection_string": "", "username": "user"}`),
			err:  plugin.NewTestConnError(codeInvalidSpec, assert.AnError),
			builderFunc: func() (plugin.Client, error) {
				return &Client{}, nil
			},
		},
		{
			name: "error/unreachable",
			spec: []byte(`{"connection_string": "file", "username": "user", "password": "pass"}`),
			err:  plugin.NewTestConnError(codeUnreachable, assert.AnError),
			builderFunc: func() (plugin.Client, error) {
				return nil, errUnreachable
			},
		},
		{
			name: "error/unauthorized",
			spec: []byte(`{"connection_string": "file", "username": "user", "password": "pass"}`),
			err:  plugin.NewTestConnError(codeUnauthorized, assert.AnError),
			builderFunc: func() (plugin.Client, error) {
				return nil, errUnauthorized
			},
		},
		{
			name: "error/connection_failed",
			spec: []byte(`{"connection_string": "file", "username": "user", "password": "pass"}`),
			err:  plugin.NewTestConnError(codeConnectionFailed, assert.AnError),
			builderFunc: func() (plugin.Client, error) {
				return nil, assert.AnError
			},
		},
	}

	for idx := range cases {
		tc := cases[idx]

		t.Run(tc.name, func(t *testing.T) {
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				spec := &Spec{}
				if err := json.Unmarshal(specBytes, spec); err != nil {
					return nil, errInvalidSpec
				}
				spec.SetDefaults()
				if err := spec.Validate(); err != nil {
					return nil, errInvalidSpec
				}

				return tc.builderFunc()
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
