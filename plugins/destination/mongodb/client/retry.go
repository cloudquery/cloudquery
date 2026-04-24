package client

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	retryMaxTries       = 5
	retryInitialBackoff = 500 * time.Millisecond
	retryMaxBackoff     = 10 * time.Second
	retryMaxElapsedTime = 30 * time.Second
)

// retryWrite runs op with exponential backoff on transient MongoDB network
// errors. The MongoDB Go driver retries retryable writes once; this adds an
// extra layer to absorb longer bursts of connection instability (e.g. TCP
// broken pipe against MongoDB Atlas private-link endpoints).
func retryWrite(ctx context.Context, logger zerolog.Logger, op func() error) error {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = retryInitialBackoff
	b.MaxInterval = retryMaxBackoff

	_, err := backoff.Retry(ctx, func() (struct{}, error) {
		if err := op(); err != nil {
			if isRetryableWriteError(err) {
				return struct{}{}, err
			}
			return struct{}{}, backoff.Permanent(err)
		}
		return struct{}{}, nil
	},
		backoff.WithBackOff(b),
		backoff.WithMaxTries(retryMaxTries),
		backoff.WithMaxElapsedTime(retryMaxElapsedTime),
		backoff.WithNotify(func(err error, next time.Duration) {
			logger.Warn().Err(err).Dur("retry_after", next).Msg("retrying MongoDB write after transient error")
		}),
	)
	return err
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
