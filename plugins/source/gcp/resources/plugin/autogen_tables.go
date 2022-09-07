// Code generated by codegen; DO NOT EDIT directly.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/resources/services/bigquery"
	"github.com/cloudquery/plugins/source/gcp/resources/services/billing"
	"github.com/cloudquery/plugins/source/gcp/resources/services/compute"
	"github.com/cloudquery/plugins/source/gcp/resources/services/container"
	"github.com/cloudquery/plugins/source/gcp/resources/services/dns"
	"github.com/cloudquery/plugins/source/gcp/resources/services/domains"
	"github.com/cloudquery/plugins/source/gcp/resources/services/functions"
	"github.com/cloudquery/plugins/source/gcp/resources/services/iam"
	"github.com/cloudquery/plugins/source/gcp/resources/services/kms"
	"github.com/cloudquery/plugins/source/gcp/resources/services/logging"
	"github.com/cloudquery/plugins/source/gcp/resources/services/monitoring"
	"github.com/cloudquery/plugins/source/gcp/resources/services/redis"
	"github.com/cloudquery/plugins/source/gcp/resources/services/resourcemanager"
	"github.com/cloudquery/plugins/source/gcp/resources/services/run"
	"github.com/cloudquery/plugins/source/gcp/resources/services/secretmanager"
	"github.com/cloudquery/plugins/source/gcp/resources/services/serviceusage"
	"github.com/cloudquery/plugins/source/gcp/resources/services/sql"
	"github.com/cloudquery/plugins/source/gcp/resources/services/storage"
)

func PluginAutoGeneratedTables() []*schema.Table {
	return []*schema.Table{
		compute.Addresses(),
		compute.Autoscalers(),
		compute.BackendServices(),
		compute.DiskTypes(),
		compute.Disks(),
		compute.ForwardingRules(),
		compute.Instances(),
		compute.SslCertificates(),
		compute.Subnetworks(),
		compute.TargetHttpProxies(),
		compute.UrlMaps(),
		compute.VpnGateways(),
		compute.InstanceGroups(),
		compute.Images(),
		compute.Firewalls(),
		compute.Networks(),
		compute.SslPolicies(),
		compute.Interconnects(),
		compute.TargetSslProxies(),
		compute.Projects(),
		dns.Policies(),
		dns.ManagedZones(),
		domains.Registrations(),
		iam.Roles(),
		iam.ServiceAccounts(),
		kms.Keyrings(),
		container.Clusters(),
		logging.Metrics(),
		logging.Sinks(),
		redis.Instances(),
		monitoring.AlertPolicies(),
		secretmanager.Secrets(),
		serviceusage.Services(),
		sql.Instances(),
		storage.Buckets(),
		bigquery.Datasets(),
		billing.BillingAccounts(),
		billing.Services(),
		resourcemanager.Folders(),
		resourcemanager.Projects(),
		functions.Functions(),
		run.Services(),
	}
}
