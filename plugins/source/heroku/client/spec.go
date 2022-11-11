package client

import (
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// Spec defines Heroku source plugin Spec
type Spec struct {
	Token string `yaml:"token,omitempty" json:"token"`
}

type BackoffSettings struct {
	Gax        gax.Backoff
	Backoff    backoff.Config
	MaxRetries int
}

func (Spec) ClientOptions() []option.ClientOption {
	p := grpc.ConnectParams{
		Backoff: Backoff().Backoff,
	}
	return []option.ClientOption{
		option.WithGRPCDialOption(grpc.WithConnectParams(p)),
	}
}

func Backoff() BackoffSettings {
	b := BackoffSettings{
		Backoff:    backoff.DefaultConfig,
		MaxRetries: 3,
	}

	b.Gax.Initial = b.Backoff.BaseDelay
	b.Gax.Max = b.Backoff.MaxDelay
	b.Gax.Multiplier = b.Backoff.Multiplier

	return b
}
