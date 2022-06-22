package resources

import (
	"github.com/cloudquery/cq-provider-okta/client"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "okta",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"users": Users(),
		},
		Config: func(f cqproto.ConfigFormat) provider.Config {
			return client.NewConfig(f)
		},
		ErrorClassifier: client.ErrorClassifier,
	}
}
