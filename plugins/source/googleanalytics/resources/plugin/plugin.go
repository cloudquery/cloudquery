package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/client"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
)

var Version = "Development"

var googleAnalyticsExceptions = map[string]string{
	"googleanalytics": "Google Analytics",
}

func titleTransformer(table *schema.Table) string {
	if table.Title != "" {
		return table.Title
	}
	exceptions := maps.Clone(source.DefaultTitleExceptions)
	for k, v := range googleAnalyticsExceptions {
		exceptions[k] = v
	}
	csr := caser.New(caser.WithCustomExceptions(exceptions))
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"googleanalytics",
		Version,
		nil,
		client.Configure,
		source.WithDynamicTableOption(client.GetTables),
		source.WithTitleTransformer(titleTransformer),
	)
}
