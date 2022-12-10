// Code generated by codegen2; DO NOT EDIT.
package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/aad"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/advisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/analysisservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/apimanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/appcomplianceautomation"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/appconfiguration"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/applicationinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/appplatform"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/appservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/automanage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/automation"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/avs"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/azurearcdata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/azuredata"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/billing"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/botservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cognitiveservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/confluent"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/connectedvmware"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/containerinstance"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/containerregistry"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/containerservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/costmanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/customerinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/dashboard"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/databox"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datadog"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datafactory"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datalakeanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datalakestore"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datamigration"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/desktopvirtualization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/devhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/devops"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/dns"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/dnsresolver"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/elastic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/engagementfabric"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/eventgrid"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/eventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/frontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/hanaonazure"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/hdinsight"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/healthbot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/healthcareapis"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/hybridcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/hybriddatamanager"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/kusto"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/maintenance"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/managementgroups"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/marketplace"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/migrate"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysqlflexibleservers"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/netapp"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/networkfunction"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/nginx"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/notificationhubs"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/operationalinsights"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/orbital"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/peering"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/policy"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/portal"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresqlflexibleservers"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresqlhsc"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/powerbidedicated"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/privatedns"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/providerhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/redhatopenshift"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/relay"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/reservations"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/saas"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sqlvirtualmachine"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storagecache"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/support"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/synapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/windowsesu"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/windowsiot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/workloads"
	"github.com/cloudquery/plugin-sdk/schema"
)

func generatedTables() []*schema.Table {
	return []*schema.Table{
		aad.PrivateLinkForAzureAd(),
		advisor.RecommendationMetadata(),
		advisor.Recommendations(),
		advisor.Suppressions(),
		analysisservices.Servers(),
		apimanagement.Service(),
		appcomplianceautomation.Reports(),
		appconfiguration.ConfigurationStores(),
		applicationinsights.Components(),
		applicationinsights.WebTests(),
		appplatform.Services(),
		appservice.CertificateOrders(),
		appservice.Certificates(),
		appservice.DeletedWebApps(),
		appservice.Domains(),
		appservice.Environments(),
		appservice.Plans(),
		appservice.Recommendations(),
		appservice.ResourceHealthMetadata(),
		appservice.StaticSites(),
		appservice.TopLevelDomains(),
		appservice.WebApps(),
		authorization.ClassicAdministrators(),
		authorization.ProviderOperationsMetadata(),
		authorization.RoleAssignments(),
		authorization.RoleDefinitions(),
		automanage.ConfigurationProfileAssignments(),
		automation.Account(),
		avs.PrivateClouds(),
		azurearcdata.PostgresInstances(),
		azurearcdata.SqlManagedInstances(),
		azurearcdata.SqlServerInstances(),
		azuredata.SqlServerRegistrations(),
		batch.Account(),
		billing.Accounts(),
		billing.EnrollmentAccounts(),
		billing.Periods(),
		botservice.Bots(),
		cdn.EdgeNodes(),
		cdn.ManagedRuleSets(),
		cdn.Policies(),
		cdn.Profiles(),
		cognitiveservices.Accounts(),
		cognitiveservices.DeletedAccounts(),
		compute.AvailabilitySets(),
		compute.CloudServices(),
		compute.DiskAccesses(),
		compute.DiskEncryptionSets(),
		compute.Disks(),
		compute.Galleries(),
		compute.Images(),
		compute.RestorePointCollections(),
		compute.Snapshots(),
		compute.VirtualMachineScaleSets(),
		compute.VirtualMachines(),
		confluent.MarketplaceAgreements(),
		connectedvmware.Clusters(),
		connectedvmware.Datastores(),
		connectedvmware.Hosts(),
		connectedvmware.ResourcePools(),
		connectedvmware.VCenters(),
		connectedvmware.VirtualMachineTemplates(),
		connectedvmware.VirtualMachines(),
		connectedvmware.VirtualNetworks(),
		containerinstance.ContainerGroups(),
		containerregistry.Registries(),
		containerservice.ManagedClusters(),
		containerservice.Snapshots(),
		cosmos.Locations(),
		cosmos.RestorableDatabaseAccounts(),
		cosmos.DatabaseAccounts(),
		costmanagement.Views(),
		customerinsights.Hubs(),
		dashboard.Grafana(),
		databox.Jobs(),
		datadog.MarketplaceAgreements(),
		datadog.Monitors(),
		datafactory.Factories(),
		datalakeanalytics.Accounts(),
		datalakestore.Accounts(),
		datamigration.Services(),
		desktopvirtualization.HostPools(),
		devhub.Workflow(),
		devops.PipelineTemplateDefinitions(),
		dns.Zones(),
		dnsresolver.DnsForwardingRulesets(),
		dnsresolver.DnsResolvers(),
		elastic.Monitors(),
		engagementfabric.Accounts(),
		eventgrid.TopicTypes(),
		eventhub.Namespaces(),
		frontdoor.FrontDoors(),
		frontdoor.ManagedRuleSets(),
		frontdoor.NetworkExperimentProfiles(),
		frontdoor.Policies(),
		hanaonazure.SapMonitors(),
		hdinsight.Clusters(),
		healthbot.Bots(),
		healthcareapis.Services(),
		hybridcompute.PrivateLinkScopes(),
		hybriddatamanager.DataManagers(),
		kusto.Clusters(),
		maintenance.Configurations(),
		maintenance.PublicMaintenanceConfigurations(),
		managementgroups.Entities(),
		mariadb.Servers(),
		marketplace.PrivateStore(),
		migrate.Projects(),
		monitor.LogProfiles(),
		monitor.PrivateLinkScopes(),
		monitor.TenantActivityLogs(),
		mysql.Servers(),
		mysqlflexibleservers.Servers(),
		netapp.Accounts(),
		network.ApplicationGateways(),
		network.ApplicationSecurityGroups(),
		network.AzureFirewalls(),
		network.BastionHosts(),
		network.BgpServiceCommunities(),
		network.CustomIpPrefixes(),
		network.DdosProtectionPlans(),
		network.DscpConfiguration(),
		network.ExpressRouteCircuits(),
		network.ExpressRoutePorts(),
		network.ExpressRoutePortsLocations(),
		network.ExpressRouteServiceProviders(),
		network.FirewallPolicies(),
		network.IpAllocations(),
		network.IpGroups(),
		network.Interfaces(),
		network.LoadBalancers(),
		network.LocalNetworkGateways(),
		network.Managers(),
		network.NatGateways(),
		network.PrivateEndpoints(),
		network.PrivateLinkServices(),
		network.Profiles(),
		network.PublicIpAddresses(),
		network.PublicIpPrefixes(),
		network.RouteFilters(),
		network.RouteTables(),
		network.SecurityGroups(),
		network.SecurityPartnerProviders(),
		network.ServiceEndpointPolicies(),
		network.SubscriptionNetworkManagerConnections(),
		network.VpnGateways(),
		network.VpnServerConfigurations(),
		network.VpnSites(),
		network.VirtualAppliances(),
		network.VirtualHubs(),
		network.VirtualNetworkGatewayConnections(),
		network.VirtualNetworkGateways(),
		network.VirtualNetworks(),
		network.VirtualRouters(),
		network.VirtualWans(),
		network.Watchers(),
		network.WebApplicationFirewallPolicies(),
		networkfunction.AzureTrafficCollectorsBySubscription(),
		nginx.Deployments(),
		notificationhubs.Namespaces(),
		operationalinsights.Clusters(),
		operationalinsights.Workspaces(),
		orbital.ContactProfiles(),
		orbital.Spacecrafts(),
		peering.ServiceCountries(),
		peering.ServiceLocations(),
		peering.ServiceProviders(),
		policy.Assignments(),
		policy.DataPolicyManifests(),
		policy.Definitions(),
		policy.Exemptions(),
		policy.SetDefinitions(),
		portal.ListTenantConfigurationViolations(),
		portal.TenantConfigurations(),
		postgresql.Servers(),
		postgresqlflexibleservers.Servers(),
		postgresqlhsc.ServerGroups(),
		powerbidedicated.Capacities(),
		privatedns.PrivateZones(),
		providerhub.ProviderRegistrations(),
		redhatopenshift.OpenShiftClusters(),
		relay.Namespaces(),
		reservations.ReservationOrder(),
		saas.Applications(),
		saas.Resources(),
		security.Alerts(),
		security.AlertsSuppressionRules(),
		security.AllowedConnections(),
		security.Applications(),
		security.AssessmentsMetadata(),
		security.AutoProvisioningSettings(),
		security.Automations(),
		security.Connectors(),
		security.DiscoveredSecuritySolutions(),
		security.ExternalSecuritySolutions(),
		security.GovernanceRule(),
		security.JitNetworkAccessPolicies(),
		security.Locations(),
		security.RegulatoryComplianceStandards(),
		security.SecureScoreControlDefinitions(),
		security.SecureScoreControls(),
		security.SecureScores(),
		security.Solutions(),
		security.Tasks(),
		security.Topology(),
		security.WorkspaceSettings(),
		servicebus.Namespaces(),
		sql.DeletedServers(),
		sql.InstancePools(),
		sql.ManagedInstances(),
		sql.Servers(),
		sql.VirtualClusters(),
		sqlvirtualmachine.Groups(),
		sqlvirtualmachine.SqlVirtualMachines(),
		storage.Accounts(),
		storage.DeletedAccounts(),
		storagecache.Caches(),
		streamanalytics.StreamingJobs(),
		subscription.Subscriptions(),
		subscription.Tenants(),
		support.Services(),
		support.Tickets(),
		synapse.PrivateLinkHubs(),
		synapse.Workspaces(),
		windowsesu.MultipleActivationKeys(),
		windowsiot.Services(),
		workloads.Monitors(),
	}
}
