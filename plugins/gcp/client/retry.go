package client

import (
	"context"
	"fmt"
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/grpc/status"
)

// https://github.com/googleapis/google-api-go-client/issues/142
// https://github.com/googleapis/google-cloud-go/blob/1063c601a4c4a99217b45be0b25caa460e7157a1/internal/retry.go#L34
// Currently there is no easy way of configuring a custom retrier in the SDK (maybe will be in future clients)
// So we have to implement a wrapper

// Annotate prepends msg to the error message in err, attempting
// to preserve other information in err, like an error code.
//
// Annotate panics if err is nil.
//
// Annotate knows about these error types:
// - "google.golang.org/grpc/status".Status
// - "google.golang.org/api/googleapi".Error
// If the error is not one of these types, Annotate behaves
// like
//   fmt.Errorf("%s: %v", msg, err)
func annotate(err error, msg string) error {
	if err == nil {
		panic("Annotate called with nil")
	}
	if s, ok := status.FromError(err); ok {
		p := s.Proto()
		p.Message = msg + ": " + p.Message
		return status.ErrorProto(p)
	}
	if g, ok := err.(*googleapi.Error); ok {
		g.Message = msg + ": " + g.Message
		return g
	}
	return fmt.Errorf("%s: %v", msg, err)
}

// Annotatef uses format and args to format a string, then calls Annotate.
func annotatef(err error, format string, args ...interface{}) error {
	return annotate(err, fmt.Sprintf(format, args...))
}

func (c Client) RetryWithDefaultBackoffIgnoreErrors(ctx context.Context, f func() (stop bool, err error), ignoreCodes map[int]bool) error {
	return c.RetryWithDefaultBackoff(ctx, func() (stop bool, err error) {
		stop, err = f()
		if g, ok := err.(*googleapi.Error); ok && ignoreCodes[g.Code] {
			c.Logger().Debug(fmt.Sprintf("Retrying... Got %s", err))
			return false, err
		}
		return stop, err
	})
}

func (c Client) RetryWithDefaultBackoff(ctx context.Context, f func() (stop bool, err error)) error {
	return c.Retry(ctx, gax.Backoff{
		Initial: 60 * time.Second,
		Max:     5 * time.Minute,
	}, f)
}

//nolint:interfacer
func (c Client) Retry(ctx context.Context, bo gax.Backoff, f func() (stop bool, err error)) error {
	var lastErr error
	for {
		stop, err := f()
		if stop {
			return err
		}
		// Remember the last "real" error from f.
		if err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			lastErr = err
		}
		p := bo.Pause()
		if cerr := gax.Sleep(ctx, p); cerr != nil {
			if lastErr != nil {
				return annotatef(lastErr, "retry failed with %v; last error", cerr)
			}
			return cerr
		}
	}
}
