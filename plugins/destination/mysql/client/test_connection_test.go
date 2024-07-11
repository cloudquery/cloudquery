package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConnectionTester(t *testing.T) {
	tests := []struct {
		name          string
		spec          []byte
		err           *plugin.TestConnError
		clientbuilder func() (plugin.Client, error)
	}{
		{
			name: "ok",
			spec: []byte(`{"connection_string": "connstr"}`),
			err:  nil,
		},
		{
			name: "invalid spec",
			spec: []byte(`{null}`),
			err:  plugin.NewTestConnError(codeInvalidSpec, nil),
		},
		{
			name: "connection failed",
			spec: []byte(`{"connection_string": "connstr"}`),
			err:  plugin.NewTestConnError(codeConnectionFailed, errValidateConnectionFailed),
			clientbuilder: func() (plugin.Client, error) {
				return nil, errValidateConnectionFailed
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.clientbuilder == nil {
				tc.clientbuilder = func() (plugin.Client, error) {
					return &Client{}, nil
				}
			}
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				sp := &Spec{}
				if err := json.Unmarshal(specBytes, &sp); err != nil {
					return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
				}
				sp.SetDefaults()
				if err := sp.Validate(); err != nil {
					return nil, err
				}

				return tc.clientbuilder()
			})

			err := tester(context.Background(), zerolog.Nop(), tc.spec)
			if tc.err == nil {
				require.NoError(t, err)
				return
			}

			var expErr *plugin.TestConnError
			require.ErrorAs(t, err, &expErr)
			require.Equal(t, tc.err.Code, err.(*plugin.TestConnError).Code)
		})
	}
}
