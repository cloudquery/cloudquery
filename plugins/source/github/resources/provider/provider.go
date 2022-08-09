package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/external_groups"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/hooks"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/installations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/issues"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/repositories"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/teams"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "github",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"billing.actions":  billing.ActionBillings(),
			"billing.storage":  billing.StorageBillings(),
			"billing.packages": billing.PackageBillings(),
			"issues":           issues.Issues(),
			"hooks":            hooks.Hooks(),
			"installations":    installations.Installations(),
			"organizations":    organizations.Organizations(),
			"repositories":     repositories.Repositories(),
			"teams":            teams.Teams(),
			"external_groups":  external_groups.ExternalGroups(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
