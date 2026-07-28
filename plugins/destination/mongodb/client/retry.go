package client

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/avast/retry-go/v5"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func retryWrite(ctx context.Context, logger zerolog.Logger, cfg *spec.WriteRetryConfig, collection string, op func() error) error {
	if cfg == nil || cfg.MaxAttempts <= 1 || cfg.MaxBackoff == nil {
		return op()
	}

	start := time.Now()
	var attempts uint

	err := retry.New(
		retry.Context(ctx),
		retry.Attempts(uint(cfg.MaxAttempts)),
		retry.MaxDelay(cfg.MaxBackoff.Duration()),
		retry.LastErrorOnly(true),
		retry.RetryIf(isRetryableWriteError),
		retry.OnRetry(func(n uint, err error) {
			logger.Warn().
				Err(err).
				Str("collection", collection).
				Uint("attempt", n+1).
				Int("max_attempts", cfg.MaxAttempts).
				Msg("retrying MongoDB write after transient error")
		}),
	).Do(func() error {
		attempts++
		return op()
	})

	if err != nil {
		logger.Error().
			Err(err).
			Str("collection", collection).
			Uint("attempts", attempts).
			Dur("elapsed", time.Since(start)).
			Msg("giving up on MongoDB write after retries")
		return err
	}
	if attempts > 1 {
		logger.Info().
			Str("collection", collection).
			Uint("attempts", attempts).
			Dur("elapsed", time.Since(start)).
			Msg("MongoDB write succeeded after retries")
	}
	return nil
}

func isRetryableWriteError(err error) bool {
	if err == nil {
		return false
	}
	if mongo.IsNetworkError(err) || mongo.IsTimeout(err) {
		return true
	}
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return true
	}
	var labeled mongo.LabeledError
	if errors.As(err, &labeled) && labeled.HasErrorLabel("RetryableWriteError") {
		return true
	}
	var retryable interface{ Retryable() bool }
	if errors.As(err, &retryable) && retryable.Retryable() {
		return true
	}
	return false
}
