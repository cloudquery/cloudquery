package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/codegen"
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
			"account_features":           codegen.AccountFeatures(),
			"add_ons":                    codegen.AddOns(),
			"add_on_attachments":         codegen.AddOnAttachments(),
			"add_on_configs":             codegen.AddOnConfigs(),
			"add_on_region_capabilities": codegen.AddOnRegionCapabilities(),
			"add_on_services":            codegen.AddOnServices(),
			"add_on_webhooks":            codegen.AddOnWebhooks(),
			"add_on_webhook_deliveries":  codegen.AddOnWebhookDeliveries(),
			"add_on_webhook_events":      codegen.AddOnWebhookEvents(),
			"apps":                       codegen.Apps(),
			"app_features":               codegen.AppFeatures(),
			"app_transfers":              codegen.AppTransfers(),
			"app_webhooks":               codegen.AppWebhooks(),
			"app_webhook_deliveries":     codegen.AppWebhookDeliveries(),
			"app_webhook_events":         codegen.AppWebhookEvents(),
			"builds":                     codegen.Builds(),
			"buildpack_installations":    codegen.BuildpackInstallations(),
			"collaborators":              codegen.Collaborators(),
			"credits":                    codegen.Credits(),
			"domains":                    codegen.Domains(),
			"dynos":                      codegen.Dynos(),
			"dyno_sizes":                 codegen.DynoSizes(),
			"enterprise_accounts":        codegen.EnterpriseAccounts(),
			"enterprise_account_members": codegen.EnterpriseAccountMembers(),
			"formations":                 codegen.Formations(),
			"inbound_rulesets":           codegen.InboundRulesets(),
			"invoices":                   codegen.Invoices(),
			"keys":                       codegen.Keys(),
			"log_drains":                 codegen.LogDrains(),
			"oauth_authorizations":       codegen.OAuthAuthorizations(),
			"oauth_clients":              codegen.OAuthClients(),
			"outbound_rulesets":          codegen.OutboundRulesets(),
			"peerings":                   codegen.Peerings(),
			"permission_entities":        codegen.PermissionEntities(),
			"pipelines":                  codegen.Pipelines(),
			"pipeline_builds":            codegen.PipelineBuilds(),
			"pipeline_couplings":         codegen.PipelineCouplings(),
			"pipeline_deployments":       codegen.PipelineDeployments(),
			"pipeline_releases":          codegen.PipelineReleases(),
			"regions":                    codegen.Regions(),
			"releases":                   codegen.Releases(),
			"review_apps":                codegen.ReviewApps(),
			"spaces":                     codegen.Spaces(),
			"space_app_accesses":         codegen.SpaceAppAccesses(),
			"stacks":                     codegen.Stacks(),
			"team_app_permissions":       codegen.TeamAppPermissions(),
			"team_features":              codegen.TeamFeatures(),
			"teams":                      codegen.Teams(),
			"team_invitations":           codegen.TeamInvitations(),
			"team_invoices":              codegen.TeamInvoices(),
			"team_members":               codegen.TeamMembers(),
			"team_spaces":                codegen.TeamSpaces(),
			"vpn_connections":            codegen.VPNConnections(),
		},
		Config: func() sdkprovider.Config {
			return &client.Config{}
		},
	}
}
