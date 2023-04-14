package client

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
)

func TestRetryOnRateLimitError(t *testing.T) {
	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := Client{
		logger:     logger,
		maxRetries: 1,
		backoff:    1 * time.Microsecond,
	}
	ctx := context.Background()
	t.Run("no_error", func(t *testing.T) {
		got := c.RetryOnError(ctx, "table_name", func() error {
			return nil
		})
		if got != nil {
			t.Errorf("RetryOnError returned error: %v, want nil", got)
		}
	})

	t.Run("with_error", func(t *testing.T) {
		got := c.RetryOnError(ctx, "table_name", func() error {
			return errors.New("test error")
		})
		if got == nil || got.Error() != "test error" {
			t.Errorf("RetryOnError returned error: %v, want %v", got, "test error")
		}
	})

	t.Run("retryable_error_that_never_succeeds", func(t *testing.T) {
		err := errors.New("server error: 500 Internal Server Error")
		got := c.RetryOnError(ctx, "table_name", func() error {
			return err
		})
		if got != err {
			t.Errorf("RetryOnError returned error: %v, want %v", got, err)
		}
	})

	t.Run("temporary_error", func(t *testing.T) {
		i := 0
		got := c.RetryOnError(ctx, "table_name", func() error {
			if i == 0 {
				i++
				return &snyk.ErrorResponse{
					Response: &snyk.Response{
						Response: &http.Response{
							StatusCode: http.StatusInternalServerError,
							Request: &http.Request{
								Method: "POST",
								URL: &url.URL{
									Scheme: "https",
									Host:   "test.test.test",
									Path:   "/api/v1/orgs/1234567890/projects",
								},
							},
						},
						SnykRequestID: "",
					},
					ErrorElement: snyk.ErrorElement{},
				}
			}
			return nil
		})
		if got != nil {
			t.Errorf("RetryOnError returned error: %v, want %v", got, nil)
		}
	})
}
