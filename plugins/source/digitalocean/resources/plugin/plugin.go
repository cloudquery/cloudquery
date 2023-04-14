package plugin

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v2/caser"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"digitalocean": "DigitalOcean",
	"cdns":         "CDNs",
	"cors":         "CORS",
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
	t := csr.ToTitle(table.Name)
	return strings.Trim(strings.ReplaceAll(t, "  ", " "), " ")
}

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"digitalocean",
		Version,
		Tables(),
		client.New,
		source.WithTitleTransformer(titleTransformer),
	)
}
