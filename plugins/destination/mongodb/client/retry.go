package client

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// retryWrite runs op with exponential backoff on transient MongoDB errors. The
// MongoDB Go driver retries retryable writes once; this adds an extra layer to
// absorb longer bursts of connection instability (e.g. TCP broken pipe against
// MongoDB Atlas private-link endpoints). See ENG-3281.
//
// collection is used purely for log context so operators can tell which table
// is experiencing retries.
func retryWrite(ctx context.Context, logger zerolog.Logger, cfg *spec.WriteRetryConfig, collection string, op func() error) error {
	maxAttempts := cfg.GetMaxAttempts()
	if maxAttempts <= 1 {
		return op()
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = cfg.GetInitialBackoff()
	b.MaxInterval = cfg.GetMaxBackoff()

	start := time.Now()
	var attempts int
	_, err := backoff.Retry(ctx, func() (struct{}, error) {
		attempts++
		if err := op(); err != nil {
			if isRetryableWriteError(err) {
				return struct{}{}, err
			}
			return struct{}{}, backoff.Permanent(err)
		}
		return struct{}{}, nil
	},
		backoff.WithBackOff(b),
		backoff.WithMaxTries(uint(maxAttempts)),
		backoff.WithMaxElapsedTime(cfg.GetMaxElapsed()),
		backoff.WithNotify(func(err error, next time.Duration) {
			logger.Warn().
				Err(err).
				Str("collection", collection).
				Int("attempt", attempts).
				Int("max_attempts", maxAttempts).
				Dur("retry_after", next).
				Msg("retrying MongoDB write after transient error")
		}),
	)

	if err != nil {
		logger.Error().
			Err(err).
			Str("collection", collection).
			Int("attempts", attempts).
			Dur("elapsed", time.Since(start)).
			Msg("giving up on MongoDB write after retries")
		return err
	}
	if attempts > 1 {
		logger.Info().
			Str("collection", collection).
			Int("attempts", attempts).
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
	// The MongoDB server and driver attach the "RetryableWriteError" label to
	// write errors that are safe to retry (e.g. InterruptedAtShutdown,
	// NotWritablePrimary, PrimarySteppedDown). See
	// https://www.mongodb.com/docs/manual/core/retryable-writes/ for the spec.
	var labeled mongo.LabeledError
	if errors.As(err, &labeled) && labeled.HasErrorLabel("RetryableWriteError") {
		return true
	}
	return false
}
