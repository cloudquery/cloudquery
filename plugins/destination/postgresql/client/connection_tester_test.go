package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	validConnectionString           = "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	invalidConnectionString         = "invalid"
	unknownHostConnectionString     = "postgresql://postgres:pass@unknownhost:5432/postgres?sslmode=disable"
	unknownDatabaseConnectionString = "postgresql://postgres:pass@localhost:5432/unknowndb?sslmode=disable"
	unknownUserConnectionString     = "postgresql://unknownuser:pass@localhost:5432/postgres?sslmode=disable"
	unknownPasswordConnectionString = "postgresql://postgres:unknownpass@localhost:5432/postgres?sslmode=disable"
)

func TestConnectionTester(t *testing.T) {
	type wantErr struct {
		Code             plugin.TestConnFailureCode
		ErrorDescription string
	}

	tests := []struct {
		name      string
		specBytes []byte
		wantErr   *wantErr
	}{
		{
			name:      "should return an error for an invalid spec",
			specBytes: []byte("invalid"),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeUnknown,
				ErrorDescription: "failed to unmarshal spec: invalid character 'i' looking for beginning of value",
			},
		},
		{
			name:      "should not return an error for a valid connection string",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: validConnectionString}),
		},
		{
			name:      "should return an error for an invalid connection string",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: invalidConnectionString}),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeInvalidCredentials,
				ErrorDescription: "cannot parse `invalid`: failed to parse as DSN (invalid dsn)",
			},
		},
		{
			name:      "should return an error for an unknown host",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: unknownHostConnectionString}),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeInvalidCredentials,
				ErrorDescription: "no such host \"unknownhost\"",
			},
		},
		{
			name:      "should return an error for an unknown database",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: unknownDatabaseConnectionString}),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeInvalidCredentials,
				ErrorDescription: "database \"unknowndb\" does not exist",
			},
		},
		{
			name:      "should return an error for an unknown user",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: unknownUserConnectionString}),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeInvalidCredentials,
				ErrorDescription: "password authentication failed for user \"unknownuser\"",
			},
		},
		{
			name:      "should return an error for an unknown password",
			specBytes: marshalSpec(t, &spec.Spec{ConnectionString: unknownPasswordConnectionString}),
			wantErr: &wantErr{
				Code:             plugin.TestConnFailureCodeInvalidCredentials,
				ErrorDescription: "password authentication failed for user \"postgres\"",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := plugin.NewPlugin(
				"postgresql",
				"development",
				New,
				plugin.WithConnectionTester(ConnectionTester),
			)

			logger := zerolog.New(os.Stdout)

			err := p.TestConnection(context.Background(), logger, tt.specBytes)
			if tt.wantErr != nil {
				var target *plugin.TestConnError
				require.Error(t, err)
				require.ErrorAs(t, err, &target)
				assert.Equal(t, tt.wantErr.Code, target.Code)
				assert.Equal(t, tt.wantErr.ErrorDescription, target.Message.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func marshalSpec(t *testing.T, s *spec.Spec) []byte {
	b, err := json.Marshal(s)
	require.NoError(t, err)

	return b
}
