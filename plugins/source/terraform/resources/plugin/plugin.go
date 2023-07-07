package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var Version = "development"

func Terraform() *plugin.Plugin {
	return plugin.NewPlugin(
		"terraform",
		Version,
		configure,
	)
}
