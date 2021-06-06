package resources

import (
	"github.com/cloudquery/cq-provider-sdk/plugin"
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
	"github.com/cloudquery/cq-provider-template/client"
)

func Provider() *plugin.Provider {
	return &plugin.Provider{
		Name:      "your_provider_name",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"demo_resource": DemoResource(),
		},
		Config: func() interface{} {
			return &client.Config{}
		},
		DefaultConfigGenerator: func() (string, error) {
			return "", nil
		},
	}

}
