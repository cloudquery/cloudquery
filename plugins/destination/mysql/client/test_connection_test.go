package client

import (
	"context"
	"encoding/json"
	"errors"
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
		errCode       string
		clientbuilder func() (plugin.Client, error)
	}{
		{
			name:    "ok",
			spec:    []byte(`{"connection_string": "connstr"}`),
			errCode: "",
		},
		{
			name:    "invalid spec",
			spec:    []byte(`{null}`),
			errCode: "INVALID_SPEC",
		},
		{
			name:    "connection failed",
			spec:    []byte(`{"connection_string": "connstr"}`),
			errCode: "CONNECT_FAILED",
			clientbuilder: func() (plugin.Client, error) {
				err := plugin.NewTestConnError("CONNECT_FAILED", errors.New("failed"))
				return nil, fmt.Errorf("failed to validate connection: %w", err)
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
					return nil, plugin.NewTestConnError("INVALID_SPEC", err)
				}
				sp.SetDefaults()
				if err := sp.Validate(); err != nil {
					return nil, plugin.NewTestConnError("INVALID_SPEC", err)
				}

				return tc.clientbuilder()
			})

			err := tester(context.Background(), zerolog.Nop(), tc.spec)
			if tc.errCode == "" {
				require.NoError(t, err)
				return
			}

			var expErr *plugin.TestConnError
			require.ErrorAs(t, err, &expErr)
			require.Equal(t, tc.errCode, err.(*plugin.TestConnError).Code)
		})
	}
}
