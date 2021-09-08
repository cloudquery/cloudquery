package resources

import (
	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "digitalocean",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"droplets":        Droplets(),
			"vpcs":            Vpcs(),
			"sizes":           Sizes(),
			"regions":         Regions(),
			"keys":            Keys(),
			"snapshots":       Snapshots(),
			"account":         Account(),
			"projects":        Projects(),
			"balance":         Balance(),
			"images":          Images(),
			"domains":         Domains(),
			"billing_history": BillingHistory(),
			"volumes":         Volumes(),
			"spaces":          Spaces(),
			"floating_ips":    FloatingIps(),
			"registry":        Registries(),
			"databases":       Databases(),
			"firewalls":       Firewalls(),
			"cdns":            Cdns(),
			"certificates":    Certificates(),
			"load_balancers":  LoadBalancers(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
