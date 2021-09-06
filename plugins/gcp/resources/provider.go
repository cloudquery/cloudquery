package resources

import (
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "gcp",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"kms.keys":                     KmsKeyrings(),
			"compute.addresses":            ComputeAddresses(),
			"compute.autoscalers":          ComputeAutoscalers(),
			"compute.backend_services":     ComputeBackendServices(),
			"compute.disk_types":           ComputeDiskTypes(),
			"compute.images":               ComputeImages(),
			"compute.instances":            ComputeInstances(),
			"compute.interconnects":        ComputeInterconnects(),
			"compute.networks":             ComputeNetworks(),
			"compute.disks":                ComputeDisks(),
			"compute.ssl_certificates":     ComputeSslCertificates(),
			"compute.vpn_gateways":         ComputeVpnGateways(),
			"compute.subnetworks":          ComputeSubnetworks(),
			"compute.firewalls":            ComputeFirewalls(),
			"compute.forwarding_rules":     ComputeForwardingRules(),
			"compute.projects":             ComputeProjects(),
			"compute.target_ssl_proxies":   ComputeTargetSslProxies(),
			"compute.target_https_proxies": ComputeTargetHTTPSProxies(),
			"compute.ssl_policies":         ComputeSslPolicies(),
			"compute.url_maps":             ComputeURLMaps(),
			"cloudfunctions.functions":     CloudfunctionsFunction(),
			"dns.managed_zones":            DNSManagedZones(),
			"dns.policies":                 DNSPolicies(),
			"iam.project_roles":            IamRoles(),
			"iam.service_accounts":         IamServiceAccounts(),
			"logging.metrics":              LoggingMetrics(),
			"logging.sinks":                LoggingSinks(),
			"monitoring.alert_policies":    MonitoringAlertPolicies(),
			"resource_manager.projects":    ResourceManagerProjects(),
			"resource_manager.folders":     ResourceManagerFolders(),
			"storage.buckets":              StorageBuckets(),
			"sql.instances":                SQLInstances(),
			"domains.registrations":        DomainsRegistration(),
			"crm.projects":                 CrmProjects(),
			"bigquery.datasets":            BigqueryDatasets(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
