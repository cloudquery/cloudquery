package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/resources/ads"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/resources/campaigns"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/resources/customers"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
)

var Version = "Development"

var googleAdsExceptions = map[string]string{
	"googleads": "Google Ads",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := maps.Clone(source.DefaultTitleExceptions)
	for k, v := range googleAdsExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"googleads",
		Version,
		[]*schema.Table{
			ads.Groups(),
			campaigns.Campaigns(),
			customers.Customers(),
		},
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
