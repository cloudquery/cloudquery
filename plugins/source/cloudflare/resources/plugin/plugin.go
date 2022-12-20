package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/access_groups"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/accounts"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/certificate_packs"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/dns_records"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/images"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/ips"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/waf_overrides"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/waf_packages"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/worker_meta_data"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/worker_routes"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/resources/services/zones"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cloudflare",
		Version,
		[]*schema.Table{
			access_groups.AccessGroups(),
			accounts.Accounts(),
			certificate_packs.CertificatePacks(),
			dns_records.DNSRecords(),
			images.Images(),
			ips.IPs(),
			waf_packages.WAFPackages(),
			waf_overrides.WAFOverrides(),
			worker_meta_data.WorkerMetaData(),
			worker_routes.WorkerRoutes(),
			zones.Zones(),
		},
		client.Configure,
	)
}
