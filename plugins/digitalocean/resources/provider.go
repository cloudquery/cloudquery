package resources

import (
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "your_provider_name",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"demo_resource": DemoResource(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
