package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/resources/services"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "heroku",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"account_features":   services.AccountFeatures(),
			"add_ons":            services.AddOns(),
			"add_on_attachments": services.AddOnAttachments(),
			"stacks":             services.Stacks(),
			"teams":              services.Teams(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
