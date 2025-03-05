package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/destination/firehose/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			spec: []byte(`{"stream_arn":"arn:aws:firehose:::deliverystream/test"}`),
			clientBuilder: func() (plugin.Client, error) {
				return &plugin.Plugin{}, nil
			},
		},
		{
			name: "invalid spec",
			spec: []byte(`{sad}`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "invalid stream arn",
			spec: []byte(`{"stream_arn":"arn:test"}`),
			err:  &plugin.TestConnError{Code: codeInvalidSpec},
		},
		{
			name: "unauthorized",
			spec: []byte(`{"stream_arn":"arn:aws:firehose:::deliverystream/test"}`),
			err:  &plugin.TestConnError{Code: codeUnauthorized},
			clientBuilder: func() (plugin.Client, error) {
				return nil, errUnauthorized
			},
		},
		{
			name: "other errors",
			spec: []byte(`{"stream_arn":"arn:aws:firehose:::deliverystream/test"}`),
			err:  &plugin.TestConnError{Code: codeConnectionFailed},
			clientBuilder: func() (plugin.Client, error) {
				return nil, assert.AnError
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tester := NewConnectionTester(func(_ context.Context, _ zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
				sp := &spec.Spec{}
				if err := json.Unmarshal(specBytes, sp); err != nil {
					return nil, errInvalidSpec
				}
				sp.SetDefaults()
				if err := sp.Validate(); err != nil {
					return nil, errInvalidSpec
				}
				if _, err := arn.Parse(sp.StreamARN); err != nil {
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
			require.Equal(t, tt.err.Code, expErr.Code)
		})
	}
}
