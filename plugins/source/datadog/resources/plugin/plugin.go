package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/caser"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "Development"

	customExceptions = map[string]string{
		"slo":  "SLO",
		"slos": "SLOs",
	}
)

func Plugin() *source.Plugin {
	allTables := Tables()
	// here you can append custom non-generated tables
	return source.NewPlugin(
		"datadog",
		Version,
		allTables,
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
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
