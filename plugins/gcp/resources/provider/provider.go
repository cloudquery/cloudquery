package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-gcp/resources/services/bigquery"
	"github.com/cloudquery/cq-provider-gcp/resources/services/cloudfunctions"
	"github.com/cloudquery/cq-provider-gcp/resources/services/compute"
	"github.com/cloudquery/cq-provider-gcp/resources/services/dns"
	"github.com/cloudquery/cq-provider-gcp/resources/services/domains"
	"github.com/cloudquery/cq-provider-gcp/resources/services/iam"
	"github.com/cloudquery/cq-provider-gcp/resources/services/kms"
	"github.com/cloudquery/cq-provider-gcp/resources/services/logging"
	"github.com/cloudquery/cq-provider-gcp/resources/services/monitoring"
	"github.com/cloudquery/cq-provider-gcp/resources/services/resource_manager"
	"github.com/cloudquery/cq-provider-gcp/resources/services/sql"
	"github.com/cloudquery/cq-provider-gcp/resources/services/storage"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	gcpMigrations embed.FS
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
		Migrations:      gcpMigrations,
		ResourceMap: map[string]*schema.Table{
			"kms.keys":                     kms.KmsKeyrings(),
			"compute.addresses":            compute.ComputeAddresses(),
			"compute.autoscalers":          compute.ComputeAutoscalers(),
			"compute.backend_services":     compute.ComputeBackendServices(),
			"compute.disk_types":           compute.ComputeDiskTypes(),
			"compute.images":               compute.ComputeImages(),
			"compute.instances":            compute.ComputeInstances(),
			"compute.interconnects":        compute.ComputeInterconnects(),
			"compute.networks":             compute.ComputeNetworks(),
			"compute.disks":                compute.ComputeDisks(),
			"compute.ssl_certificates":     compute.ComputeSslCertificates(),
			"compute.vpn_gateways":         compute.ComputeVpnGateways(),
			"compute.subnetworks":          compute.ComputeSubnetworks(),
			"compute.firewalls":            compute.ComputeFirewalls(),
			"compute.forwarding_rules":     compute.ComputeForwardingRules(),
			"compute.projects":             compute.ComputeProjects(),
			"compute.target_ssl_proxies":   compute.ComputeTargetSslProxies(),
			"compute.target_https_proxies": compute.ComputeTargetHTTPSProxies(),
			"compute.ssl_policies":         compute.ComputeSslPolicies(),
			"compute.url_maps":             compute.ComputeURLMaps(),
			"cloudfunctions.functions":     cloudfunctions.CloudfunctionsFunction(),
			"dns.managed_zones":            dns.DNSManagedZones(),
			"dns.policies":                 dns.DNSPolicies(),
			"iam.project_roles":            iam.IamRoles(),
			"iam.service_accounts":         iam.IamServiceAccounts(),
			"logging.metrics":              logging.LoggingMetrics(),
			"logging.sinks":                logging.LoggingSinks(),
			"monitoring.alert_policies":    monitoring.MonitoringAlertPolicies(),
			"resource_manager.projects":    resource_manager.ResourceManagerProjects(),
			"resource_manager.folders":     resource_manager.ResourceManagerFolders(),
			"storage.buckets":              storage.StorageBuckets(),
			"sql.instances":                sql.SQLInstances(),
			"domains.registrations":        domains.DomainsRegistration(),
			"bigquery.datasets":            bigquery.BigqueryDatasets(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
