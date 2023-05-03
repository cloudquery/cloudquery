package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/groups"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/projects"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/settings"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/resources/services/users"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"gitlab",
		Version,
		schema.Tables{
			groups.Groups(),
			projects.Projects(),
			settings.Settings(),
			users.Users(),
		},
		client.Configure,
	)
}
