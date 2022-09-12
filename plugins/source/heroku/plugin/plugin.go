package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/codegen"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# Required. OAuth token to authenticate with Heroku API
token: <token>
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"heroku",
		Version,
		[]*schema.Table{
			codegen.AccountFeatures(),
			codegen.AddOns(),
			codegen.AddOnAttachments(),
			codegen.AddOnConfigs(),
			codegen.AddOnRegionCapabilities(),
			codegen.AddOnServices(),
			codegen.AddOnWebhooks(),
			codegen.AddOnWebhookDeliveries(),
			codegen.AddOnWebhookEvents(),
			codegen.Apps(),
			codegen.AppFeatures(),
			codegen.AppTransfers(),
			codegen.AppWebhooks(),
			codegen.AppWebhookDeliveries(),
			codegen.AppWebhookEvents(),
			codegen.Builds(),
			codegen.BuildpackInstallations(),
			codegen.Collaborators(),
			codegen.Credits(),
			codegen.Domains(),
			codegen.Dynos(),
			codegen.DynoSizes(),
			codegen.EnterpriseAccounts(),
			codegen.EnterpriseAccountMembers(),
			codegen.Formations(),
			codegen.InboundRulesets(),
			codegen.Invoices(),
			codegen.Keys(),
			codegen.LogDrains(),
			codegen.OAuthAuthorizations(),
			codegen.OAuthClients(),
			codegen.OutboundRulesets(),
			codegen.Peerings(),
			codegen.PermissionEntities(),
			codegen.Pipelines(),
			codegen.PipelineBuilds(),
			codegen.PipelineCouplings(),
			codegen.PipelineDeployments(),
			codegen.PipelineReleases(),
			codegen.Regions(),
			codegen.Releases(),
			codegen.ReviewApps(),
			codegen.Spaces(),
			codegen.SpaceAppAccesses(),
			codegen.Stacks(),
			codegen.TeamAppPermissions(),
			codegen.TeamFeatures(),
			codegen.Teams(),
			codegen.TeamInvitations(),
			codegen.TeamInvoices(),
			codegen.TeamMembers(),
			codegen.TeamSpaces(),
			codegen.VPNConnections(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
