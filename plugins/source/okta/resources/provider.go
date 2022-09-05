package resources

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "okta",
		Version:   Version,
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"users": Users(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
		ErrorClassifier: client.ErrorClassifier,
	}
}
