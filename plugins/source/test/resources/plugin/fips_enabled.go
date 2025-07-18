//go:build fipsEnabled

package plugin

import "github.com/cloudquery/plugin-sdk/v4/plugin"

func buildTargets() []plugin.BuildTarget {
	// default build targets for FIPS builds
	// fipsEnabled is used to enable FIPS mode
	targets := make([]plugin.BuildTarget, len(plugin.DefaultBuildTargets))
	for i := range plugin.DefaultBuildTargets {
		targets[i] = plugin.DefaultBuildTargets[i]
		targets[i].Tags = append(targets[i].Tags, "fipsEnabled")
	}
	return targets
}
