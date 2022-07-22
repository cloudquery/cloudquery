package provider

import (
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-gcp/resources/services/bigquery"
	"github.com/cloudquery/cq-provider-gcp/resources/services/cloudbilling"
	"github.com/cloudquery/cq-provider-gcp/resources/services/cloudfunctions"
	"github.com/cloudquery/cq-provider-gcp/resources/services/cloudrun"
	"github.com/cloudquery/cq-provider-gcp/resources/services/compute"
	"github.com/cloudquery/cq-provider-gcp/resources/services/dns"
	"github.com/cloudquery/cq-provider-gcp/resources/services/domains"
	"github.com/cloudquery/cq-provider-gcp/resources/services/iam"
	"github.com/cloudquery/cq-provider-gcp/resources/services/kms"
	"github.com/cloudquery/cq-provider-gcp/resources/services/kubernetes"
	"github.com/cloudquery/cq-provider-gcp/resources/services/logging"
	"github.com/cloudquery/cq-provider-gcp/resources/services/memorystore"
	"github.com/cloudquery/cq-provider-gcp/resources/services/monitoring"
	"github.com/cloudquery/cq-provider-gcp/resources/services/resource_manager"
	"github.com/cloudquery/cq-provider-gcp/resources/services/security"
	"github.com/cloudquery/cq-provider-gcp/resources/services/sql"
	"github.com/cloudquery/cq-provider-gcp/resources/services/storage"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            "gcp",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"bigquery.datasets":            bigquery.BigqueryDatasets(),
			"cloudbilling.accounts":        cloudbilling.Accounts(),
			"cloudbilling.services":        cloudbilling.Services(),
			"cloudfunctions.functions":     cloudfunctions.CloudfunctionsFunction(),
			"cloudrun.services":            cloudrun.Services(),
			"compute.addresses":            compute.ComputeAddresses(),
			"compute.autoscalers":          compute.ComputeAutoscalers(),
			"compute.backend_services":     compute.ComputeBackendServices(),
			"compute.disk_types":           compute.ComputeDiskTypes(),
			"compute.disks":                compute.ComputeDisks(),
			"compute.firewalls":            compute.ComputeFirewalls(),
			"compute.forwarding_rules":     compute.ComputeForwardingRules(),
			"compute.images":               compute.ComputeImages(),
			"compute.instance_groups":      compute.InstanceGroups(),
			"compute.instances":            compute.ComputeInstances(),
			"compute.interconnects":        compute.ComputeInterconnects(),
			"compute.networks":             compute.ComputeNetworks(),
			"compute.projects":             compute.ComputeProjects(),
			"compute.ssl_certificates":     compute.ComputeSslCertificates(),
			"compute.ssl_policies":         compute.ComputeSslPolicies(),
			"compute.subnetworks":          compute.ComputeSubnetworks(),
			"compute.target_http_proxies":  compute.ComputeTargetHTTPProxies(),
			"compute.target_https_proxies": compute.ComputeTargetHTTPSProxies(),
			"compute.target_ssl_proxies":   compute.ComputeTargetSslProxies(),
			"compute.url_maps":             compute.ComputeURLMaps(),
			"compute.vpn_gateways":         compute.ComputeVpnGateways(),
			"dns.managed_zones":            dns.DNSManagedZones(),
			"dns.policies":                 dns.DNSPolicies(),
			"domains.registrations":        domains.DomainsRegistration(),
			"iam.project_roles":            iam.IamRoles(),
			"iam.service_accounts":         iam.IamServiceAccounts(),
			"kms.keys":                     kms.KmsKeyrings(),
			"kubernetes.clusters":          kubernetes.Clusters(),
			"logging.metrics":              logging.LoggingMetrics(),
			"logging.sinks":                logging.LoggingSinks(),
			"memorystore.redis_instances":  memorystore.RedisInstances(),
			"monitoring.alert_policies":    monitoring.MonitoringAlertPolicies(),
			"resource_manager.folders":     resource_manager.ResourceManagerFolders(),
			"resource_manager.projects":    resource_manager.ResourceManagerProjects(),
			"sql.instances":                sql.SQLInstances(),
			"security.secrets":             security.Secrets(),
			"storage.buckets":              storage.StorageBuckets(),
			"storage.metrics":              storage.Metrics(),
			// "serviceusage.services":        serviceusage.Services(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
