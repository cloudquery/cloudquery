package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/users"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"gitlab",
		Version,
		[]*schema.Table{
			users.Groups(),
		},
		client.Configure,
	)
}
