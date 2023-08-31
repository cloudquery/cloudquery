package client

import (
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	Logger     zerolog.Logger
	Spec       *Spec
	Homebrew   *homebrew.Client
	MaxRetries int
	Backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (*Client) ID() string {
	return "homebrew"
}
