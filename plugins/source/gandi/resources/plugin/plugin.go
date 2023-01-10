package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/certificates"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/domains"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/livedns"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/resources/services/simplehosting"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

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
	)
}
