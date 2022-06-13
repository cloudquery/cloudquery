package provider

import (
	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:         Version,
		Name:            "digitalocean",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"droplets":        resources.Droplets(),
			"vpcs":            resources.Vpcs(),
			"sizes":           resources.Sizes(),
			"regions":         resources.Regions(),
			"keys":            resources.Keys(),
			"snapshots":       resources.Snapshots(),
			"account":         resources.Account(),
			"projects":        resources.Projects(),
			"balance":         resources.Balance(),
			"images":          resources.Images(),
			"domains":         resources.Domains(),
			"billing_history": resources.BillingHistory(),
			"volumes":         resources.Volumes(),
			"spaces":          resources.Spaces(),
			"floating_ips":    resources.FloatingIps(),
			"registry":        resources.Registries(),
			"databases":       resources.Databases(),
			"firewalls":       resources.Firewalls(),
			"cdns":            resources.Cdns(),
			"certificates":    resources.Certificates(),
			"load_balancers":  resources.LoadBalancers(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
