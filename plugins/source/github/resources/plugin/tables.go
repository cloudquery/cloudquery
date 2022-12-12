package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/actions"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/external"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/hooks"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/installations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/issues"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/repositories"
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/teams"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Tables() []*schema.Table {
	return []*schema.Table{
		actions.Workflows(),
		billing.Action(),
		billing.Storage(),
		billing.Package(),
		external.Groups(),
		issues.Issues(),
		hooks.Hooks(),
		installations.Installations(),
		organizations.Organizations(),
		repositories.Repositories(),
		teams.Teams(),
	}
}
