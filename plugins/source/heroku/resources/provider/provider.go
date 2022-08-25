package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/resources/services"
	sdkprovider "github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *sdkprovider.Provider {
	return &sdkprovider.Provider{
		Version:   Version,
		Name:      "heroku",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"account_features":           services.AccountFeatures(),
			"add_ons":                    services.AddOns(),
			"add_on_attachments":         services.AddOnAttachments(),
			"add_on_configs":             services.AddOnConfigs(),
			"add_on_region_capabilities": services.AddOnRegionCapabilities(),
			"add_on_services":            services.AddOnServices(),
			"add_on_webhooks":            services.AddOnWebhooks(),
			"add_on_webhook_deliveries":  services.AddOnWebhookDeliveries(),
			"add_on_webhook_events":      services.AddOnWebhookEvents(),
			"apps":                       services.Apps(),
			"app_features":               services.AppFeatures(),
			"app_transfers":              services.AppTransfers(),
			"app_webhooks":               services.AppWebhooks(),
			"app_webhook_deliveries":     services.AppWebhookDeliveries(),
			"app_webhook_events":         services.AppWebhookEvents(),
			"builds":                     services.Builds(),
			"buildpack_installations":    services.BuildpackInstallations(),
			"collaborators":              services.Collaborators(),
			"credits":                    services.Credits(),
			"domains":                    services.Domains(),
			"dynos":                      services.Dynos(),
			"dyno_sizes":                 services.DynoSizes(),
			"enterprise_accounts":        services.EnterpriseAccounts(),
			"enterprise_account_members": services.EnterpriseAccountMembers(),
			"formations":                 services.Formations(),
			"inbound_rulesets":           services.InboundRulesets(),
			"invoices":                   services.Invoices(),
			"keys":                       services.Keys(),
			"log_drains":                 services.LogDrains(),
			"oauth_authorizations":       services.OAuthAuthorizations(),
			"oauth_clients":              services.OAuthClients(),
			"outbound_rulesets":          services.OutboundRulesets(),
			"peerings":                   services.Peerings(),
			"permission_entities":        services.PermissionEntities(),
			"pipelines":                  services.Pipelines(),
			"pipeline_builds":            services.PipelineBuilds(),
			"pipeline_couplings":         services.PipelineCouplings(),
			"pipeline_deployments":       services.PipelineDeployments(),
			"pipeline_releases":          services.PipelineReleases(),
			"regions":                    services.Regions(),
			"releases":                   services.Releases(),
			"review_apps":                services.ReviewApps(),
			"spaces":                     services.Spaces(),
			"space_app_accesses":         services.SpaceAppAccesses(),
			"stacks":                     services.Stacks(),
			"team_app_permissions":       services.TeamAppPermissions(),
			"team_features":              services.TeamFeatures(),
			"teams":                      services.Teams(),
			"team_invitations":           services.TeamInvitations(),
			"team_invoices":              services.TeamInvoices(),
			"team_members":               services.TeamMembers(),
			"team_spaces":                services.TeamSpaces(),
			"vpn_connections":            services.VPNConnections(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
