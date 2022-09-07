package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/external_groups"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/hooks"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/installations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/issues"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/repositories"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/teams"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		billing.ActionBillings(),
		billing.StorageBillings(),
		billing.PackageBillings(),
		issues.Issues(),
		hooks.Hooks(),
		installations.Installations(),
		organizations.Organizations(),
		repositories.Repositories(),
		teams.Teams(),
		external_groups.ExternalGroups(),
	}
}
