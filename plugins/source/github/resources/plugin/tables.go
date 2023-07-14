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
	"github.com/cloudquery/cloudquery/plugins/source/github/resources/services/traffic"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func getTables() []*schema.Table {
	tables := []*schema.Table{
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
		traffic.Clones(),
		traffic.Paths(),
		traffic.Views(),
		traffic.Referrers(),
	}

	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}

	return tables
}
