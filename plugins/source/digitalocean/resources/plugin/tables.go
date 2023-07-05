package plugin

import (
	"strings"

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
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/load_balancers"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/monitoring"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/projects"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/regions"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/registries"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/sizes"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/snapshots"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/spaces"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/services/vpcs"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getTables() []*schema.Table {
	tables := []*schema.Table{
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
		load_balancers.LoadBalancers(),
		projects.Projects(),
		regions.Regions(),
		registries.Registries(),
		sizes.Sizes(),
		snapshots.Snapshots(),
		spaces.Spaces(),
		storage.Volumes(),
		vpcs.Vpcs(),
	}

	for _, t := range tables {
		if err := t.Transform(t); err != nil {
			panic(err)
		}
		t.Title = titleTransformer(t)

		for _, rel := range t.Relations {
			if err := rel.Transform(rel); err != nil {
				panic(err)
			}
			rel.Title = titleTransformer(rel)
		}

		schema.AddCqIDs(t)
	}

	return tables
}

var customExceptions = map[string]string{
	"digitalocean": "DigitalOcean",
	"cdns":         "CDNs",
	"cors":         "CORS",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	csr := caser.New(caser.WithCustomExceptions(customExceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}
