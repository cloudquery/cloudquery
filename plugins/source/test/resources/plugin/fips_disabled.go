//go:build !fipsEnabled

package plugin

import "github.com/cloudquery/plugin-sdk/v4/plugin"

func buildTargets() []plugin.BuildTarget {
	return plugin.DefaultBuildTargets
}
