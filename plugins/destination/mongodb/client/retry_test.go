package client

import (
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func newTestRetryConfig(maxAttempts int) *spec.WriteRetryConfig {
	d := configtype.NewDuration(10 * time.Millisecond)
	return &spec.WriteRetryConfig{
		MaxAttempts: maxAttempts,
		MaxBackoff:  &d,
	}
}

func TestRetryWrite_RetriesUntilSuccess(t *testing.T) {
	var calls int
	transientErr := io.ErrUnexpectedEOF
	err := retryWrite(context.Background(), zerolog.Nop(), newTestRetryConfig(5), "t", func() error {
		calls++
		if calls < 3 {
			return transientErr
		}
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, 3, calls)
}

func TestRetryWrite_StopsAtMaxAttempts(t *testing.T) {
	var calls int
	transientErr := io.ErrUnexpectedEOF
	err := retryWrite(context.Background(), zerolog.Nop(), newTestRetryConfig(3), "t", func() error {
		calls++
		return transientErr
	})
	require.ErrorIs(t, err, transientErr)
	require.Equal(t, 3, calls)
}

func TestRetryWrite_NonRetryableNotRetried(t *testing.T) {
	var calls int
	permanentErr := errors.New("schema validation failed")
	err := retryWrite(context.Background(), zerolog.Nop(), newTestRetryConfig(5), "t", func() error {
		calls++
		return permanentErr
	})
	require.ErrorIs(t, err, permanentErr)
	require.Equal(t, 1, calls)
}

func TestRetryWrite_DisabledByMaxAttemptsOne(t *testing.T) {
	var calls int
	err := retryWrite(context.Background(), zerolog.Nop(), newTestRetryConfig(1), "t", func() error {
		calls++
		return io.ErrUnexpectedEOF
	})
	require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	require.Equal(t, 1, calls)
}

func TestRetryWrite_NilConfigNoRetry(t *testing.T) {
	var calls int
	err := retryWrite(context.Background(), zerolog.Nop(), nil, "t", func() error {
		calls++
		return io.ErrUnexpectedEOF
	})
	require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	require.Equal(t, 1, calls)
}

func TestRetryWrite_ContextCancelStopsRetries(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var calls int
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	err := retryWrite(ctx, zerolog.Nop(), newTestRetryConfig(100), "t", func() error {
		calls++
		return io.ErrUnexpectedEOF
	})
	require.Error(t, err)
	require.Less(t, calls, 100)
}
