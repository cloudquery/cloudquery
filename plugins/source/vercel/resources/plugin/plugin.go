package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/deployment"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/domain"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/project"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/services/team"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"vercel",
		Version,
		[]*schema.Table{
			domain.Domains(),
			team.Teams(),
			project.Projects(),
			deployment.Deployments(),
		},
		client.Configure,
	)
}
