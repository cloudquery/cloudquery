package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/applications"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/users"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

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
