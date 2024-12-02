package cmd

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
)

func SpecRegistryToPlugin(registry specs.Registry) managedplugin.Registry {
	switch registry {
	case specs.RegistryGitHub:
		return managedplugin.RegistryGithub
	case specs.RegistryLocal:
		return managedplugin.RegistryLocal
	case specs.RegistryGRPC:
		return managedplugin.RegistryGrpc
	case specs.RegistryDocker:
		return managedplugin.RegistryDocker
	case specs.RegistryCloudQuery:
		return managedplugin.RegistryCloudQuery
	default:
		panic(fmt.Sprintf("unknown registry %q", registry.String()))
	}
}
