package client

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

func TestConnectionTester(t *testing.T) {
	tests := []struct {
		name          string
		spec          []byte
		err           *plugin.TestConnError
		clientBuilder func() (plugin.Client, error)
	}{
		{
			name: "ok",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  nil,
		},
		{
			name: "invalid spec",
			spec: []byte(`{"connection_string": "test"}`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "invalid spec JSON",
			spec: []byte(`{"connection_string": 12`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "connection failed",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  &plugin.TestConnError{Code: codeConnectionFailed},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errConnectionFailed
			},
		},
		{
			name: "unreachable",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  &plugin.TestConnError{Code: codeUnreachable},
			clientBuilder: func() (plugin.Client, error) {
				return nil, topology.ServerSelectionError{}
			},
		},
		{
			name: "unauthorized",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  &plugin.TestConnError{Code: codeUnauthorized},
			clientBuilder: func() (plugin.Client, error) {
				err := topology.ConnectionError{
					Wrapped: &auth.Error{},
				}
				return nil, err
			},
		},
		{
			name: "connection error without wrapped auth error",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  &plugin.TestConnError{Code: codeConnectionFailed},
			clientBuilder: func() (plugin.Client, error) {
				return nil, topology.ConnectionError{}
			},
		},
		{
			name: "unrecognized error",
			spec: []byte(`{"connection_string": "test", "database":"test"}`),
			err:  &plugin.TestConnError{Code: codeConnectionFailed},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errors.New("unrecognized error")
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.clientBuilder == nil {
				tt.clientBuilder = func() (plugin.Client, error) {
					return &Client{}, nil
				}
			}
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				sp := &spec.Spec{}
				if err := json.Unmarshal(specBytes, sp); err != nil {
					return nil, errInvalidSpec
				}
				sp.SetDefaults()
				if err := sp.Validate(); err != nil {
					return nil, errInvalidSpec
				}
				return tt.clientBuilder()
			})
			err := tester(context.Background(), zerolog.Nop(), tt.spec)
			if tt.err == nil {
				require.NoError(t, err)
				return
			}
			var expErr *plugin.TestConnError
			require.ErrorAs(t, err, &expErr)
			require.Equal(t, tt.err.Code, err.(*plugin.TestConnError).Code)
		})
	}
}
