package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/codegen"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# You may use either the API token or the API key. (API key requires setting the API email field)
# API token is preferred

# API token to access Cloudflare resources, also can be set with the CLOUDFLARE_API_TOKEN environment variable
api_token: "<YOUR_API_TOKEN_HERE>"
# API key to access Cloudflare resources, also can be set with the CLOUDFLARE_API_KEY environment variable
#api_key: "<YOUR_API_KEY_HERE>"
# API email to access Cloudflare resources, also can be set with the CLOUDFLARE_API_EMAIL environment variable
#api_email: "<YOUR_API_EMAIL_HERE>"

# List of accounts to target, if empty, all accounts will be targeted
#accounts:
# - "<YOUR_ACCOUNT_ID>"

# List of zones to target, if empty, all available zones will be targeted
#zones:
# - "<YOUR_ZONE_ID>"
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"cloudflare",
		Version,
		[]*schema.Table{
			codegen.AccessGroups(),
			codegen.Accounts(),
			codegen.CertificatePacks(),
			codegen.DNSRecords(),
			codegen.Images(),
			services.Ips(),
			codegen.WAFPackages(),
			codegen.WAFOverrides(),
			codegen.WorkerMetaData(),
			codegen.WorkerRoutes(),
			codegen.Zones(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
