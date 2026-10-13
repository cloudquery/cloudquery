package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNewNoConnection(t *testing.T) {
	ctx := context.Background()
	client, err := New(ctx, zerolog.Nop(), nil, plugin.NewClientOptions{NoConnection: true})
	require.NoError(t, err)
	require.NoError(t, client.Close(ctx))
}
