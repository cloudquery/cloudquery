package client

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

func TestRetryOnRateLimitError(t *testing.T) {
	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	c := Client{
		logger: logger,
	}

	t.Run("no_error", func(t *testing.T) {
		got := c.RetryOnRateLimitError("table_name", func() error {
			return nil
		})
		if got != nil {
			t.Errorf("RetryOnRateLimitError returned error: %v, want nil", got)
		}
	})

	t.Run("with_error", func(t *testing.T) {
		got := c.RetryOnRateLimitError("table_name", func() error {
			return errors.New("test error")
		})
		if got == nil || got.Error() != "test error" {
			t.Errorf("RetryOnRateLimitError returned error: %v, want %v", got, "test error")
		}
	})

	t.Run("with_slack_error", func(t *testing.T) {
		calls := 0
		got := c.RetryOnRateLimitError("table_name", func() error {
			if calls == 0 {
				calls++
				return &slack.RateLimitedError{
					RetryAfter: time.Microsecond,
				}
			}
			return nil
		})
		if got != nil {
			t.Errorf("RetryOnRateLimitError returned error: %v, want nil", got)
		}
	})
}
