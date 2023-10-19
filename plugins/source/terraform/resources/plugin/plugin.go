package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "terraform"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Terraform() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
