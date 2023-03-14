package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

var customExceptions = map[string]string{
	"pagerduty": "PagerDuty",
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

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"pagerduty",
		Version,
		AllTables(),
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
