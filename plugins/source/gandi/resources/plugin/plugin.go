package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/certificates"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/domains"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/livedns"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/simplehosting"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

var customExceptions = map[string]string{
	"dnssec":        "DNSSEC",
	"livedns":       "LiveDNS",
	"simplehosting": "Simple Hosting",
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
		"gandi",
		Version,
		[]*schema.Table{
			certificates.Certificates(),
			certificates.CertificatePackages(),
			domains.Domains(),
			livedns.LiveDNSDomains(),
			simplehosting.SimplehostingInstances(),
		},
		client.Configure,
		source.WithTitleTransformer(titleTransformer),
	)
}
