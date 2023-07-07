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
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"golang.org/x/exp/maps"
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

	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, titleTransformer); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}

	return tables
}

var customExceptions = map[string]string{
	"digitalocean": "DigitalOcean",
	"cdns":         "CDNs",
	"cors":         "CORS",
}

func titleTransformer(table *schema.Table) error {
	if table.Title != "" {
		return nil
	}
	exceptions := maps.Clone(docs.DefaultTitleExceptions)
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	table.Title = strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
	return nil
}
