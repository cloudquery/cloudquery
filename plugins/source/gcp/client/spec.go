package client

import (
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// Spec defines GCP source plugin Spec
type Spec struct {
	ProjectIDs            []string `yaml:"project_ids,omitempty" json:"project_ids"`
	FolderIDs             []string `yaml:"folder_ids,omitempty" json:"folder_ids"`
	FolderMaxDepth        uint     `yaml:"folders_max_depth,omitempty" json:"folders_max_depth"`
	ServiceAccountKeyJSON string   `yaml:"service_account_key_json,omitempty" json:"service_account_key_json"`

	BaseDelay         int     `yaml:"backoff_base_delay,omitempty" hcl:"backoff_base_delay,optional" default:"-1"`
	Multiplier        float64 `yaml:"backoff_multiplier,omitempty" hcl:"backoff_multiplier,optional"`
	MaxDelay          int     `yaml:"backoff_max_delay,omitempty" hcl:"backoff_max_delay,optional"`
	Jitter            float64 `yaml:"backoff_jitter,omitempty" hcl:"backoff_jitter,optional"`
	MinConnectTimeout int     `yaml:"backoff_min_connect_timeout,omitempty" hcl:"backoff_min_connect_timeout,optional"`
	MaxRetries        int     `yaml:"max_retries,omitempty" hcl:"max_retries,optional" default:"3"`
}

type BackoffSettings struct {
	Gax        gax.Backoff
	Backoff    backoff.Config
	MaxRetries int
}

func (c Spec) ClientOptions() []option.ClientOption {
	p := grpc.ConnectParams{
		Backoff: c.Backoff().Backoff,
	}
	if c.MinConnectTimeout >= 0 {
		p.MinConnectTimeout = time.Duration(c.MinConnectTimeout) * time.Second
	}
	return []option.ClientOption{
		option.WithGRPCDialOption(grpc.WithConnectParams(p)),
	}
}

func (c Spec) Backoff() BackoffSettings {
	b := BackoffSettings{
		Backoff:    backoff.DefaultConfig,
		MaxRetries: 3,
	}
	if c.BaseDelay >= 0 {
		b.Backoff.BaseDelay = time.Duration(c.BaseDelay) * time.Second
	}
	if c.Multiplier > 0 {
		b.Backoff.Multiplier = c.Multiplier
	}
	if c.MaxDelay > 0 {
		b.Backoff.MaxDelay = time.Duration(c.MaxDelay) * time.Second
	}
	if c.Jitter != 0 {
		b.Backoff.Jitter = c.Jitter
	}
	if c.MaxRetries != 0 {
		b.MaxRetries = c.MaxRetries
	}

	b.Gax.Initial = b.Backoff.BaseDelay
	b.Gax.Max = b.Backoff.MaxDelay
	b.Gax.Multiplier = b.Backoff.Multiplier

	return b
}
