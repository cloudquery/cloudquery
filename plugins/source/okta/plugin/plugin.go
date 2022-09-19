package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# Optional. Okta Token to access API, you can set this with OKTA_API_TOKEN env variable
# token: "<YOUR_OKTA_TOKEN>"
# Required. You okta domain name
# domain: "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
`

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"okta",
		Version,
		[]*schema.Table{
			services.Users(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
