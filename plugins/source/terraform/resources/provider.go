package resources

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "terraform",
		Version:   Version,
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"tf.data": TFData(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
