package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/resources/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"heroku",
		Version,
		[]*schema.Table{
			services.AccountFeatures(),
			services.AddOns(),
			services.AddOnAttachments(),
			services.AddOnConfigs(),
			services.AddOnRegionCapabilities(),
			services.AddOnServices(),
			services.AddOnWebhooks(),
			services.AddOnWebhookDeliveries(),
			services.AddOnWebhookEvents(),
			services.Apps(),
			services.AppFeatures(),
			services.AppTransfers(),
			services.AppWebhooks(),
			services.AppWebhookDeliveries(),
			services.AppWebhookEvents(),
			services.Builds(),
			services.BuildpackInstallations(),
			services.Collaborators(),
			services.Credits(),
			services.Domains(),
			services.Dynos(),
			services.DynoSizes(),
			services.EnterpriseAccounts(),
			services.EnterpriseAccountMembers(),
			services.Formations(),
			services.InboundRulesets(),
			services.Invoices(),
			services.Keys(),
			services.LogDrains(),
			services.OAuthAuthorizations(),
			services.OAuthClients(),
			services.OutboundRulesets(),
			services.Peerings(),
			services.PermissionEntities(),
			services.Pipelines(),
			services.PipelineBuilds(),
			services.PipelineCouplings(),
			services.PipelineDeployments(),
			services.PipelineReleases(),
			services.Regions(),
			services.Releases(),
			services.ReviewApps(),
			services.Spaces(),
			services.SpaceAppAccesses(),
			services.Stacks(),
			services.TeamAppPermissions(),
			services.TeamFeatures(),
			services.Teams(),
			services.TeamInvitations(),
			services.TeamInvoices(),
			services.TeamMembers(),
			services.TeamSpaces(),
			services.VPNConnections(),
		},
		client.Configure,
	)
}
