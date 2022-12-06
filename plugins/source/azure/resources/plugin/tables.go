// Code generated by codegen; DO NOT EDIT.
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
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/customerlockbox"
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
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/kusto"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/maintenance"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/managedservices"
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
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/saas"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sqlvirtualmachine"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storagecache"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscriptions"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/support"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/synapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/virtualmachineimagebuilder"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/windowsesu"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/windowsiot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/workloads"
	"github.com/cloudquery/plugin-sdk/schema"
)

func generatedTables() []*schema.Table {
	return []*schema.Table{
		aad.PrivateLinkPolicy(),
		advisor.ResourceRecommendationBase(),
		advisor.MetadataEntity(),
		advisor.SuppressionContract(),
		analysisservices.Server(),
		apimanagement.ServiceResource(),
		appcomplianceautomation.ReportResource(),
		appconfiguration.ConfigurationStore(),
		applicationinsights.Component(),
		applicationinsights.WebTest(),
		appplatform.ServiceResource(),
		appservice.Domain(),
		appservice.TopLevelDomain(),
		appservice.ResourceHealthMetadata(),
		appservice.CertificateOrder(),
		appservice.StaticSiteArmResource(),
		appservice.EnvironmentResource(),
		appservice.Plan(),
		appservice.AppCertificate(),
		appservice.DeletedSite(),
		appservice.Recommendation(),
		appservice.Site(),
		authorization.RoleAssignment(),
		authorization.ProviderOperationsMetadata(),
		authorization.ClassicAdministrator(),
		automanage.ConfigurationProfileAssignment(),
		automation.Account(),
		avs.PrivateCloud(),
		azurearcdata.SqlManagedInstance(),
		azurearcdata.SqlServerInstance(),
		azurearcdata.PostgresInstance(),
		azuredata.SqlServerRegistration(),
		batch.Account(),
		billing.Account(),
		billing.Period(),
		billing.EnrollmentAccountSummary(),
		botservice.Bot(),
		cdn.EdgeNode(),
		cdn.Profile(),
		cdn.WebApplicationFirewallPolicy(),
		cdn.ManagedRuleSetDefinition(),
		cognitiveservices.Account(),
		cognitiveservices.Account(),
		compute.CloudService(),
		compute.AvailabilitySet(),
		compute.Disk(),
		compute.RestorePointCollection(),
		compute.Snapshot(),
		compute.DiskEncryptionSet(),
		compute.Gallery(),
		compute.VirtualMachineScaleSet(),
		compute.VirtualMachine(),
		compute.Image(),
		compute.DiskAccess(),
		confluent.AgreementResource(),
		connectedvmware.VirtualMachine(),
		connectedvmware.Cluster(),
		connectedvmware.VirtualNetwork(),
		connectedvmware.VirtualMachineTemplate(),
		connectedvmware.Datastore(),
		connectedvmware.Host(),
		connectedvmware.ResourcePool(),
		connectedvmware.VCenter(),
		containerinstance.ContainerGroup(),
		containerregistry.Registry(),
		containerservice.ManagedCluster(),
		containerservice.Snapshot(),
		cosmos.DatabaseAccountGetResults(),
		cosmos.LocationGetResult(),
		cosmos.RestorableDatabaseAccountGetResult(),
		costmanagement.View(),
		customerinsights.Hub(),
		customerlockbox.LockboxRequestResponse(),
		dashboard.ManagedGrafana(),
		databox.JobResource(),
		datadog.AgreementResource(),
		datadog.MonitorResource(),
		datafactory.Factory(),
		datalakeanalytics.AccountBasic(),
		datalakestore.AccountBasic(),
		datamigration.Service(),
		desktopvirtualization.HostPool(),
		devhub.Workflow(),
		devops.PipelineTemplateDefinition(),
		dns.Zone(),
		dnsresolver.DnsResolver(),
		dnsresolver.DnsForwardingRuleset(),
		elastic.MonitorResource(),
		engagementfabric.Account(),
		eventgrid.TopicTypeInfo(),
		eventhub.EhNamespace(),
		frontdoor.WebApplicationFirewallPolicy(),
		frontdoor.Profile(),
		frontdoor.FrontDoor(),
		frontdoor.ManagedRuleSetDefinition(),
		hanaonazure.SapMonitor(),
		hdinsight.Cluster(),
		healthbot.HealthBot(),
		healthcareapis.ServicesDescription(),
		hybridcompute.PrivateLinkScope(),
		hybriddatamanager.DataManager(),
		keyvault.Resource(),
		kusto.Cluster(),
		maintenance.Configuration(),
		maintenance.ApplyUpdate(),
		maintenance.ApplyUpdate(),
		maintenance.Configuration(),
		maintenance.Configuration(),
		managedservices.MarketplaceRegistrationDefinition(),
		managementgroups.EntityInfo(),
		mariadb.Server(),
		marketplace.PrivateStore(),
		migrate.Project(),
		monitor.AzureMonitorPrivateLinkScope(),
		monitor.LogProfileResource(),
		monitor.EventData(),
		mysql.Server(),
		mysqlflexibleservers.Server(),
		netapp.Account(),
		network.PrivateLinkService(),
		network.ApplicationGateway(),
		network.BastionHost(),
		network.ExpressRouteCrossConnection(),
		network.VirtualNetwork(),
		network.Profile(),
		network.ServiceEndpointPolicy(),
		network.RouteFilter(),
		network.NatGateway(),
		network.VirtualRouter(),
		network.LocalNetworkGateway(),
		network.PublicIpAddress(),
		network.VirtualHub(),
		network.VirtualAppliance(),
		network.RouteTable(),
		network.ApplicationSecurityGroup(),
		network.PublicIpPrefix(),
		network.IpAllocation(),
		network.Watcher(),
		network.Manager(),
		network.WebApplicationFirewallPolicy(),
		network.AzureFirewall(),
		network.ExpressRoutePortsLocation(),
		network.PrivateEndpoint(),
		network.Interface(),
		network.ManagerConnection(),
		network.VirtualApplianceSku(),
		network.LoadBalancer(),
		network.VirtualNetworkGatewayConnection(),
		network.FirewallPolicy(),
		network.VirtualNetworkGateway(),
		network.BgpServiceCommunity(),
		network.IpGroup(),
		network.VpnGateway(),
		network.CustomIpPrefix(),
		network.SecurityGroup(),
		network.ExpressRoutePort(),
		network.DscpConfiguration(),
		network.VirtualWan(),
		network.VpnServerConfiguration(),
		network.ExpressRouteServiceProvider(),
		network.ExpressRouteCircuit(),
		network.DdosProtectionPlan(),
		network.VpnSite(),
		network.SecurityPartnerProvider(),
		networkfunction.AzureTrafficCollector(),
		networkfunction.AzureTrafficCollector(),
		nginx.Deployment(),
		notificationhubs.NamespaceResource(),
		operationalinsights.Workspace(),
		operationalinsights.Workspace(),
		operationalinsights.Cluster(),
		orbital.ContactProfile(),
		orbital.Spacecraft(),
		peering.ServiceCountry(),
		peering.ServiceProvider(),
		peering.ServiceLocation(),
		policy.Exemption(),
		policy.SetDefinition(),
		policy.Assignment(),
		policy.DataPolicyManifest(),
		policy.Definition(),
		portal.Violation(),
		portal.Configuration(),
		postgresql.Server(),
		postgresqlflexibleservers.Server(),
		postgresqlhsc.ServerGroup(),
		powerbidedicated.DedicatedCapacity(),
		privatedns.PrivateZone(),
		providerhub.ProviderRegistration(),
		redhatopenshift.OpenShiftCluster(),
		relay.Namespace(),
		reservations.ReservationOrderResponse(),
		resources.TagDetails(),
		resources.Provider(),
		resources.ResourceGroup(),
		saas.Resource(),
		saas.App(),
		security.Application(),
		security.WorkspaceSetting(),
		security.Automation(),
		security.Solution(),
		security.GovernanceRule(),
		security.SecureScoreControlDefinitionItem(),
		security.SecureScoreItem(),
		security.Task(),
		security.SecureScoreControlDetails(),
		security.AscLocation(),
		security.AutoProvisioningSetting(),
		security.DiscoveredSecuritySolution(),
		security.AllowedConnectionsResource(),
		security.JitNetworkAccessPolicy(),
		security.ExternalSecuritySolution(),
		security.Alert(),
		security.TopologyResource(),
		security.Contact(),
		security.AlertsSuppressionRule(),
		security.ConnectorSetting(),
		security.IngestionSetting(),
		security.AssessmentMetadataResponse(),
		security.Connector(),
		security.RegulatoryComplianceStandard(),
		servicebus.SbNamespace(),
		sql.DeletedServer(),
		sql.ManagedInstance(),
		sql.VirtualCluster(),
		sql.InstancePool(),
		sql.Server(),
		sqlvirtualmachine.Group(),
		sqlvirtualmachine.SqlVirtualMachine(),
		storage.DeletedAccount(),
		storage.Account(),
		storagecache.Cache(),
		streamanalytics.StreamingJob(),
		subscription.Subscription(),
		subscription.TenantIdDescription(),
		subscriptions.TenantIdDescription(),
		support.TicketDetails(),
		support.Service(),
		synapse.PrivateLinkHub(),
		synapse.Workspace(),
		virtualmachineimagebuilder.ImageTemplate(),
		windowsesu.MultipleActivationKey(),
		windowsiot.DeviceService(),
		workloads.Monitor(),
	}
}
