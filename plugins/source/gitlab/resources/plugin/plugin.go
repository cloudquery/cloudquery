package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/groups"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/projects"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/settings"
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
			groups.Groups(),
			projects.Projects(),
			settings.Settings(),
		},
		client.Configure,
	)
}
