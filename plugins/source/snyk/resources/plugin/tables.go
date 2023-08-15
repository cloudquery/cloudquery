package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/dependency"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/group"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/integration"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/organization"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/project"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/reporting"
	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func getTables() schema.Tables {
	tables := schema.Tables{
		dependency.Dependencies(),
		integration.Integrations(),
		group.Groups(),
		organization.Organizations(),
		project.Projects(),
		reporting.Issues(),
		reporting.LatestIssues(),
	}
	if err := transformers.Apply(tables, func(t *schema.Table) error {
		t.Title = docs.DefaultTitleTransformer(t)
		return nil
	}); err != nil {
		panic(err)
	}
	return client.Transform(tables)
}
