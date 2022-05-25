package client

import (
	"time"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// Config defines Provider Configuration
type Config struct {
	ProjectFilter         string   `hcl:"project_filter,optional"` // Deprecated
	ProjectIDs            []string `hcl:"project_ids,optional"`
	FolderIDs             []string `hcl:"folder_ids,optional"`
	FolderMaxDepth        uint     `hcl:"folders_max_depth,optional"`
	ServiceAccountKeyJSON string   `hcl:"service_account_key_json,optional"`

	BaseDelay         int     `hcl:"backoff_base_delay,optional" default:"-1"`
	Multiplier        float64 `hcl:"backoff_multiplier,optional"`
	MaxDelay          int     `hcl:"backoff_max_delay,optional"`
	Jitter            float64 `hcl:"backoff_jitter,optional"`
	MinConnectTimeout int     `hcl:"backoff_min_connect_timeout,optional"`
	MaxRetries        int     `hcl:"max_retries,optional" default:"3"`
}

type BackoffSettings struct {
	Gax        gax.Backoff
	Backoff    backoff.Config
	MaxRetries int
}

func (Config) Example() string {
	return `configuration {
				// Optional. List of folders to get projects from. Required permission: resourcemanager.projects.list
				// folder_ids = [ "organizations/<ORG_ID>", "folders/<FOLDER_ID>" ]
				// Optional. Maximum level of folders to recurse into
				// folders_max_depth = 5
				// Optional. If not specified either using all projects accessible.
				// project_ids = [<CHANGE_THIS_TO_YOUR_PROJECT_ID>]
				// Optional. ServiceAccountKeyJSON passed as value instead of a file path, can be passed also via env: CQ_SERVICE_ACCOUNT_KEY_JSON
				// service_account_key_json = <YOUR_JSON_SERVICE_ACCOUNT_KEY_DATA>
				// Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
				// backoff_base_delay = 1
				// backoff_multiplier = 1.6
				// backoff_max_delay = 120
				// backoff_jitter = 0.2
				// backoff_min_connect_timeout = 0
				// Optional. Max amount of retries for retrier, defaults to max 3 retries.
				// max_retries = 3
			}`
}

func (c Config) ClientOptions() []option.ClientOption {
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

func (c Config) Backoff() BackoffSettings {
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
