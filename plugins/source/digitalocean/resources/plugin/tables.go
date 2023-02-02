package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/accounts"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/balances"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/billing_history"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/cdns"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/certificates"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/databases"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/domains"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/droplets"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/firewalls"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/floating_ips"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/images"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/keys"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/monitoring"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/projects"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/registries"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/sizes"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/snapshots"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/spaces"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/vpcs"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		accounts.Accounts(),
		cdns.Cdns(),
		billing_history.BillingHistory(),
		monitoring.AlertPolicies(),
		balances.Balances(),
		certificates.Certificates(),
		databases.Databases(),
		domains.Domains(),
		droplets.Droplets(),
		firewalls.Firewalls(),
		floating_ips.FloatingIps(),
		images.Images(),
		keys.Keys(),
		projects.Projects(),
		registries.Registries(),
		sizes.Sizes(),
		snapshots.Snapshots(),
		spaces.Spaces(),
		storage.Volumes(),
		vpcs.Vpcs(),
	}
}
