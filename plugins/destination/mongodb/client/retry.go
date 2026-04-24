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

// retryWrite runs op with exponential backoff on transient MongoDB errors. The
// MongoDB Go driver retries retryable writes once; this adds an extra layer to
// absorb longer bursts of connection instability (e.g. TCP broken pipe against
// MongoDB Atlas private-link endpoints). See ENG-3281.
//
// collection is used purely for log context so operators can tell which table
// is experiencing retries. cfg must already have defaults applied via
// spec.Spec.SetDefaults().
//
// Duplicate-write caveat: for tables without a primary key (InsertMany path),
// a retry triggered by an ambiguous failure -- server processed the write but
// response was lost -- can produce duplicate documents. The driver's built-in
// single retry avoids this via a session txnNumber that the server dedupes on,
// but that tuple is not exposed in the public API, so our app-level retries
// cannot participate in it. Upsert-keyed tables (with primary keys, BulkWrite
// path) are naturally idempotent and unaffected.
func retryWrite(ctx context.Context, logger zerolog.Logger, cfg *spec.WriteRetryConfig, collection string, op func() error) error {
	if cfg.MaxAttempts <= 1 {
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
	// The server/driver tag errors that are safe to retry with the
	// "RetryableWriteError" label (e.g. InterruptedAtShutdown,
	// NotWritablePrimary, PrimarySteppedDown, Atlas primary failovers). The
	// driver itself retries at most once on these; if one reaches us it means
	// the driver's single retry also failed and we should keep going. See
	// https://www.mongodb.com/docs/manual/core/retryable-writes/.
	var labeled mongo.LabeledError
	if errors.As(err, &labeled) && labeled.HasErrorLabel("RetryableWriteError") {
		return true
	}
	return false
}
