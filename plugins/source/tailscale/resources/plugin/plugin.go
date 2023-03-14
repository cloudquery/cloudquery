package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var Version = "Development"

var customExceptions = map[string]string{
	"acls":        "Access Control Lists (ACLs)",
	"dns":         "Domain Name System (DNS)",
	"nameservers": "Name Servers",
	"searchpaths": "Search Paths",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := make(map[string]string)
	for k, v := range source.DefaultTitleExceptions {
		exceptions[k] = v
	}
	for k, v := range customExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	return csr.ToTitle(table.Name)
}

func Tailscale() *source.Plugin {
	return source.NewPlugin(
		"tailscale",
		Version,
		tables(),
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
