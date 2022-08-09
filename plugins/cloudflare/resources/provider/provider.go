package provider

import (
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/resources/services"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:            "cloudflare",
		Version:         Version,
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"access_groups":     services.AccessGroups(),
			"accounts":          services.Accounts(),
			"certificate_packs": services.CertificatePacks(),
			"dns_records":       services.DNSRecords(),
			"ips":               services.Ips(),
			"wafs":              services.Wafs(),
			"workers_scripts":   services.WorkersScripts(),
			"workers_routes":    services.WorkersRoutes(),
			"zones":             services.Zones(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
