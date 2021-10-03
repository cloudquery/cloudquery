package resources

import (
	"embed"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	gcpMigrations embed.FS
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:       "gcp",
		Configure:  client.Configure,
		Migrations: gcpMigrations,
		ResourceMap: map[string]*schema.Table{
			"bigquery.datasets":            BigqueryDatasets(),
			"cloudfunctions.functions":     CloudfunctionsFunction(),
			"compute.addresses":            ComputeAddresses(),
			"compute.autoscalers":          ComputeAutoscalers(),
			"compute.backend_services":     ComputeBackendServices(),
			"compute.disk_types":           ComputeDiskTypes(),
			"compute.disks":                ComputeDisks(),
			"compute.firewalls":            ComputeFirewalls(),
			"compute.forwarding_rules":     ComputeForwardingRules(),
			"compute.images":               ComputeImages(),
			"compute.instances":            ComputeInstances(),
			"compute.interconnects":        ComputeInterconnects(),
			"compute.networks":             ComputeNetworks(),
			"compute.projects":             ComputeProjects(),
			"compute.ssl_certificates":     ComputeSslCertificates(),
			"compute.ssl_policies":         ComputeSslPolicies(),
			"compute.subnetworks":          ComputeSubnetworks(),
			"compute.target_https_proxies": ComputeTargetHTTPSProxies(),
			"compute.target_ssl_proxies":   ComputeTargetSslProxies(),
			"compute.url_maps":             ComputeURLMaps(),
			"compute.vpn_gateways":         ComputeVpnGateways(),
			"dns.managed_zones":            DNSManagedZones(),
			"dns.policies":                 DNSPolicies(),
			"domains.registrations":        DomainsRegistration(),
			"iam.project_roles":            IamRoles(),
			"iam.service_accounts":         IamServiceAccounts(),
			"kms.keys":                     KmsKeyrings(),
			"logging.metrics":              LoggingMetrics(),
			"logging.sinks":                LoggingSinks(),
			"monitoring.alert_policies":    MonitoringAlertPolicies(),
			"resource_manager.folders":     ResourceManagerFolders(),
			"resource_manager.projects":    ResourceManagerProjects(),
			"sql.instances":                SQLInstances(),
			"storage.buckets":              StorageBuckets(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
