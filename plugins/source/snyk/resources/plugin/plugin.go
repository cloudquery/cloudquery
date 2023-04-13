package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/dependency"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/integration"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/organization"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/project"
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/services/reporting"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var Version = "Development"

func Snyk() *source.Plugin {
	return source.NewPlugin(
		"snyk",
		Version,
		[]*schema.Table{
			dependency.Dependencies(),
			integration.Integrations(),
			organization.Organizations(),
			project.Projects(),
			reporting.Issues(),
		},
		client.Configure,
	)
}
