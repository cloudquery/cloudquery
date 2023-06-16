package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/applications"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/users"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"okta",
		Version,
		[]*schema.Table{
			users.Users(),
			groups.Groups(),
			applications.Applications(),
		},
		client.Configure,
	)
}
