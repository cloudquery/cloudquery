package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# Required. OAuth token to authenticate with Heroku API
token: <token>

## Optional. GRPC Retry/backoff configuration, time units in seconds. Documented in https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
# backoff_base_delay: 1
# backoff_multiplier: 1.6
# backoff_max_delay: 120
# backoff_jitter: 0.2
# backoff_min_connect_timeout = 0
## Optional. Max amount of retries for retrier, defaults to max 3 retries.
# max_retries: 3
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"cloudflare",
		Version,
		[]*schema.Table{
			services.AccessGroups(),
			services.Accounts(),
			services.CertificatePacks(),
			services.DNSRecords(),
			services.Images(),
			services.Ips(),
			services.Wafs(),
			services.WafOverrides(),
			services.WorkersScripts(),
			services.WorkersRoutes(),
			services.Zones(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
